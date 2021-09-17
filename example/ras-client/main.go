package main

//
// func main() {
//
// 	var host string
// 	flag.StringVar(&host, "server", "srv-uk-app31:1545", "Адрес сервера и порт")
//
// 	flag.Parse()
//
// 	ctx := context.Background()
//
// 	client := simpleClient.NewClient(host)
// 	defer func() {
// 		err := client.Close()
// 		if err != nil {
// 			panic(err)
// 		}
// 	}()
//
// 	err := client.Connect(ctx)
// 	if err != nil {
// 		panic(err)
// 	}
//
// 	endpointService, err := client.Open("10.0")
// 	if err != nil {
// 		panic(err)
// 	}
//
// 	ras := clientv1.NewRasService(endpointService)
// 	clusters, err := ras.GetClusters(&messagesv1.GetClustersRequest{})
// 	if err != nil {
// 		panic(err)
// 	}
//
// 	_, err = ras.AuthenticateCluster(&messagesv1.ClusterAuthenticateRequest{ClusterId: clusters.Clusters[0].Uuid})
// 	if err != nil {
// 		panic(err)
// 	}
//
// 	sessions, err := ras.GetSessions(&messagesv1.GetSessionsRequest{ClusterId: clusters.Clusters[0].Uuid})
// 	if err != nil {
// 		panic(err)
// 	}
// 	// pp.SetDefaultMaxDepth(1)
// 	// pp.Println(resp.Sessions)
// 	fmt.Println("Список полученных сессий ", len(sessions.Sessions))
// 	for _, session := range sessions.Sessions {
//
// 		log.Println(session.String())
// 	}
//
// }
