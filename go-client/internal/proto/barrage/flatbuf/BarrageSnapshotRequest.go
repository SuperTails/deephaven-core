// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package flatbuf

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// Describes the snapshot the client would like to acquire.
type BarrageSnapshotRequest struct {
	_tab flatbuffers.Table
}

func GetRootAsBarrageSnapshotRequest(buf []byte, offset flatbuffers.UOffsetT) *BarrageSnapshotRequest {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &BarrageSnapshotRequest{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *BarrageSnapshotRequest) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *BarrageSnapshotRequest) Table() flatbuffers.Table {
	return rcv._tab
}

/// Ticket for the source data set.
func (rcv *BarrageSnapshotRequest) Ticket(j int) int8 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetInt8(a + flatbuffers.UOffsetT(j*1))
	}
	return 0
}

func (rcv *BarrageSnapshotRequest) TicketLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

/// Ticket for the source data set.
func (rcv *BarrageSnapshotRequest) MutateTicket(j int, n int8) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateInt8(a+flatbuffers.UOffsetT(j*1), n)
	}
	return false
}

/// The bitset of columns to request. If not provided then all columns are requested.
func (rcv *BarrageSnapshotRequest) Columns(j int) int8 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetInt8(a + flatbuffers.UOffsetT(j*1))
	}
	return 0
}

func (rcv *BarrageSnapshotRequest) ColumnsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

/// The bitset of columns to request. If not provided then all columns are requested.
func (rcv *BarrageSnapshotRequest) MutateColumns(j int, n int8) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateInt8(a+flatbuffers.UOffsetT(j*1), n)
	}
	return false
}

/// This is an encoded and compressed RowSet in position-space to subscribe to. If not provided then the entire
/// table is requested.
func (rcv *BarrageSnapshotRequest) Viewport(j int) int8 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetInt8(a + flatbuffers.UOffsetT(j*1))
	}
	return 0
}

func (rcv *BarrageSnapshotRequest) ViewportLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

/// This is an encoded and compressed RowSet in position-space to subscribe to. If not provided then the entire
/// table is requested.
func (rcv *BarrageSnapshotRequest) MutateViewport(j int, n int8) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateInt8(a+flatbuffers.UOffsetT(j*1), n)
	}
	return false
}

/// Options to configure your subscription.
func (rcv *BarrageSnapshotRequest) SnapshotOptions(obj *BarrageSnapshotOptions) *BarrageSnapshotOptions {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(BarrageSnapshotOptions)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

/// Options to configure your subscription.
/// When this is set the viewport RowSet will be inverted against the length of the table. That is to say
/// every position value is converted from `i` to `n - i - 1` if the table has `n` rows.
func (rcv *BarrageSnapshotRequest) ReverseViewport() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

/// When this is set the viewport RowSet will be inverted against the length of the table. That is to say
/// every position value is converted from `i` to `n - i - 1` if the table has `n` rows.
func (rcv *BarrageSnapshotRequest) MutateReverseViewport(n bool) bool {
	return rcv._tab.MutateBoolSlot(12, n)
}

func BarrageSnapshotRequestStart(builder *flatbuffers.Builder) {
	builder.StartObject(5)
}
func BarrageSnapshotRequestAddTicket(builder *flatbuffers.Builder, ticket flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(ticket), 0)
}
func BarrageSnapshotRequestStartTicketVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(1, numElems, 1)
}
func BarrageSnapshotRequestAddColumns(builder *flatbuffers.Builder, columns flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(columns), 0)
}
func BarrageSnapshotRequestStartColumnsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(1, numElems, 1)
}
func BarrageSnapshotRequestAddViewport(builder *flatbuffers.Builder, viewport flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(viewport), 0)
}
func BarrageSnapshotRequestStartViewportVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(1, numElems, 1)
}
func BarrageSnapshotRequestAddSnapshotOptions(builder *flatbuffers.Builder, snapshotOptions flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(snapshotOptions), 0)
}
func BarrageSnapshotRequestAddReverseViewport(builder *flatbuffers.Builder, reverseViewport bool) {
	builder.PrependBoolSlot(4, reverseViewport, false)
}
func BarrageSnapshotRequestEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
