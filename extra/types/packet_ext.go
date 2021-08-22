package types

import (
	"fmt"
	"github.com/v8platform/protos/extra/encoding/rasbinary"
	messagesv1 "github.com/v8platform/protos/gen/ras/messages/v1"
	protocolv1 "github.com/v8platform/protos/gen/ras/protocol/v1"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
)

type Endpoint struct {
	*protocolv1.EndpointOpenAck
}

func (e *Endpoint) Message(m proto.Message) (*protocolv1.EndpointMessage, error) {

	switch t := m.(type) {

	case *protocolv1.Packet:
		switch t.Type {
		case protocolv1.PacketType_PACKET_TYPE_ENDPOINT_MESSAGE:
			var data protocolv1.EndpointMessage
			err := UnpackPacketDataTo(t, &data)
			if err != nil {
				return nil, err
			}
			return &data, nil
		case protocolv1.PacketType_PACKET_TYPE_ENDPOINT_FAILURE:

			var failure protocolv1.EndpointFailureAck
			err := UnpackPacketDataTo(t, &failure)
			if err != nil {
				return nil, err
			}
			return nil, fmt.Errorf(failure.String())
		}
	}

	isTypeMessage := proto.HasExtension(m.ProtoReflect().Descriptor().Options(), messagesv1.E_MessageType)

	if !isTypeMessage {

		messageType := proto.GetExtension(m.ProtoReflect().Descriptor().Options(), messagesv1.E_MessageType).(messagesv1.MessageType)

		enc := rasbinary.MarshalOptions{ProtocolVersion: e.Version}
		bytesData, err := enc.Marshal(m)
		if err != nil {
			return nil, err
		}
		return &protocolv1.EndpointMessage{
			EndpointId: e.EndpointId,
			Format:     255,
			Type:       protocolv1.EndpointDataType_ENDPOINT_DATA_TYPE_MESSAGE,
			Data: &protocolv1.EndpointMessage_Message{Message: &protocolv1.EndpointDataMessage{
				Type:  messageType,
				Bytes: bytesData,
			}},
		}, nil
	}

	return nil, fmt.Errorf("no type message")
}

func NewEndpointFromPacket(p *protocolv1.Packet) (*Endpoint, error) {

	var openAck protocolv1.EndpointOpenAck

	err := UnpackPacketDataTo(p, &openAck)
	if err != nil {
		return nil, err
	}
	return NewEndpoint(&openAck)
}

func NewEndpoint(p *protocolv1.EndpointOpenAck) (*Endpoint, error) {

	return &Endpoint{
		p,
	}, nil
}

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

	if p.Type != packetType {
		return fmt.Errorf("this is not packet message: <%s>", proto.MessageName(m))
	}

	err := rasbinary.Unmarshal(p.Data, m)
	if err != nil {
		return err
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
