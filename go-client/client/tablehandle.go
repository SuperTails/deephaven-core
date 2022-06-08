package client

import (
	"context"

	"github.com/apache/arrow/go/arrow"
	"github.com/apache/arrow/go/arrow/array"

	ticketpb2 "github.com/deephaven/deephaven-core/go-client/internal/proto/ticket"
)

type TableHandle struct {
	client   *Client
	ticket   *ticketpb2.Ticket
	schema   *arrow.Schema
	size     int64
	isStatic bool
}

func newTableHandle(client *Client, ticket *ticketpb2.Ticket, schema *arrow.Schema, size int64, isStatic bool) TableHandle {
	return TableHandle{
		client:   client,
		ticket:   ticket,
		schema:   schema,
		size:     size,
		isStatic: isStatic,
	}
}

// Downloads the current state of the table on the server and returns it as a Record.
//
// If a Record is returned successfully, it must be freed later with `record.Release()`
func (th *TableHandle) Snapshot(ctx context.Context) (array.Record, error) {
	return th.client.SnapshotRecord(ctx, th.ticket)
}

// Returns a new table without the given columns.
func (th *TableHandle) DropColumns(ctx context.Context, cols []string) (TableHandle, error) {
	return th.client.DropColumns(ctx, th, cols)
}

// Returns a new table with additional columns calculated according to the formulas.
func (th *TableHandle) Update(ctx context.Context, formulas []string) (TableHandle, error) {
	return th.client.Update(ctx, th, formulas)
}

func (th *TableHandle) Query() *QueryBuilder {
	qb := newQueryBuilder(th)
	return &qb
}

/* ... more table methods would go here ... */
