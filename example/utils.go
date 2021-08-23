package main

import (
	"fmt"
	"github.com/v8platform/protos/extra/encoding/rasbinary"
	protocolv1 "github.com/v8platform/protos/gen/ras/protocol/v1"
	"google.golang.org/protobuf/proto"
	"io"
)

func writePacket(writer io.Writer, packet *protocolv1.Packet) (int, error) {

	packetBytes, err := rasbinary.Marshal(packet)

	if err != nil {
		return 0, err
	}

	return writer.Write(packetBytes)
}

func ReadPacket(reader io.Reader, unmarshalData bool) (*protocolv1.Packet, error) {

	var packet protocolv1.Packet

	u := rasbinary.UnmarshalOptions{}

	err := u.UnmarshalReader(reader, &packet)

	// pp.Println(packet)

	return &packet, err

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
