package main

import (
	"context"
	"fmt"
	"time"

	"github.com/deephaven/deephaven-core/go-client/client"
)

func main() {
	ctx := context.Background()

	cl, err := client.NewClient(ctx, "localhost", "10000", "python")
	fmt.Println(err)

	q1 := cl.TimeTableQuery(time.Second, time.Now()).Update("foo = i % 10").Tail(5).Where("foo < 5")
	q2 := client.MergeQuery("", q1)

	tbls, err := cl.ExecBatch(ctx, q2)
	fmt.Println(err)
	tbl := tbls[0]

	_, err = cl.Subscribe(ctx, tbl)
	if err != nil {
		fmt.Println(err)
	}
}
