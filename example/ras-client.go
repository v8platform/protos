package main

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"github.com/k0kubun/pp"
	ras "github.com/v8platform/encoder/ras"
	ras2 "github.com/v8platform/protos/example/ras"
	extpb "github.com/v8platform/protos/gen/ras/encoding"
	messagesv1 "github.com/v8platform/protos/gen/ras/messages/v1"
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

	//ctx := context.Background()
	conn := bytes.NewBufferString("")

	//conn, err := connect(ctx, host)
	//if err != nil {
	//	//panic(err)
	//}

	negotiate, err := NewPacket(&protocolv1.NegotiateMessage{
		Magic:    475223888, // Константы
		Protocol: 256,       // Константы
		Version:  256,       // Константы
	})

	clustersReq, err := newEndpointMessage(1, 255, &messagesv1.GetClustersRequest{})

	p, err := NewPacket(clustersReq)

	if err != nil {
		panic(err)
	}

	pp.Println(negotiate)
	pp.Println(p)

	_, err = writePacket(conn, negotiate)
	if err != nil {
		panic(err)
	}

	pp.Println(conn.Bytes())

	packet, err := ReadPacket(bytes.NewReader(conn.Bytes()), false)
	if err != nil {
		panic(err)
	}

	pp.Println(packet.GetBytes())

}

func newEndpointMessage(id int32, format int32, m proto.Message) (*protocolv1.EndpointMessage, error) {

	md := m.ProtoReflect().Descriptor()

	isPacketMessage := proto.HasExtension(md.Options(), messagesv1.E_MessageType)

	if !isPacketMessage {
		return nil, fmt.Errorf("this is not packet message: <%s>", proto.MessageName(m))
	}

	dataType := proto.GetExtension(md.Options(), messagesv1.E_MessageType).(messagesv1.EndpointMessageType)

	bytes, err := encode(m)

	if err != nil {
		return nil, err
	}

	endpointM := &protocolv1.EndpointMessage{
		EndpointId: id,
		Format:     format,
		Type:       protocolv1.EndpointDataType_ENDPOINT_DATA_TYPE_MESSAGE,
		Data: &protocolv1.EndpointDataMessage{
			Type: dataType,
			Data: &protocolv1.EndpointDataMessage_Bytes{bytes},
		},
	}

	return endpointM, nil

}

func encode(message proto.Message) ([]byte, error) {

	//message, ok := m.(proto.Message)

	//if !ok {
	//	panic("Non proto message")
	//}

	buf := &bytes.Buffer{}

	var fields []encodeField
	pRef := message.ProtoReflect()
	mFields := pRef.Descriptor().Fields()
	mOneOfs := pRef.Descriptor().Oneofs()

	idx := make(map[protoreflect.FieldNumber]struct{})
	for i := 0; i < mOneOfs.Len(); i++ {
		fieldDescr := mOneOfs.Get(i)

		descriptor := pRef.WhichOneof(fieldDescr)

		for i := 0; i < fieldDescr.Fields().Len(); i++ {
			idx[fieldDescr.Fields().Get(i).Number()] = struct{}{}
		}

		encoderOptions := proto.GetExtension(descriptor.Options(), extpb.E_Field).(*extpb.EncodingFieldOptions)
		value := pRef.Get(descriptor)

		fields = append(fields, encodeField{
			fd:      descriptor,
			order:   encoderOptions.GetOrder(),
			encoder: encoderOptions.GetEncoder(),
			value:   value,
		})
	}

	for i := 0; i < mFields.Len(); i++ {
		descriptor := mFields.Get(i)
		if _, ok := idx[descriptor.Number()]; ok {
			continue
		}
		value := pRef.Get(descriptor)
		encoderOptions := proto.GetExtension(descriptor.Options(), extpb.E_Field).(*extpb.EncodingFieldOptions)

		fields = append(fields, encodeField{
			fd:      descriptor,
			order:   encoderOptions.GetOrder(),
			encoder: encoderOptions.GetEncoder(),
			value:   value,
		})
	}

	sort.Slice(fields, func(i, j int) bool {
		return fields[i].order < fields[j].order
	})

	for _, field := range fields {

		switch field.fd.Kind() {
		case protoreflect.MessageKind:
			i, err := encode(field.value.Message().Interface())
			if err != nil {
				return nil, err
			}
			buf.Write(i)
		case protoreflect.EnumKind:

			_, err := ras2.EncodeValue(field.encoder, buf, int32(field.value.Enum()))
			if err != nil {
				panic(err)
			}

		default:
			_, err := ras2.EncodeValue(field.encoder, buf, field.value.Interface())
			if err != nil {
				panic(err)
			}
		}

	}

	return buf.Bytes(), nil
}

type encodeField struct {
	order   int32
	encoder string
	value   protoreflect.Value
	fd      protoreflect.FieldDescriptor
	opts    *extpb.EncodingFieldOptions
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
