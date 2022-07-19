// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package flatbuf

import "strconv"

type MetadataVersion int16

const (
	/// 0.1.0 (October 2016).
	MetadataVersionV1 MetadataVersion = 0
	/// 0.2.0 (February 2017). Non-backwards compatible with V1.
	MetadataVersionV2 MetadataVersion = 1
	/// 0.3.0 -> 0.7.1 (May - December 2017). Non-backwards compatible with V2.
	MetadataVersionV3 MetadataVersion = 2
	/// >= 0.8.0 (December 2017). Non-backwards compatible with V3.
	MetadataVersionV4 MetadataVersion = 3
	/// >= 1.0.0 (July 2020. Backwards compatible with V4 (V5 readers can read V4
	/// metadata and IPC messages). Implementations are recommended to provide a
	/// V4 compatibility mode with V5 format changes disabled.
	///
	/// Incompatible changes between V4 and V5:
	/// - Union buffer layout has changed. In V5, Unions don't have a validity
	///   bitmap buffer.
	MetadataVersionV5 MetadataVersion = 4
)

var EnumNamesMetadataVersion = map[MetadataVersion]string{
	MetadataVersionV1: "V1",
	MetadataVersionV2: "V2",
	MetadataVersionV3: "V3",
	MetadataVersionV4: "V4",
	MetadataVersionV5: "V5",
}

var EnumValuesMetadataVersion = map[string]MetadataVersion{
	"V1": MetadataVersionV1,
	"V2": MetadataVersionV2,
	"V3": MetadataVersionV3,
	"V4": MetadataVersionV4,
	"V5": MetadataVersionV5,
}

func (v MetadataVersion) String() string {
	if s, ok := EnumNamesMetadataVersion[v]; ok {
		return s
	}
	return "MetadataVersion(" + strconv.FormatInt(int64(v), 10) + ")"
}
