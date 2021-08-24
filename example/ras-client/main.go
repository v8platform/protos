package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/v8platform/protos/example/ras-client/simpleClient"
	messagesv1 "github.com/v8platform/protos/gen/ras/messages/v1"
	"log"
)

func main() {

	var host string
	flag.StringVar(&host, "server", "localhost:1545", "Адрес сервера и порт")

	flag.Parse()

	ctx := context.Background()

	client := simpleClient.NewClient(host)

	err := client.Connect(ctx)
	if err != nil {
		panic(err)
	}

	endpoint, err := client.Open("10.0")
	if err != nil {
		panic(err)
	}

	log.Printf("%v", endpoint)

	err = endpoint.SendMessage(client, &messagesv1.GetClustersRequest{})
	if err != nil {
		panic(err)
	}

	var Response messagesv1.GetClustersResponse

	err = endpoint.ReadMessage(client, &Response)
	if err != nil {
		panic(err)
	}

	log.Println(Response.Clusters)

	err = endpoint.SendMessage(client, &messagesv1.ClusterAuthenticateRequest{
		ClusterId: Response.Clusters[0].Uuid,
	})
	if err != nil {
		panic(err)
	}

	err = endpoint.ReadVoidMessage(client)
	if err != nil {
		panic(err)
	}

	err = endpoint.SendMessage(client, &messagesv1.GetSessionsRequest{
		ClusterId: Response.Clusters[0].Uuid,
	})
	if err != nil {
		panic(err)
	}

	var resp messagesv1.GetSessionsResponse

	err = endpoint.ReadMessage(client, &resp)
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
