package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/v8platform/protos/example/ras-client/simpleClient"
	clientv1 "github.com/v8platform/protos/gen/ras/client/v1"
	messagesv1 "github.com/v8platform/protos/gen/ras/messages/v1"
	serializev1 "github.com/v8platform/protos/gen/v8platform/serialize/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"net"

	apiv1 "github.com/v8platform/protos/gen/ras/service/api/v1"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {

	var host string
	flag.StringVar(&host, "server", "localhost:1545", "Адрес сервера и порт")

	flag.Parse()

	listenOn := "127.0.0.1:8080"
	listener, err := net.Listen("tcp", listenOn)
	if err != nil {
		return fmt.Errorf("failed to listen on %s: %w", listenOn, err)
	}

	srv := &rasClusterServiceServer{Host: host, client: simpleClient.NewClient(host)}
	server := grpc.NewServer()
	apiv1.RegisterClustersServiceServer(server, srv)
	RunGW(srv)

	log.Println("Listening on", listenOn)
	if err := server.Serve(listener); err != nil {
		return fmt.Errorf("failed to serve gRPC server: %w", err)
	}

	return nil
}

type rasClusterServiceServer struct {
	apiv1.UnimplementedClustersServiceServer
	client *simpleClient.Client

	Host string
}

func (s rasClusterServiceServer) Clusters(ctx context.Context, req *apiv1.GetClustersRequest) (*apiv1.GetClustersResponse, error) {

	err := s.client.Connect(ctx)
	if err != nil {
		return nil, err
	}
	defer s.client.Close()

	endpointService, err := s.client.Open("10.0")
	if err != nil {
		return nil, err
	}
	endpoint := endpointService.(clientv1.EndpointServiceImpl)

	clustersService := clientv1.NewClustersService(endpoint)

	resp, err := clustersService.GetClusters(&messagesv1.GetClustersRequest{})
	if err != nil {
		return nil, err
	}

	return &apiv1.GetClustersResponse{
		Items: resp.Clusters,
	}, nil
}
func (rasClusterServiceServer) GetCluster(ctx context.Context, req *apiv1.GetClusterRequest) (*serializev1.ClusterInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCluster not implemented")
}
func (rasClusterServiceServer) AddCluster(ctx context.Context, req *apiv1.AddClusterRequest) (*apiv1.AddClusterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCluster not implemented")
}
func (rasClusterServiceServer) DeleteCluster(ctx context.Context, req *apiv1.DeleteClusterRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCluster not implemented")
}
