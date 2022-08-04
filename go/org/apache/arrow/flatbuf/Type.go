// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package flatbuf

import "strconv"

/// ----------------------------------------------------------------------
/// Top-level Type value, enabling extensible type-specific metadata. We can
/// add new logical types to Type without breaking backwards compatibility
type Type byte

const (
	TypeNONE            Type = 0
	TypeNull            Type = 1
	TypeInt             Type = 2
	TypeFloatingPoint   Type = 3
	TypeBinary          Type = 4
	TypeUtf8            Type = 5
	TypeBool            Type = 6
	TypeDecimal         Type = 7
	TypeDate            Type = 8
	TypeTime            Type = 9
	TypeTimestamp       Type = 10
	TypeInterval        Type = 11
	TypeList            Type = 12
	TypeStruct_         Type = 13
	TypeUnion           Type = 14
	TypeFixedSizeBinary Type = 15
	TypeFixedSizeList   Type = 16
	TypeMap             Type = 17
	TypeDuration        Type = 18
	TypeLargeBinary     Type = 19
	TypeLargeUtf8       Type = 20
	TypeLargeList       Type = 21
)

var EnumNamesType = map[Type]string{
	TypeNONE:            "NONE",
	TypeNull:            "Null",
	TypeInt:             "Int",
	TypeFloatingPoint:   "FloatingPoint",
	TypeBinary:          "Binary",
	TypeUtf8:            "Utf8",
	TypeBool:            "Bool",
	TypeDecimal:         "Decimal",
	TypeDate:            "Date",
	TypeTime:            "Time",
	TypeTimestamp:       "Timestamp",
	TypeInterval:        "Interval",
	TypeList:            "List",
	TypeStruct_:         "Struct_",
	TypeUnion:           "Union",
	TypeFixedSizeBinary: "FixedSizeBinary",
	TypeFixedSizeList:   "FixedSizeList",
	TypeMap:             "Map",
	TypeDuration:        "Duration",
	TypeLargeBinary:     "LargeBinary",
	TypeLargeUtf8:       "LargeUtf8",
	TypeLargeList:       "LargeList",
}

var EnumValuesType = map[string]Type{
	"NONE":            TypeNONE,
	"Null":            TypeNull,
	"Int":             TypeInt,
	"FloatingPoint":   TypeFloatingPoint,
	"Binary":          TypeBinary,
	"Utf8":            TypeUtf8,
	"Bool":            TypeBool,
	"Decimal":         TypeDecimal,
	"Date":            TypeDate,
	"Time":            TypeTime,
	"Timestamp":       TypeTimestamp,
	"Interval":        TypeInterval,
	"List":            TypeList,
	"Struct_":         TypeStruct_,
	"Union":           TypeUnion,
	"FixedSizeBinary": TypeFixedSizeBinary,
	"FixedSizeList":   TypeFixedSizeList,
	"Map":             TypeMap,
	"Duration":        TypeDuration,
	"LargeBinary":     TypeLargeBinary,
	"LargeUtf8":       TypeLargeUtf8,
	"LargeList":       TypeLargeList,
}

func (v Type) String() string {
	if s, ok := EnumNamesType[v]; ok {
		return s
	}
	return "Type(" + strconv.FormatInt(int64(v), 10) + ")"
}
