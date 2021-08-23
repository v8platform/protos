package extra

import (
	"fmt"
	"github.com/v8platform/protos/extra/encoding/rasbinary"
	messagesv1 "github.com/v8platform/protos/gen/ras/messages/v1"
	protocolv1 "github.com/v8platform/protos/gen/ras/protocol/v1"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"io"
)

func UnpackPacketDataNew(p *protocolv1.Packet) (proto.Message, error) {

	var packetType protoreflect.MessageType

	protoregistry.GlobalTypes.RangeMessages(func(messageType protoreflect.MessageType) bool {

		isPacketMessage := proto.HasExtension(messageType.Descriptor().Options(), protocolv1.E_PacketType)

		if !isPacketMessage {
			return true
		}

		mPacketType := proto.GetExtension(messageType.Descriptor().Options(), protocolv1.E_PacketType).(protocolv1.PacketType)

		if mPacketType == p.Type {
			packetType = messageType
			return false
		}

		return true
	})
	val := packetType.New().Interface()

	err := rasbinary.Unmarshal(p.Data, val)
	if err != nil {
		return nil, err
	}
	return val, nil
}

func UnpackEndpointMessageNew(m *protocolv1.EndpointMessage, version int32) (proto.Message, error) {

	if err := m.GetFailure(); err != nil {
		return nil, fmt.Errorf(err.String())
	}

	if void := m.GetVoidMessage(); void != nil {
		return void, nil
	}

	messageData := m.GetMessage()
	if messageData == nil {
		return messageData, nil
	}

	var packetType protoreflect.MessageType

	protoregistry.GlobalTypes.RangeMessages(func(messageType protoreflect.MessageType) bool {

		isPacketMessage := proto.HasExtension(messageType.Descriptor().Options(), messagesv1.E_MessageType)

		if !isPacketMessage {
			return true
		}

		mPacketType := proto.GetExtension(messageType.Descriptor().Options(), messagesv1.E_MessageType).(messagesv1.MessageType)

		if mPacketType == messageData.Type {
			packetType = messageType
			return false
		}

		return true
	})

	val := packetType.New().Interface()

	err := UnpackMessageDataTo(messageData, val, version)
	if err != nil {
		return nil, err
	}

	return val, nil
}

func UnpackMessageDataTo(p *protocolv1.EndpointDataMessage, m proto.Message, v int32) error {

	// TODO Проверка на соответсвие типов
	enc := rasbinary.UnmarshalOptions{ProtocolVersion: v}
	err := enc.Unmarshal(p.Bytes, m)
	if err != nil {
		return err
	}

	return nil

}

func UnpackPacketDataTo(p *protocolv1.Packet, m proto.Message) error {

	md := m.ProtoReflect().Descriptor()

	isPacketMessage := proto.HasExtension(md.Options(), protocolv1.E_PacketType)

	if !isPacketMessage {
		return fmt.Errorf("this is not packet message: <%s>", proto.MessageName(m))
	}

	packetType := proto.GetExtension(md.Options(), protocolv1.E_PacketType).(protocolv1.PacketType)

	switch p.Type {
	case packetType:
		err := rasbinary.Unmarshal(p.Data, m)
		if err != nil {
			return err
		}
	case protocolv1.PacketType_PACKET_TYPE_ENDPOINT_FAILURE:

		var failure protocolv1.EndpointFailureAck
		err := rasbinary.Unmarshal(p.Data, &failure)
		if err != nil {
			return err
		}
		return fmt.Errorf("endpoint error <%s>", &failure)
	default:
		return fmt.Errorf("this is not packet message: <%s>", proto.MessageName(m))
	}

	return nil

}

func NewPacket(m proto.Message) (*protocolv1.Packet, error) {

	md := m.ProtoReflect().Descriptor()

	isPacketMessage := proto.HasExtension(md.Options(), protocolv1.E_PacketType)

	if !isPacketMessage {
		return nil, fmt.Errorf("this is not packet message: <%s>", proto.MessageName(m))
	}

	packetType := proto.GetExtension(md.Options(), protocolv1.E_PacketType).(protocolv1.PacketType)
	bytes, err := rasbinary.Marshal(m)

	if err != nil {
		return nil, err
	}
	packet := &protocolv1.Packet{
		Type: packetType,
		Size: int32(len(bytes)),
		Data: bytes,
	}

	return packet, nil
}

func WritePacketMessage(writer io.Writer, m proto.Message) (int, error) {

	p, err := NewPacket(m)

	if err != nil {
		return 0, err
	}
	return WritePacket(writer, p)
}

func WritePacket(writer io.Writer, p *protocolv1.Packet) (int, error) {

	b, err := rasbinary.Marshal(p)
	if err != nil {
		return 0, err
	}
	return writer.Write(b)
}

func ReadPacket(reader io.Reader) (*protocolv1.Packet, error) {

	var packet protocolv1.Packet

	u := rasbinary.UnmarshalOptions{}

	err := u.UnmarshalReader(reader, &packet)
	if err != nil {
		return nil, err
	}

	return &packet, nil
}

func ReadPacketMessage(reader io.Reader, m proto.Message) error {

	packet, err := ReadPacket(reader)
	if err != nil {
		return err
	}

	err = UnpackPacketDataTo(packet, m)
	if err != nil {
		return err
	}

	return nil
}
