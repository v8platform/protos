package extra

import (
	"fmt"
	"github.com/v8platform/protos/extra/encoding/rasbinary"
	messagesv1 "github.com/v8platform/protos/gen/ras/messages/v1"
	protocolv1 "github.com/v8platform/protos/gen/ras/protocol/v1"
	"google.golang.org/protobuf/proto"
	"io"
	"net"
	"regexp"
	"strconv"
	"strings"
)

const protocolVersion = 256

var serviceVersions = []string{"3.0", "4.0", "5.0", "6.0", "7.0", "8.0", "9.0", "10.0"}

type Endpoint struct {
	Format         int32
	ServiceVersion int32
	*protocolv1.EndpointOpenAck
}

func (e *Endpoint) NewMessage(m proto.Message) (*protocolv1.EndpointMessage, error) {

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

		default:
			return nil, fmt.Errorf("unknown packet type <%s>", t.Type.String())
		}
	}

	isTypeMessage := proto.HasExtension(m.ProtoReflect().Descriptor().Options(), messagesv1.E_MessageType)

	if isTypeMessage {

		messageType := proto.GetExtension(m.ProtoReflect().Descriptor().Options(), messagesv1.E_MessageType).(messagesv1.MessageType)

		enc := rasbinary.MarshalOptions{ProtocolVersion: e.ServiceVersion}
		bytesData, err := enc.Marshal(m)
		if err != nil {
			return nil, err
		}
		return &protocolv1.EndpointMessage{
			EndpointId: e.EndpointId,
			Format:     e.Format,
			Type:       protocolv1.EndpointDataType_ENDPOINT_DATA_TYPE_MESSAGE,
			Data: &protocolv1.EndpointMessage_Message{Message: &protocolv1.EndpointDataMessage{
				Type:  messageType,
				Bytes: bytesData,
			}},
		}, nil
	}

	return nil, fmt.Errorf("no type message")
}

func (e *Endpoint) SendMessage(writer io.Writer, m proto.Message) error {

	em, err := e.NewMessage(m)
	if err != nil {
		return err
	}

	_, err = WritePacketMessage(writer, em)
	if err != nil {
		return err
	}

	return nil
}

func (e *Endpoint) ReadMessage(reader io.Reader, m proto.Message) error {

	p, err := ReadPacket(reader)
	if err != nil {
		return err
	}
	em, err := e.NewMessage(p)
	if err != nil {
		return err
	}

	return e.UnpackMessage(em, m)
}

func (e *Endpoint) UnpackMessage(em *protocolv1.EndpointMessage, to proto.Message) error {

	if err := em.GetFailure(); err != nil {
		return fmt.Errorf(err.String())
	}

	if void := em.GetVoidMessage(); void != nil {
		return fmt.Errorf("can not unpack void message")
	}
	isPacketMessage := proto.HasExtension(to.ProtoReflect().Descriptor().Options(), messagesv1.E_MessageType)

	if !isPacketMessage {
		return fmt.Errorf("can not unpack message unknown type")
	}

	mType := proto.GetExtension(to.ProtoReflect().Descriptor().Options(), messagesv1.E_MessageType).(messagesv1.MessageType)
	messageData := em.GetMessage()
	if mType != messageData.Type {
		return fmt.Errorf("unpack message type <%s> to <%s> mismatch ", messageData.String(), mType.String())
	}
	enc := rasbinary.UnmarshalOptions{ProtocolVersion: e.ServiceVersion}
	err := enc.Unmarshal(messageData.GetBytes(), to)
	if err != nil {
		return err
	}

	return nil
}

func NewEndpointFromPacket(p *protocolv1.Packet) (*Endpoint, error) {

	var openAck protocolv1.EndpointOpenAck

	err := UnpackPacketDataTo(p, &openAck)
	if err != nil {
		return nil, err
	}
	return NewEndpoint(&openAck)
}

func OpenEndpoint(conn net.Conn, version string) (*Endpoint, error) {

open:
	_, err := WritePacketMessage(conn, &protocolv1.EndpointOpen{
		Service: "v8.service.Admin.Cluster",
		Version: version,
	})

	if err != nil {
		return nil, err
	}

read:
	var openAck protocolv1.EndpointOpenAck
	p, err := ReadPacket(conn)

	switch p.Type {
	case protocolv1.PacketType_PACKET_TYPE_ENDPOINT_FAILURE:
		var failure protocolv1.EndpointFailureAck
		err := UnpackPacketDataTo(p, &failure)
		if err != nil {
			return nil, err
		}
		supportedVersion := detectSupportedVersion(&failure)
		if len(supportedVersion) == 0 {
			return nil, fmt.Errorf("endpoint open error <%s>", &failure)
		}
		version = supportedVersion

		goto open

	case protocolv1.PacketType_PACKET_TYPE_ENDPOINT_OPEN_ACK:
		err := UnpackPacketDataTo(p, &openAck)
		if err != nil {
			return nil, err
		}
	case protocolv1.PacketType_PACKET_TYPE_KEEP_ALIVE:
		goto read
	}

	if err != nil {
		return nil, err
	}

	return NewEndpoint(&openAck)
}

func Connect(conn net.Conn) error {

	negotationPacket, err := rasbinary.Marshal(&protocolv1.NegotiateMessage{
		Magic:    475223888, // Константы
		Protocol: 256,       // Константы
		Version:  256,       // Константы
	})

	if err != nil {
		return err
	}
	_, err = conn.Write(negotationPacket)
	if err != nil {
		return err
	}

	_, err = WritePacketMessage(conn, &protocolv1.ConnectMessage{
		// Params: map[string]protocolv1.Param{
		// 	"connect.timeout": {},
		// },
	})
	if err != nil {
		return err
	}
	var answer protocolv1.ConnectMessageAck

	err = ReadPacketMessage(conn, &answer)

	if err != nil {
		return err
	}

	return nil
}

func NewEndpoint(p *protocolv1.EndpointOpenAck) (*Endpoint, error) {

	serviceVersion, _ := strconv.ParseInt(p.Version, 10, 64)

	return &Endpoint{
		0,
		int32(serviceVersion),
		p,
	}, nil
}

func detectSupportedVersion(fail *protocolv1.EndpointFailureAck) string {

	if fail.Cause == nil {
		return ""
	}

	msg := fail.Cause.Message

	// log.Print(fail.String())

	matchs := re.FindAllString(msg, -1)

	if len(matchs) == 0 {
		return ""
	}

	supported := matchs[0]

	for i := len(serviceVersions) - 1; i >= 0; i-- {
		version := serviceVersions[i]
		if strings.Contains(supported, version) {
			return version
		}
	}

	return ""

}

var re = regexp.MustCompile(`(?m)supported=(.*?)]`)
