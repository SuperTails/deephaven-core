// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package flatbuf

import "strconv"

type BarrageMessageType int8

const (
	/// A barrage message wrapper might send a None message type
	/// if the msg_payload is empty.
	BarrageMessageTypeNone                        BarrageMessageType = 0
	/// for session management (not-yet-used)
	BarrageMessageTypeNewSessionRequest           BarrageMessageType = 1
	BarrageMessageTypeRefreshSessionRequest       BarrageMessageType = 2
	BarrageMessageTypeSessionInfoResponse         BarrageMessageType = 3
	/// for subscription parsing/management (aka DoPut, DoExchange)
	BarrageMessageTypeBarrageSerializationOptions BarrageMessageType = 4
	BarrageMessageTypeBarrageSubscriptionRequest  BarrageMessageType = 5
	BarrageMessageTypeBarrageUpdateMetadata       BarrageMessageType = 6
	BarrageMessageTypeBarrageSnapshotRequest      BarrageMessageType = 7
	BarrageMessageTypeBarragePublicationRequest   BarrageMessageType = 8
)

var EnumNamesBarrageMessageType = map[BarrageMessageType]string{
	BarrageMessageTypeNone:                        "None",
	BarrageMessageTypeNewSessionRequest:           "NewSessionRequest",
	BarrageMessageTypeRefreshSessionRequest:       "RefreshSessionRequest",
	BarrageMessageTypeSessionInfoResponse:         "SessionInfoResponse",
	BarrageMessageTypeBarrageSerializationOptions: "BarrageSerializationOptions",
	BarrageMessageTypeBarrageSubscriptionRequest:  "BarrageSubscriptionRequest",
	BarrageMessageTypeBarrageUpdateMetadata:       "BarrageUpdateMetadata",
	BarrageMessageTypeBarrageSnapshotRequest:      "BarrageSnapshotRequest",
	BarrageMessageTypeBarragePublicationRequest:   "BarragePublicationRequest",
}

var EnumValuesBarrageMessageType = map[string]BarrageMessageType{
	"None":                        BarrageMessageTypeNone,
	"NewSessionRequest":           BarrageMessageTypeNewSessionRequest,
	"RefreshSessionRequest":       BarrageMessageTypeRefreshSessionRequest,
	"SessionInfoResponse":         BarrageMessageTypeSessionInfoResponse,
	"BarrageSerializationOptions": BarrageMessageTypeBarrageSerializationOptions,
	"BarrageSubscriptionRequest":  BarrageMessageTypeBarrageSubscriptionRequest,
	"BarrageUpdateMetadata":       BarrageMessageTypeBarrageUpdateMetadata,
	"BarrageSnapshotRequest":      BarrageMessageTypeBarrageSnapshotRequest,
	"BarragePublicationRequest":   BarrageMessageTypeBarragePublicationRequest,
}

func (v BarrageMessageType) String() string {
	if s, ok := EnumNamesBarrageMessageType[v]; ok {
		return s
	}
	return "BarrageMessageType(" + strconv.FormatInt(int64(v), 10) + ")"
}
