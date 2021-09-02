package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/v8platform/protos/example/ras-client/simpleClient"
	clientv1 "go.buf.build/v8platform/go-gen-ras/v8platform/rasapis/ras/client/v1"
	messagesv1 "go.buf.build/v8platform/go-gen-ras/v8platform/rasapis/ras/messages/v1"
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

	endpointService, err := client.Open("10.0")
	if err != nil {
		panic(err)
	}

	ras := clientv1.NewRasService(endpointService)
	clusters, err := ras.GetClusters(&messagesv1.GetClustersRequest{})
	if err != nil {
		return
	}

	fmt.Println("Список полученных сессий")

	sessions, err := ras.GetSessions(&messagesv1.GetSessionsRequest{ClusterId: clusters.Clusters[0].Uuid})
	if err != nil {
		return
	}
	// pp.SetDefaultMaxDepth(1)
	// pp.Println(resp.Sessions)
	for _, session := range sessions.Sessions {

		log.Println(session.String())
	}
}
