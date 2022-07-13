// This package allows you to interface with a Deephaven server over a network connection using Go.
// It can upload, manipulate, and download tables, among other features.
//
// To get started, use client.NewClient to connect to the server. The Client can then be used to perform operations.
// See the provided examples in the examples/ folder or the individual code documentation for more.
//
// Online docs for the client can be found at /* TODO: deephaven.io link goes here once docs are hosted. */
//
// The Go API uses Records from the Apache Arrow package as tables.
// The docs for the Arrow package can be found at the following link:
// https://pkg.go.dev/github.com/apache/arrow/go/v8
//
// All methods for all structs in this package are thread-safe unless otherwise specified.
package client

import (
	"context"
	"errors"
	apppb2 "github.com/deephaven/deephaven-core/go-client/internal/proto/application"
	consolepb2 "github.com/deephaven/deephaven-core/go-client/internal/proto/console"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"log"
	"sync"
)

// ErrClosedClient is returned as an error when trying to perform a network operation on a client that has been closed.
var ErrClosedClient = errors.New("client is closed")

// Maintains a connection to a Deephaven server.
// It can be used to run scripts, create new tables, execute queries, etc.
// Check the various methods of Client to learn more.
type Client struct {
	// This lock guards isClosed.
	// Other functionality does not need a lock, since the gRPC interface is already thread-safe,
	// the tables array has its own lock, and the session token also has its own lock.
	lock     sync.Mutex
	isClosed bool // True if Close has been called (i.e. the client can no longer perform operations).

	grpcChannel *grpc.ClientConn

	sessionStub
	consoleStub
	flightStub
	tableStub
	inputTableStub

	appServiceClient apppb2.ApplicationServiceClient
	ticketFact       ticketFactory
	fieldMan         fieldManager
}

