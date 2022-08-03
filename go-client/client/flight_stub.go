package client

import (
	"context"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"net"
	"strconv"

	"github.com/apache/arrow/go/v8/arrow"
	"github.com/apache/arrow/go/v8/arrow/flight"
	"github.com/apache/arrow/go/v8/arrow/ipc"
	flatbuffers "github.com/google/flatbuffers/go"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	ticketpb2 "github.com/deephaven/deephaven-core/go-client/internal/proto/ticket"

	flatbuf_b "github.com/deephaven/deephaven-core/go-client/internal/proto/barrage/flatbuf"
	flatbuf_a "github.com/deephaven/deephaven-core/go-client/org/apache/arrow/flatbuf"
)

// flightStub wraps Arrow Flight gRPC calls.
type flightStub struct {
	client *Client

	stub flight.Client // The stub for performing Arrow Flight gRPC requests.
}

func newFlightStub(client *Client, host string, port string) (flightStub, error) {
	stub, err := flight.NewClientWithMiddleware(
		net.JoinHostPort(host, port),
		nil,
		nil,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return flightStub{}, err
	}

	return flightStub{client: client, stub: stub}, nil
}

// snapshotRecord downloads the data currently in the provided table and returns it as an Arrow Record.
func (fs *flightStub) snapshotRecord(ctx context.Context, ticket *ticketpb2.Ticket) (arrow.Record, error) {
	ctx, err := fs.client.withToken(ctx)
	if err != nil {
		return nil, err
	}

	fticket := &flight.Ticket{Ticket: ticket.GetTicket()}

	req, err := fs.stub.DoGet(ctx, fticket)
	if err != nil {
		return nil, err
	}
	defer req.CloseSend()

	reader, err := flight.NewRecordReader(req)
	defer reader.Release()
	if err != nil {
		return nil, err
	}

	rec1, err := reader.Read()
	if err != nil {
		return nil, err
	}
	rec1.Retain()

	rec2, err := reader.Read()
	if err != io.EOF {
		rec1.Release()
		rec2.Release()
		return nil, errors.New("multiple records retrieved during snapshot")
	}

	return rec1, nil
}

// ImportTable uploads a table to the Deephaven server.
// The table can then be manipulated and referenced using the returned TableHandle.
func (fs *flightStub) ImportTable(ctx context.Context, rec arrow.Record) (*TableHandle, error) {
	ctx, err := fs.client.withToken(ctx)
	if err != nil {
		return nil, err
	}

	doPut, err := fs.stub.DoPut(ctx)
	if err != nil {
		return nil, err
	}
	defer doPut.CloseSend()

	ticketNum := fs.client.ticketFact.nextId()

	//todo Seems like this should be a fixed size int cast here and not a generic int
	descr := &flight.FlightDescriptor{Type: flight.DescriptorPATH, Path: []string{"export", strconv.Itoa(int(ticketNum))}}

	writer := flight.NewRecordWriter(doPut, ipc.WithSchema(rec.Schema()))

	writer.SetFlightDescriptor(descr)
	err = writer.Write(rec)
	if err != nil {
		return nil, err
	}

	err = writer.Close()
	if err != nil {
		return nil, err
	}

	_, err = doPut.Recv()
	if err != nil {
		return nil, err
	}

	ticket := fs.client.ticketFact.makeTicket(ticketNum)

	schema := rec.Schema()

	return newTableHandle(fs.client, &ticket, schema, rec.NumRows(), true), nil
}

