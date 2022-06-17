package examples

import (
	"context"
	"fmt"
	"time"

	"github.com/apache/arrow/go/v8/arrow"
	"github.com/apache/arrow/go/v8/arrow/array"
	"github.com/apache/arrow/go/v8/arrow/memory"
	linuxproc "github.com/c9s/goprocinfo/linux"

	"github.com/deephaven/deephaven-core/go-client/client"
)

func MakeRecord(timestamp time.Time, cpuTime float64, idle float64) arrow.Record {
	pool := memory.NewGoAllocator()

	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "Timestamp", Type: arrow.PrimitiveTypes.Int64},
			{Name: "CpuTime", Type: arrow.PrimitiveTypes.Float64},
			{Name: "Idle", Type: arrow.PrimitiveTypes.Float64},
		},
		nil,
	)

	b := array.NewRecordBuilder(pool, schema)
	defer b.Release()

	b.Field(0).(*array.Int64Builder).AppendValues([]int64{timestamp.UnixNano()}, nil)
	b.Field(1).(*array.Float64Builder).AppendValues([]float64{cpuTime}, nil)
	b.Field(2).(*array.Float64Builder).AppendValues([]float64{idle}, nil)

	return b.NewRecord()
}

func ExampleCpuUsage() {
	ctx := context.Background()

	cl, err := client.NewClient(ctx, "localhost", "10000", "python")
	if err != nil {
		fmt.Println("couldn't open client", err.Error())
		return
	}
	defer cl.Close()

	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "Timestamp", Type: arrow.PrimitiveTypes.Int64},
			{Name: "CpuTime", Type: arrow.PrimitiveTypes.Float64},
			{Name: "Idle", Type: arrow.PrimitiveTypes.Float64},
		},
		nil,
	)

	inputTable, err := cl.NewAppendOnlyInputTableFromSchema(ctx, schema)
	if err != nil {
		fmt.Println("couldn't get input table", err.Error())
		return
	}
	defer inputTable.Release(ctx)

	timedQuery := inputTable.
		Query().
		Update("Timestamp = nanosToTime(Timestamp)", "PastTimestamp = Timestamp - 5 * SECOND")

	joinedQuery := timedQuery.
		AsOfJoin(timedQuery, []string{"PastTimestamp = Timestamp"}, []string{"PastTime = Timestamp", "PastCpuTime = CpuTime", "PastIdle = Idle"}, client.MatchRuleLessThanEqual).
		DropColumns("PastTimestamp").
		Update(
			"CpuTimeChange = CpuTime - PastCpuTime", "IdleTimeChange = Idle - PastIdle",
			"TimeChange = diffNanos(PastTime, Timestamp) / 10000000.0", "CpuUtil = 100.0 * (1.0 - (IdleTimeChange / CpuTimeChange))")

	cleanQuery := joinedQuery.
		DropColumns("CpuTimeChange", "IdleTimeChange", "TimeChange", "PastCpuTime", "PastIdle", "PastTime")

	tables, err := cl.ExecQuery(ctx, timedQuery, joinedQuery, cleanQuery)
	if err != nil {
		fmt.Println("couldn't get query", err.Error())
		return
	}
	timedTable, joinedTable, cleanTable := tables[0], tables[1], tables[2]
	defer timedTable.Release(ctx)
	defer joinedTable.Release(ctx)
	defer cleanTable.Release(ctx)

	err = cl.BindToVariable(ctx, "go_cpu_usage_raw", joinedTable)
	if err != nil {
		fmt.Println("couldn't bind table to variable", err.Error())
		return
	}

	err = cl.BindToVariable(ctx, "go_cpu_usage", cleanTable)
	if err != nil {
		fmt.Println("couldn't bind table to variable", err.Error())
		return
	}

	for {
		cpuStat, err := linuxproc.ReadStat("/proc/stat")
		if err != nil {
			fmt.Println("couldn't get stat", err.Error())
			return
		}

		totalTime := cpuStat.CPUStatAll.User + cpuStat.CPUStatAll.Nice + cpuStat.CPUStatAll.System + cpuStat.CPUStatAll.Idle +
			cpuStat.CPUStatAll.IOWait + cpuStat.CPUStatAll.IRQ + cpuStat.CPUStatAll.SoftIRQ
		idleTime := cpuStat.CPUStatAll.Idle

		newRec := MakeRecord(time.Now(), float64(totalTime), float64(idleTime))
		defer newRec.Release()
		newData, err := cl.ImportTable(ctx, newRec)
		if err != nil {
			fmt.Println("couldn't import table", err.Error())
			return
		}
		defer newData.Release(ctx)

		err = inputTable.AddTable(ctx, newData)
		if err != nil {
			fmt.Println("couldn't add data", err.Error())
			return
		}

		time.Sleep(time.Second)
	}
}
