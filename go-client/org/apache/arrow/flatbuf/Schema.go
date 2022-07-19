// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package flatbuf

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// ----------------------------------------------------------------------
/// A Schema describes the columns in a row batch
type Schema struct {
	_tab flatbuffers.Table
}

func GetRootAsSchema(buf []byte, offset flatbuffers.UOffsetT) *Schema {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Schema{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *Schema) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Schema) Table() flatbuffers.Table {
	return rcv._tab
}

/// endianness of the buffer
/// it is Little Endian by default
/// if endianness doesn't match the underlying system then the vectors need to be converted
func (rcv *Schema) Endianness() Endianness {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return Endianness(rcv._tab.GetInt16(o + rcv._tab.Pos))
	}
	return 0
}

/// endianness of the buffer
/// it is Little Endian by default
/// if endianness doesn't match the underlying system then the vectors need to be converted
func (rcv *Schema) MutateEndianness(n Endianness) bool {
	return rcv._tab.MutateInt16Slot(4, int16(n))
}

func (rcv *Schema) Fields(obj *Field, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *Schema) FieldsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *Schema) CustomMetadata(obj *KeyValue, j int) bool {
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

func (rcv *Schema) CustomMetadataLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

/// Features used in the stream/file.
func (rcv *Schema) Features(j int) Feature {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return Feature(rcv._tab.GetInt64(a + flatbuffers.UOffsetT(j*8)))
	}
	return 0
}

func (rcv *Schema) FeaturesLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

/// Features used in the stream/file.
func (rcv *Schema) MutateFeatures(j int, n Feature) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateInt64(a+flatbuffers.UOffsetT(j*8), int64(n))
	}
	return false
}

func SchemaStart(builder *flatbuffers.Builder) {
	builder.StartObject(4)
}
func SchemaAddEndianness(builder *flatbuffers.Builder, endianness Endianness) {
	builder.PrependInt16Slot(0, int16(endianness), 0)
}
func SchemaAddFields(builder *flatbuffers.Builder, fields flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(fields), 0)
}
func SchemaStartFieldsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func SchemaAddCustomMetadata(builder *flatbuffers.Builder, customMetadata flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(customMetadata), 0)
}
func SchemaStartCustomMetadataVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func SchemaAddFeatures(builder *flatbuffers.Builder, features flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(features), 0)
}
func SchemaStartFeaturesVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(8, numElems, 8)
}
func SchemaEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