func decodeRowSet(bytes []byte) {
	for i := 0; i < len(bytes); {
		cmd := bytes[i]
		cmdType := cmd >> 3
		valType := cmd & 0x7

		i++

		/*

			long pending = -1;
			long lastValue = 0;
			void consume(long nextOffset) {
				if (nextOffset < 0) {
					assert(pending != -1);
					lastValue = lastValue - nextOffset;
					addRowsInRange(pending, lastValue);
					pending = -1;
				} else {
					if (pending != -1) {
						addRowAt(pending);
					}
					lastValue = pending = lastValue + nextOffset;
				}
			}

			pending = -1
			lastValue = 0

			pending = 11
			lastValue = 11

			???
			lastValue =
			addRowsInRange(11, 12) ???

		*/

		switch cmdType {
		case 1:
			fmt.Print("offset ")
			switch valType {
			case 1: // short
				fmt.Print("!1")
			case 2: // int
				val := int32(binary.LittleEndian.Uint32(bytes[i : i+4]))
				i += 4
				fmt.Print(val, " ")
			case 3: // long
				fmt.Print("!3")
			case 4: // byte
				val := bytes[i]
				i++
				fmt.Print(val, " ")
			default:
				fmt.Print("???")
			}
		case 2:
			fmt.Print("short_array", valType)
		case 3:
			fmt.Print("byte_array", valType)
		case 4:
			fmt.Print("end (", valType, ")")
		default:
			fmt.Print("unknown command type ", cmdType)
		}
	}
	fmt.Println()
}

