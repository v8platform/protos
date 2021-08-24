package rasbinary

import (
	"fmt"
	"github.com/v8platform/protos/gen/ras/encoding"
	extpb "github.com/v8platform/protos/gen/ras/encoding"
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

	if val, ok := getExtension(m.ProtoReflect().Descriptor().Options(), messagesv1.E_MessageType); ok {
		return val.(messagesv1.MessageType), true
	}
	return 0, false
}

func GetPacketType(m proto.Message) (protocolv1.PacketType, bool) {

	if val, ok := getExtension(m.ProtoReflect().Descriptor().Options(), protocolv1.E_PacketType); ok {
		return val.(protocolv1.PacketType), true
	}
	return 0, false

}

func GetEncodingFieldOptions(fd pref.FieldDescriptor) (*extpb.EncodingFieldOptions, bool) {

	if val, ok := getExtension(fd.Options(), encoding.E_Field); ok {
		return val.(*extpb.EncodingFieldOptions), true
	}
	return nil, false

}

func GetEndpointDataType(m proto.Message) (protocolv1.EndpointDataType, bool) {

	if val, ok := getExtension(m.ProtoReflect().Descriptor().Options(), protocolv1.E_EndpointDataType); ok {
		return val.(protocolv1.EndpointDataType), true
	}
	return 0, false

}

func getExtension(opts pref.ProtoMessage, ext pref.ExtensionType) (interface{}, bool) {

	if !proto.HasExtension(opts, ext) {
		return nil, false
	}

	val := proto.GetExtension(opts, ext)

	return val, true
}

type unmarshalFunc func(d decoder, message pref.Message, opts *extpb.EncodingFieldOptions) error

func unmarshalTimestamp(d decoder, message pref.Message, opts *extpb.EncodingFieldOptions) error {
	var decodeFn TypeDecoderFunc
	var err error
	if opts != nil {
		decodeFn, err = GetDecodeFunc(opts.GetEncoder())
		if err != nil {
			return err
		}
	}
	if decodeFn == nil {
		decodeFn = decodeTime
	}

	var val time.Time

	_, err = decodeFn(d, &val, opts.Opts)
	if err != nil {
		return err
	}

	secs := val.Unix()
	if secs < minTimestampSeconds || secs > maxTimestampSeconds {
		return fmt.Errorf("%v value out of range: %v", Timestamp_message_fullname, val)
	}

	fds := message.Descriptor().Fields()
	fdSeconds := fds.ByNumber(pref.FieldNumber(1))
	fdNanos := fds.ByNumber(pref.FieldNumber(2))

	message.Set(fdSeconds, pref.ValueOfInt64(secs))
	message.Set(fdNanos, pref.ValueOfInt32(int32(val.Nanosecond())))

	return nil
}

// Names for google.protobuf.Timestamp.
const (
	Timestamp_message_name     pref.Name     = "Timestamp"
	Timestamp_message_fullname pref.FullName = "google.protobuf.Timestamp"
)

const (
	maxTimestampSeconds = 253402300799
	minTimestampSeconds = -62135596800
)

// wellKnownTypeUnmarshaler returns a unmarshal function if the message type
// has specialized serialization behavior. It returns nil otherwise.
func wellKnownTypeUnmarshaler(name pref.FullName) unmarshalFunc {
	switch name {
	// case genid.Any_message_name:
	// 	return decoder.unmarshalAny
	case Timestamp_message_fullname:
		return unmarshalTimestamp
		// 	case genid.Duration_message_name:
		// 		return decoder.unmarshalDuration
		// 	case genid.BoolValue_message_name,
		// 		genid.Int32Value_message_name,
		// 		genid.Int64Value_message_name,
		// 		genid.UInt32Value_message_name,
		// 		genid.UInt64Value_message_name,
		// 		genid.FloatValue_message_name,
		// 		genid.DoubleValue_message_name,
		// 		genid.StringValue_message_name,
		// 		genid.BytesValue_message_name:
		// 		return decoder.unmarshalWrapperType
		// 	case genid.Struct_message_name:
		// 		return decoder.unmarshalStruct
		// 	case genid.ListValue_message_name:
		// 		return decoder.unmarshalListValue
		// 	case genid.Value_message_name:
		// 		return decoder.unmarshalKnownValue
		// 	case genid.FieldMask_message_name:
		// 		return decoder.unmarshalFieldMask
		// 	case genid.Empty_message_name:
		// 		return decoder.unmarshalEmpty
	}
	return nil
}
