package main

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"github.com/k0kubun/pp"
	ras "github.com/v8platform/encoder/ras"
	extpb "github.com/v8platform/protos/gen/ras/encoding"
	protocolv1 "github.com/v8platform/protos/gen/ras/protocol/v1"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"io"
	"net"
	"sort"
)

var codec = ras.NewCodec()

func main() {

	var host string
	flag.StringVar(&host, "server", "localhost:1545", "Адрес сервера и порт")

	flag.Parse()
	fmt.Println("host has value ", host)

	ctx := context.Background()

	conn, err := connect(ctx, host)
	if err != nil {
		//panic(err)
	}

	negotiate := &protocolv1.NegotiateMessage{
		Magic:    475223888, // Константы
		Protocol: 256,       // Константы
		Version:  256,       // Константы
	}

	err = sendPacket(conn, protocolv1.PacketType_PACKET_TYPE_NEGOTIATE, encode(negotiate))
	if err != nil {
		panic(err)
	}

}

func sendPacket(conn net.Conn, packetType protocolv1.PacketType, data []byte) error {

	packet := &protocolv1.Packet{
		Type:   packetType,
		Data:   data,
		Length: int32(len(data)),
	}

	packetData := encode(packet)

	pp.Println(packetData)
	//err := writeMessage(conn, packetData)
	//return err
	return nil
}

func encode(message proto.Message) []byte {

	//message, ok := m.(proto.Message)

	//if !ok {
	//	panic("Non proto message")
	//}

	buf := &bytes.Buffer{}

	var fields []encodeField

	message.ProtoReflect().Range(func(descriptor protoreflect.FieldDescriptor, value protoreflect.Value) bool {

		encoderOptions := proto.GetExtension(descriptor.Options(), extpb.E_Field).(*extpb.EncodingFieldOptions)

		fields = append(fields, encodeField{
			order:   encoderOptions.Order,
			encoder: *encoderOptions.Encoder,
			value:   value.Interface(),
		})
		return true
	})

	sort.Slice(fields, func(i, j int) bool {
		return fields[i].order < fields[j].order
	})

	//pp.Println(fields)

	for _, field := range fields {

		_, err := EncodeValue(field.encoder, buf, field.value)
		if err != nil {
			panic(err)
		}
	}

	return buf.Bytes()
}

type encodeField struct {
	order   int32
	encoder string
	value   interface{}
}

func writeMessage(w io.WriteCloser, message []byte) error {

	_, err := w.Write(message)
	if err != nil {
		return err
	}

	return nil
}

func connect(ctx context.Context, addr string) (net.Conn, error) {

	_, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		return nil, err
	}

	var dialer net.Dialer

	conn, err := dialer.DialContext(ctx, "tcp", addr)
	if err != nil {
		return nil, err
	}

	return conn, nil

}
