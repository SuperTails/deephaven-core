package main

import (
	"context"
	"fmt"

	"github.com/apache/arrow/go/v8/arrow"
	"github.com/deephaven/deephaven-core/go-client/client"
	"github.com/deephaven/deephaven-core/go-client/examples/common"
)

// This example shows how to use Input Tables.
// Input Tables make are a generic interface for streaming data from any source,
// so you can use Deephaven's streaming table processing power for anything.
func main() {
	// A context is used to set timeouts and deadlines for requests or cancel requests.
	// If you don't have any specific requirements, context.Background() is a good default.
	ctx := context.Background()

	// When starting a client connection, note that the client script language "python"
	// must match the language the server was started with,
	// even if we aren't using any scripts.
	cl, err := client.NewClient(ctx, "localhost", "10000", "python")
	if err != nil {
		fmt.Println("error when connecting to localhost port 10000:", err.Error())
		return
	}
	defer cl.Close()

	// First, let's make a schema our input table is going to use.
	// This describes the name and data types each of its columns will have.
	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "Ticker", Type: arrow.BinaryTypes.String},
			{Name: "Close", Type: arrow.PrimitiveTypes.Float32},
			{Name: "Vol", Type: arrow.PrimitiveTypes.Int32},
		},
		nil,
	)

	// Then we can actually make the input table.
	// It will start empty, but we're going to add more data to it.
	// This is a key-backed input table, so it will make sure the "Ticker" column stays unique.
	// This is in contrast to an append-only table, which will append rows to the end of the table.
	inputTable, err := cl.NewKeyBackedInputTableFromSchema(ctx, schema, "Ticker")
	if err != nil {
		fmt.Println("error when creating InputTable", err.Error())
		return
	}
	// Any tables you create should be eventually released.
	defer inputTable.Release(ctx)

	// Now let's create a table derived from the input table.
	// When we update the input table, this table will update too.
	outputTable, err := inputTable.Where(ctx, "Close > 50.0")
	if err != nil {
		fmt.Println("error when filtering input table", err.Error())
		return
	}
	defer outputTable.Release(ctx)

	// Now, let's get some new data to add to the input table.
	// We import the data so that it is available on the server.
	newDataRec := common.GetExampleRecord()
	// Note that Arrow records must be eventually released.
	defer newDataRec.Release()
	newDataTable, err := cl.ImportTable(ctx, newDataRec)
	if err != nil {
		fmt.Println("error when importing new data", err.Error())
		return
	}
	defer newDataTable.Release(ctx)

	// Now we can add the new data we just imported to our input table.
	// Since this is a key-backed table, it will add any rows with new keys
	// and replace any rows with keys that already exist.
	// Since there's currently nothing in the table,
	// this call will add all the rows of the new data to the input table.
	err = inputTable.AddTable(ctx, newDataTable)
	if err != nil {
		fmt.Println("error when adding new data to table", err.Error())
		return
	}

	// Now, we take a snapshot of the outputTable to see what data it currently contains.
	// We should see the new rows we added, filtered by the condition we specified when creating outputTable.
	outputRec, err := outputTable.Snapshot(ctx)
	if err != nil {
		fmt.Println("error when snapshotting table", err.Error())
		return
	}
	defer outputRec.Release()

	fmt.Println("Got the output table!")
	fmt.Println(outputRec)
}
