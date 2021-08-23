package rasbinary

import (
	messagesv1 "github.com/v8platform/protos/gen/ras/messages/v1"
	protocolv1 "github.com/v8platform/protos/gen/ras/protocol/v1"
	"google.golang.org/protobuf/proto"
	pref "google.golang.org/protobuf/reflect/protoreflect"
	"time"
)

const (
	UTF8_CHARSET   = "UTF-8"
	SIZEOF_SHORT   = 2
	SIZEOF_INT     = 4
	SIZEOF_LONG    = 8
	NULL_BYTE      = 0x80
	TRUE_BYTE      = 1
	FALSE_BYTE     = 0
	MAX_SHIFT      = 7
	NULL_SHIFT     = 6
	BYTE_MASK      = 255
	NEXT_MASK      = -128
	NULL_NEXT_MASK = 64
	LAST_MASK      = 0
	NULL_LSB_MASK  = 63
	LSB_MASK       = 127
	TEMP_CAPACITY  = 256
)

const AgeDelta = 621355968000000

func dateFromTicks(ticks int64) time.Time {
	if ticks > 0 {

		timeT := (ticks - AgeDelta) / 10

		t := time.Unix(0, timeT*int64(time.Millisecond))

		return t

	}
	return time.Time{}
}

func dateToTicks(date time.Time) (ticks int64) {

	if !date.IsZero() {

		ticks = date.UnixNano() / int64(time.Millisecond)

		ticks = ticks*10 + AgeDelta

		return ticks

	}
	return 0
}

func getFieldValueOfNumber(m pref.Message, number int32) interface{} {
	fd := m.Descriptor().Fields().ByNumber(pref.FieldNumber(number))
	return m.Get(fd).Interface()
}

func GetMessageType(m proto.Message) (messagesv1.MessageType, bool) {

	val, ok := getExtension(m, messagesv1.E_MessageType)
	return val.(messagesv1.MessageType), ok

}

func GetPacketType(m proto.Message) (protocolv1.PacketType, bool) {

	val, ok := getExtension(m, protocolv1.E_PacketType)
	return val.(protocolv1.PacketType), ok

}

func GetEndpointMessageType(m proto.Message) (protocolv1.PacketType, bool) {

	val, ok := getExtension(m, protocolv1.E_EndpointDataType)
	return val.(protocolv1.PacketType), ok

}

func getExtension(m proto.Message, ext pref.ExtensionType) (interface{}, bool) {
	opts := m.ProtoReflect().Descriptor().Options()

	if !proto.HasExtension(opts, ext) {
		return nil, false
	}

	val := proto.GetExtension(opts, ext)

	return val, true
}