func (fs *flightStub) Subscribe(ctx context.Context, handle *TableHandle) (*int, error) {
	ctx, err := fs.client.withToken(ctx)
	if err != nil {
		return nil, err
	}

	doExchg, err := fs.stub.DoExchange(ctx)
	if err != nil {
		return nil, err
	}
	defer doExchg.CloseSend()

	builder_b := flatbuffers.NewBuilder(0)
	/*flatbuf_b.BarrageSubscriptionOptionsStart(builder_b)
	flatbuf_b.BarrageSubscriptionOptionsAddUseDeephavenNulls(builder_b, false)
	flatbuf_b.BarrageSubscriptionOptionsAddBatchSize(builder_b, 4096)
	flatbuf_b.BarrageSubscriptionOptionsAddMaxMessageSize(builder_b, 2000000000)
	opts := flatbuf_b.BarrageSubscriptionOptionsEnd(builder_b)*/

	flatbuf_b.BarrageSubscriptionRequestStartTicketVector(builder_b, len(handle.ticket.Ticket))
	for i := len(handle.ticket.Ticket) - 1; i >= 0; i-- {
		builder_b.PrependByte(handle.ticket.Ticket[i])
	}
	ticketVec := builder_b.EndVector(len(handle.ticket.Ticket))
	/*flatbuf_b.BarrageSubscriptionRequestStartColumnsVector(builder_b, 0)
	columnVec := builder_b.EndVector(0)
	flatbuf_b.BarrageSubscriptionRequestStartViewportVector(builder_b, 0)
	viewportVec := builder_b.EndVector(0)*/
	flatbuf_b.BarrageSubscriptionRequestStart(builder_b)
	flatbuf_b.BarrageSubscriptionRequestAddTicket(builder_b, ticketVec)
	//flatbuf_b.BarrageSubscriptionRequestAddColumns(builder_b, columnVec)
	//flatbuf_b.BarrageSubscriptionRequestAddViewport(builder_b, viewportVec)
	flatbuf_b.BarrageSubscriptionRequestAddReverseViewport(builder_b, false)
	//flatbuf_b.BarrageSubscriptionRequestAddSubscriptionOptions(builder_b, opts)
	bsrOff := flatbuf_b.BarrageSubscriptionRequestEnd(builder_b)
	builder_b.Finish(bsrOff)
	bsrPayload := builder_b.FinishedBytes()

	builder_c := flatbuffers.NewBuilder(0)
	flatbuf_b.BarrageMessageWrapperStartMsgPayloadVector(builder_c, len(bsrPayload))
	for i := len(bsrPayload) - 1; i >= 0; i-- {
		builder_c.PrependByte(bsrPayload[i])
	}
	bsrPayloadVec := builder_c.EndVector(len(bsrPayload))
	flatbuf_b.BarrageMessageWrapperStart(builder_c)
	flatbuf_b.BarrageMessageWrapperAddMagic(builder_c, 0x6E687064)
	flatbuf_b.BarrageMessageWrapperAddMsgType(builder_c, flatbuf_b.BarrageMessageTypeBarrageSubscriptionRequest)
	flatbuf_b.BarrageMessageWrapperAddMsgPayload(builder_c, bsrPayloadVec)
	bmwOff := flatbuf_b.BarrageMessageWrapperEnd(builder_c)
	builder_c.Finish(bmwOff)
	customMeta := builder_c.FinishedBytes()

	builder_a := flatbuffers.NewBuilder(0)
	flatbuf_a.MessageStart(builder_a)
	flatbuf_a.MessageAddHeaderType(builder_a, flatbuf_a.MessageHeaderNONE)
	msgOff := flatbuf_a.MessageEnd(builder_a)
	builder_a.Finish(msgOff)
	msgBuf := builder_a.FinishedBytes()

	desc := &flight.FlightDescriptor{Type: flight.DescriptorCMD, Cmd: []byte{0x64, 0x70, 0x68, 0x6E}}
	data := &flight.FlightData{DataHeader: msgBuf, AppMetadata: customMeta, FlightDescriptor: desc}

	msg := flatbuf_b.GetRootAsBarrageMessageWrapper(customMeta, 0)
	fmt.Println(msg)

	err = doExchg.Send(data)
	if err != nil {
		return nil, err
	}

	reader, err := flight.NewRecordReader(doExchg)
	if err != nil {
		return nil, err
	}

	for reader.Next() {
		if reader.Err() != nil {
			return nil, err
		}

		record := reader.Record()

		fmt.Println(record)

		resp := reader.LatestAppMetadata()

		if len(resp) > 0 {
			meta := resp
			msgWrapper := flatbuf_b.GetRootAsBarrageMessageWrapper(meta, 0)

			payload := make([]byte, msgWrapper.MsgPayloadLength())
			for i := 0; i < msgWrapper.MsgPayloadLength(); i++ {
				payload[i] = byte(msgWrapper.MsgPayload(i))
			}

			fmt.Println(msgWrapper.MsgType())
			if msgWrapper.MsgType() == flatbuf_b.BarrageMessageTypeBarrageUpdateMetadata {
				updateMeta := flatbuf_b.GetRootAsBarrageUpdateMetadata(payload, 0)
				fmt.Println()
				fmt.Println("first_seq:", updateMeta.FirstSeq())
				fmt.Println("last_seq:", updateMeta.LastSeq())
				fmt.Println("is_snapshot", updateMeta.IsSnapshot())

				addedRows := make([]byte, updateMeta.AddedRowsLength())
				for i := 0; i < updateMeta.AddedRowsLength(); i++ {
					addedRows[i] = byte(updateMeta.AddedRows(i))
				}

				removedRows := make([]byte, updateMeta.RemovedRowsLength())
				for i := 0; i < updateMeta.RemovedRowsLength(); i++ {
					removedRows[i] = byte(updateMeta.RemovedRows(i))
				}

				addedRowsIncluded := make([]byte, updateMeta.AddedRowsIncludedLength())
				for i := 0; i < updateMeta.AddedRowsIncludedLength(); i++ {
					addedRowsIncluded[i] = byte(updateMeta.AddedRowsIncluded(i))
				}

				shiftData := make([]byte, updateMeta.ShiftDataLength())
				for i := 0; i < updateMeta.ShiftDataLength(); i++ {
					shiftData[i] = byte(updateMeta.ShiftData(i))
				}

				fmt.Print("added rows (", updateMeta.AddedRowsLength(), "): ")
				decodeRowSet(addedRows)

				fmt.Print("added rows inc (", updateMeta.AddedRowsIncludedLength(), "): ")
				decodeRowSet(addedRowsIncluded)

				fmt.Print("removed rows (", updateMeta.RemovedRowsLength(), "): ")
				decodeRowSet(removedRows)

				fmt.Print("shift data (", updateMeta.ShiftDataLength(), "): ")
				decodeRowSet(shiftData)

				fmt.Println(updateMeta.ShiftDataLength())

			}
		}
	}

	return nil, err
}

// Close closes the flight stub and frees any associated resources.
// The flight stub should not be used after calling this function.
// The client lock should be held when calling this function.
func (fs *flightStub) Close() error {
	if fs.stub != nil {
		err := fs.stub.Close()
		if err != nil {
			return err
		}
		fs.stub = nil
	}
	return nil
}
