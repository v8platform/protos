package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/v8platform/protos/extra"
	"github.com/v8platform/protos/extra/encoding/rasbinary"
	messagesv1 "github.com/v8platform/protos/gen/ras/messages/v1"
	protocolv1 "github.com/v8platform/protos/gen/ras/protocol/v1"
	"google.golang.org/protobuf/proto"
	"io"
	"log"
	"net"
)

func main() {

	var host string
	flag.StringVar(&host, "server", "srv-uk-app10:1545", "Адрес сервера и порт")

	flag.Parse()
	fmt.Println("host has value ", host)

	ctx := context.Background()

	conn, err := connect(ctx, host)
	if err != nil {
		panic(err)
	}

	// err = conn.SetReadDeadline(time.Now().Add(5 * time.Second))
	// if err != nil {
	// 	panic(err)
	// }

	err = extra.Connect(conn)
	if err != nil {
		panic(err)
	}

	endpoint, err := extra.OpenEndpoint(conn, "10.0")
	if err != nil {
		panic(err)
	}

	log.Printf("%v", endpoint)

	err = endpoint.SendMessage(conn, &messagesv1.GetClustersRequest{})
	if err != nil {
		panic(err)
	}

	var Response messagesv1.GetClustersResponse

	err = endpoint.ReadMessage(conn, &Response)
	if err != nil {
		panic(err)
	}

	log.Println(Response.Clusters)

	// negotiate, err := NewPacket(&protocolv1.NegotiateMessage{
	//	Magic:    475223888, // Константы
	//	Protocol: 256,       // Константы
	//	Version:  256,       // Константы
	// })

	// clustersReq, err := newEndpointMessage(1, 255, &messagesv1.GetClustersRequest{})
	//
	// p, err := NewPacket(clustersReq)
	//
	// if err != nil {
	// 	panic(err)
	// }
	//
	// // pp.Println(negotiate)
	// log.Println(p.String())
	// //
	// _, err = writePacket(conn, p)
	// if err != nil {
	// 	panic(err)
	// }
	// //
	// // pp.Println(conn.Bytes())
	// //
	//
	// clustersResp, err := newEndpointMessage(1, 255, &messagesv1.GetClustersResponse{})
	// p, err = NewPacket(clustersResp)
	//
	// conn.Reset()
	// _, err = writePacket(conn, p)
	// if err != nil {
	// 	panic(err)
	// }
	// packet, err := ReadPacket(bytes.NewReader(conn.Bytes()), false)
	// if err != nil {
	// 	panic(err)
	// }
	//
	// log.Println(packet.String())
	// var v protocolv1.EndpointMessage
	// err = types.UnpackPacketDataTo(packet, &v)
	//
	// if err != nil {
	// 	panic(err)
	// }
	// log.Println(v.String())
	//
	// val, _ := types.UnpackPacketDataNew(packet)
	// log.Println(val.(*protocolv1.EndpointMessage).String())

}

func newEndpointMessage(id int32, format int32, m proto.Message) (*protocolv1.EndpointMessage, error) {

	md := m.ProtoReflect().Descriptor()

	isPacketMessage := proto.HasExtension(md.Options(), messagesv1.E_MessageType)

	if !isPacketMessage {
		return nil, fmt.Errorf("this is not packet message: <%s>", proto.MessageName(m))
	}

	dataType := proto.GetExtension(md.Options(), messagesv1.E_MessageType).(messagesv1.MessageType)

	encoder := rasbinary.MarshalOptions{
		ProtocolVersion: format,
	}

	b, err := encoder.Marshal(m)

	if err != nil {
		return nil, err
	}

	endpointM := &protocolv1.EndpointMessage{
		EndpointId: id,
		Format:     format,
		Type:       protocolv1.EndpointDataType_ENDPOINT_DATA_TYPE_MESSAGE,
		Data: &protocolv1.EndpointMessage_Message{
			&protocolv1.EndpointDataMessage{
				Type:  dataType,
				Bytes: b,
			}},
	}

	return endpointM, nil

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
