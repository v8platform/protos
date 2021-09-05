package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/v8platform/protos/example/ras-client/simpleClient"
	clientv1 "github.com/v8platform/protos/gen/ras/client/v1"
	"github.com/v8platform/protos/gen/ras/protocol/messages"
	"log"
)

func main() {

	var host string
	flag.StringVar(&host, "server", "srv-uk-app31:1545", "Адрес сервера и порт")

	flag.Parse()

	ctx := context.Background()

	client := simpleClient.NewClient(host)
	defer func() {
		err := client.Close()
		if err != nil {
			panic(err)
		}
	}()

	err := client.Connect(ctx)
	if err != nil {
		panic(err)
	}

	endpointService, err := client.Open("10.0")
	if err != nil {
		panic(err)
	}

	ras := clientv1.NewRasService(endpointService)
	clusters, err := ras.GetClusters(&messages.GetClustersRequest{})
	if err != nil {
		panic(err)
	}

	_, err = ras.AuthenticateCluster(&messages.ClusterAuthenticateRequest{ClusterId: clusters.Clusters[0].Uuid})
	if err != nil {
		panic(err)
	}

	sessions, err := ras.GetSessions(&messages.GetSessionsRequest{ClusterId: clusters.Clusters[0].Uuid})
	if err != nil {
		panic(err)
	}
	// pp.SetDefaultMaxDepth(1)
	// pp.Println(resp.Sessions)
	fmt.Println("Список полученных сессий ", len(sessions.Sessions))
	for _, session := range sessions.Sessions {

		log.Println(session.String())
	}

}
