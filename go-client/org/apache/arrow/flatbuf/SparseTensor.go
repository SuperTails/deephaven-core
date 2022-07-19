// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package flatbuf

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type SparseTensor struct {
	_tab flatbuffers.Table
}

func GetRootAsSparseTensor(buf []byte, offset flatbuffers.UOffsetT) *SparseTensor {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &SparseTensor{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *SparseTensor) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *SparseTensor) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *SparseTensor) TypeType() Type {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return Type(rcv._tab.GetByte(o + rcv._tab.Pos))
	}
	return 0
}

func (rcv *SparseTensor) MutateTypeType(n Type) bool {
	return rcv._tab.MutateByteSlot(4, byte(n))
}

/// The type of data contained in a value cell.
/// Currently only fixed-width value types are supported,
/// no strings or nested types.
func (rcv *SparseTensor) Type(obj *flatbuffers.Table) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		rcv._tab.Union(obj, o)
		return true
	}
	return false
}

/// The type of data contained in a value cell.
/// Currently only fixed-width value types are supported,
/// no strings or nested types.
/// The dimensions of the tensor, optionally named.
func (rcv *SparseTensor) Shape(obj *TensorDim, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *SparseTensor) ShapeLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

/// The dimensions of the tensor, optionally named.
/// The number of non-zero values in a sparse tensor.
func (rcv *SparseTensor) NonZeroLength() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

/// The number of non-zero values in a sparse tensor.
func (rcv *SparseTensor) MutateNonZeroLength(n int64) bool {
	return rcv._tab.MutateInt64Slot(10, n)
}

func (rcv *SparseTensor) SparseIndexType() SparseTensorIndex {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return SparseTensorIndex(rcv._tab.GetByte(o + rcv._tab.Pos))
	}
	return 0
}

func (rcv *SparseTensor) MutateSparseIndexType(n SparseTensorIndex) bool {
	return rcv._tab.MutateByteSlot(12, byte(n))
}

/// Sparse tensor index
func (rcv *SparseTensor) SparseIndex(obj *flatbuffers.Table) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		rcv._tab.Union(obj, o)
		return true
	}
	return false
}

/// Sparse tensor index
/// The location and size of the tensor's data
func (rcv *SparseTensor) Data(obj *Buffer) *Buffer {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		x := o + rcv._tab.Pos
		if obj == nil {
			obj = new(Buffer)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

/// The location and size of the tensor's data
func SparseTensorStart(builder *flatbuffers.Builder) {
	builder.StartObject(7)
}
func SparseTensorAddTypeType(builder *flatbuffers.Builder, typeType Type) {
	builder.PrependByteSlot(0, byte(typeType), 0)
}
func SparseTensorAddType(builder *flatbuffers.Builder, type_ flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(type_), 0)
}
func SparseTensorAddShape(builder *flatbuffers.Builder, shape flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(shape), 0)
}
func SparseTensorStartShapeVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func SparseTensorAddNonZeroLength(builder *flatbuffers.Builder, nonZeroLength int64) {
	builder.PrependInt64Slot(3, nonZeroLength, 0)
}
func SparseTensorAddSparseIndexType(builder *flatbuffers.Builder, sparseIndexType SparseTensorIndex) {
	builder.PrependByteSlot(4, byte(sparseIndexType), 0)
}
func SparseTensorAddSparseIndex(builder *flatbuffers.Builder, sparseIndex flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(5, flatbuffers.UOffsetT(sparseIndex), 0)
}
func SparseTensorAddData(builder *flatbuffers.Builder, data flatbuffers.UOffsetT) {
	builder.PrependStructSlot(6, flatbuffers.UOffsetT(data), 0)
}
func SparseTensorEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
