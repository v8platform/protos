package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/v8platform/protos/extra"
	messagesv1 "github.com/v8platform/protos/gen/ras/messages/v1"
	"log"
	"net"
)

func main() {

	var host string
	flag.StringVar(&host, "server", "localhost:1545", "Адрес сервера и порт")

	flag.Parse()
	fmt.Println("host has value ", host)

	ctx := context.Background()

	conn, err := connect(ctx, host)
	if err != nil {
		panic(err)
	}

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

	err = endpoint.SendMessage(conn, &messagesv1.ClusterAuthenticateRequest{
		ClusterId: Response.Clusters[0].Uuid,
	})
	if err != nil {
		panic(err)
	}

	err = endpoint.ReadVoidMessage(conn)
	if err != nil {
		panic(err)
	}

	err = endpoint.SendMessage(conn, &messagesv1.GetSessionsRequest{
		ClusterId: Response.Clusters[0].Uuid,
	})
	if err != nil {
		panic(err)
	}

	var resp messagesv1.GetSessionsResponse

	err = endpoint.ReadMessage(conn, &resp)
	if err != nil {
		panic(err)
	}

	fmt.Println("Список полученных сессий")
	// pp.SetDefaultMaxDepth(1)
	// pp.Println(resp.Sessions)
	for _, session := range resp.Sessions {

		log.Println(session.String())
	}
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
