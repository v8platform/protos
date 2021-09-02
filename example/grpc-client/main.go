package main

//
// import (
// 	"context"
// 	"fmt"
// 	apiv1 "github.com/v8platform/protos/gen/ras/api/v1"
// 	"google.golang.org/grpc"
// 	"log"
// )
//
// func main() {
// 	if err := run(); err != nil {
// 		log.Fatal(err)
// 	}
// }
// func run() error {
// 	connectTo := "127.0.0.1:8080"
// 	conn, err := grpc.Dial(connectTo, grpc.WithBlock(), grpc.WithInsecure())
// 	if err != nil {
// 		return fmt.Errorf("failed to connect to ClustersService on %s: %w", connectTo, err)
// 	}
// 	log.Println("Connected to", connectTo)
//
// 	clusterService := apiv1.NewClustersServiceClient(conn)
// 	resp, err := clusterService.Clusters(context.Background(), &apiv1.GetClustersRequest{})
// 	if err != nil {
// 		return fmt.Errorf("failed to Clusters: %w", err)
// 	}
//
// 	log.Printf("Successfully Clusters \n %v", resp.Items)
// 	return nil
// }