// NewClient starts a connection to a Deephaven server.
//
// scriptLanguage can be either "python" or "groovy", and must match the language used on the server. Python is the default.
//
// The client should be closed using Close() after it is done being used.
//
// Keepalive messages are sent automatically by the client to the server at a regular interval (~30 seconds)
// so that the connection remains open. The provided context is saved and used to send keepalive messages.
func NewClient(ctx context.Context, host string, port string, scriptLanguage string) (*Client, error) {
	grpcChannel, err := grpc.Dial(host+":"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	client := &Client{grpcChannel: grpcChannel, isClosed: false}

	client.ticketFact = newTicketFactory()

	client.sessionStub, err = newSessionStub(ctx, client)
	if err != nil {
		client.Close()
		return nil, err
	}

	client.consoleStub, err = newConsoleStub(ctx, client, scriptLanguage)
	if err != nil {
		client.Close()
		return nil, err
	}

	client.flightStub, err = newFlightStub(client, host, port)
	if err != nil {
		client.Close()
		return nil, err
	}

	client.tableStub = newTableStub(client)

	client.inputTableStub = newInputTableStub(client)

	client.appServiceClient = apppb2.NewApplicationServiceClient(client.grpcChannel)

	client.fieldMan = newFieldManager(client)

	return client, nil
}

// Closed checks if the client is closed, i.e. it can no longer perform operations on the server.
func (client *Client) Closed() bool {
	return client.isClosed
}

// FetchTablesOnce fetches the list of tables from the server.
// This allows the client to see the list of named global tables on the server,
// and thus allows the client to open them using OpenTable.
// Tables created in scripts run by the current client are immediately visible and do not require a FetchTables call.
func (client *Client) FetchTablesOnce(ctx context.Context) error {
	ctx, err := client.withToken(ctx)
	if err != nil {
		return err
	}

	return client.fieldMan.FetchTablesOnce(ctx, client.appServiceClient)
}

// FetchTablesRepeating starts up a goroutine that fetches the list of tables from the server continuously.
// This allows the client to see the list of named global tables on the server,
// and thus allows the client to open them using OpenTable.
// Tables created in scripts run by the current client are immediately visible and do not require a FetchTables call.
// The returned error channel is guaranteed to send zero or one errors and then close,
// however it is not specified at what time this will occur.
func (client *Client) FetchTablesRepeating(ctx context.Context) <-chan error {
	ctx, err := client.withToken(ctx)
	if err != nil {
		chanError := make(chan error, 1)
		chanError <- err
		close(chanError)
		return chanError
	}

	return client.fieldMan.FetchTablesRepeating(ctx, client.appServiceClient)
}

// ListOpenableTables returns a list of the (global) tables that can be opened with OpenTable.
// Tables that are created by other clients or in the web UI are not listed here automatically.
// Tables that are created in scripts run by this client, however, are immediately available,
// and will be added to/removed from the list as soon as the script finishes.
// FetchTablesOnce or FetchTablesRepeating can be used to update the list
// to reflect what tables are currently available from other clients or the web UI.
func (client *Client) ListOpenableTables() []string {
	return client.fieldMan.ListOpenableTables()
}

// ExecSerial executes several table operations on the server and returns the resulting tables.
//
// This function makes a request for each table operation.
// Consider using ExecBatch to batch all of the table operations into a single request,
// which can be more efficient.
//
// If this function completes successfully,
// the number of tables returned will always match the number of query nodes passed.
// The first table in the returned list corresponds to the first node argument,
// the second table in the returned list corresponds to the second node argument,
// etc.
//
// This may return a QueryError if the query is invalid.
func (client *Client) ExecSerial(ctx context.Context, nodes ...QueryNode) ([]*TableHandle, error) {
	return execSerial(ctx, client, nodes)
}

// ExecBatch executes a batched query on the server and returns the resulting tables.
//
// All of the operations in the query will be performed in a single request,
// so ExecBatch is usually more efficient than ExecSerial.
//
// If this function completes successfully,
// the number of tables returned will always match the number of query nodes passed.
// The first table in the returned list corresponds to the first node argument,
// the second table in the returned list corresponds to the second node argument,
// etc.
//
// This may return a QueryError if the query is invalid.
func (client *Client) ExecBatch(ctx context.Context, nodes ...QueryNode) ([]*TableHandle, error) {
	return execBatch(client, ctx, nodes)
}

// Close closes the connection to the server and frees any associated resources.
// Once this method is called, the client and any TableHandles from it cannot be used.
func (client *Client) Close() error {
	client.lock.Lock()
	defer client.lock.Unlock()

	if client.isClosed {
		return nil
	}

	client.isClosed = true

	client.sessionStub.Close()

	if client.grpcChannel != nil {
		client.grpcChannel.Close()
		client.grpcChannel = nil
	}

	// This is logged because most of the time this method is used with defer,
	// which will discard the error value.
	err := client.flightStub.Close()
	if err != nil {
		log.Println("unable to close client:", err.Error())
	}

	client.fieldMan.Close()

	return err
}

// withToken attaches the current session token to a context as metadata.
func (client *Client) withToken(ctx context.Context) (context.Context, error) {
	tok, err := client.getToken()

	if err != nil {
		return nil, err
	}

	return metadata.NewOutgoingContext(ctx, metadata.Pairs("deephaven_session_id", string(tok))), nil
}

// RunScript executes a script on the deephaven server.
// The script language depends on the scriptLanguage argument passed when creating the client.
func (client *Client) RunScript(ctx context.Context, script string) error {
	ctx, err := client.consoleStub.client.withToken(ctx)
	if err != nil {
		return err
	}

	f := func() (*apppb2.FieldsChangeUpdate, error) {
		// TODO: This might fall victim to the double-responses problem if a FetchTablesRepeating
		// request is already active.

		req := consolepb2.ExecuteCommandRequest{ConsoleId: client.consoleStub.consoleId, Code: script}
		rst, err := client.consoleStub.stub.ExecuteCommand(ctx, &req)
		if err != nil {
			return nil, err
		} else {
			return rst.Changes, nil
		}
	}

	return client.fieldMan.ExecAndUpdate(f)
}
