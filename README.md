# proto

Public interface definitions of 1C Enterprise

## How to use

1. Install `buf`on your machine

```shell
go install github.com/bufbuild/buf/cmd/buf # required
go install github.com/bufbuild/buf/cmd/protoc-gen-buf-breaking
go install github.com/bufbuild/buf/cmd/protoc-gen-buf-lint
```

2. Install protobuf & plugins

* Install `protoc` https://grpc.io/docs/protoc-installation/
* Install `plugins`

```shell
go install google.golang.org/protobuf/cmd/protoc-gen-go
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
go install github.com/srikrsna/protoc-gen-gotag
```

3. Add to your project file `buf.gen.yaml` to generate `*.pb.go`

```yaml
---
version: v1
managed:
  enabled: true
  go_package_prefix:
    default: github.com/your-name/repo-name/gen/go
    except:
      - buf.build/googleapis/googleapis
      - googleapies/google/api

plugins:
  - name: go
    out: ./gen/go
    opt: paths=source_relative
  - name: go-grpc
    out: ./gen/go
    opt: paths=source_relative,require_unimplemented_servers=false
  #  - name: grpc-gateway
  #    out: ./gen/go
  #    opt:
  #      - paths=source_relative
  #      - generate_unbound_methods=true
  #      - grpc_api_configuration=grpc-rest-bindings.yml
  - name: gotag
    out: .
    opt:
      - paths=source_relative
      - outdir=./gen/go

```

4. Create file `grpc-rest-bindings.yaml` for generate `grpc-gateway`:

```yaml
---
---
type: google.api.Service
config_version: 3

name: ras.api.v1
title: RAS API v1

apis:
  - name: ClustersService
  - mane: ClusterAdminService
  - mane: InfobasesService
http:
  rules:
    # Clusters
    - selector: ras.api.v1.ClustersService.Clusters
      get: '/clusters'
    - selector: ras.api.v1.ClustersService.GetCluster
      get: '/clusters/{cluster.id}'
    - selector: ras.api.v1.ClustersService.AddCluster
      post: '/clusters'
      body: '*'
    - selector: ras.api.v1.ClustersService.DeleteCluster
      delete: '/clusters/{cluster.id}'

    # Cluster admin
    - selector: ras.api.v1.ClusterAdminService.Admins
      get: '/clusters/{cluster.id}/admins'
      additional_bindings:
        - get: "/admins"
    - selector: ras.api.v1.ClusterAdminService.AddAdmin
      post: '/clusters/{cluster.id}/admins'
      body: 'admin_info'
      additional_bindings:
        - post: "/admins"
          body: 'admin_info'
    - selector: ras.api.v1.ClusterAdminService.DeleteAdmin
      delete: '/clusters/{cluster.id}/admins/{admin_name}'
      additional_bindings:
        - delete: "/admins/{admin_name}"

    # Infobases
    - selector: ras.api.v1.InfobasesService.Infobases
      get: '/infobases'
      additional_bindings:
        - get: '/clusters/{cluster.id}/infobases'
    - selector: ras.api.v1.InfobasesService.LookupInfobase
      get: "/infobases/lookup"
      additional_bindings:
        - get: '/clusters/{cluster.id}/infobases/lookup'
```

5. Run generate command

```shell
buf generate https://github.com/v8platform/protos.git
```

6. Code your client or server file

* file for server `Server/main.go`

```go
package main

import (
	"context"
	"fmt"
	"log"
	"net"

	// This import path is based on the name declaration in the go.mod,
	// and the gen/proto/go output location in the buf.gen.yaml.
	rasv1 "github.com/your-name/repo-name/gen/go/ras/api/v1"
	serializev1 "github.com/your-name/repo-name/gen/go/v8platform/protocol/v1"

	"google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	listenOn := "127.0.0.1:8080"
	listener, err := net.Listen("tcp", listenOn)
	if err != nil {
		return fmt.Errorf("failed to listen on %s: %w", listenOn, err)
	}

	server := grpc.NewServer()
	rasv1.RegisterClustersServiceServer(server, &rasClusterServiceServer{})
	log.Println("Listening on", listenOn)
	if err := server.Serve(listener); err != nil {
		return fmt.Errorf("failed to serve gRPC server: %w", err)
	}

	return nil
}

// petStoreServiceServer implements the PetStoreService API.
type rasClusterServiceServer struct {
	rasv1.UnimplementedClustersServiceServer
}

func (rasClusterServiceServer) Clusters(ctx context.Context, req *rasv1.GetClustersRequest) (*rasv1.GetClustersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Clusters not implemented")
}
func (rasClusterServiceServer) GetCluster(ctx context.Context, req *rasv1.GetClusterRequest) (*serializev1.ClusterInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCluster not implemented")
}
func (rasClusterServiceServer) AddCluster(ctx context.Context, req *rasv1.AddClusterRequest) (*rasv1.AddClusterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCluster not implemented")
}
func (rasClusterServiceServer) DeleteCluster(ctx context.Context, req *rasv1.DeleteClusterRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCluster not implemented")
}
```

* file for client `client/main.go`

```go
package main

import (
	"context"
	"fmt"
	"log"

	// This import path is based on the name declaration in the go.mod,
	// and the gen/proto/go output location in the buf.gen.yaml.
	rasv1 "github.com/your-name/repo-name/gen/go/ras/api/v1"
	"google.golang.org/grpc"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
func run() error {
	connectTo := "127.0.0.1:8080"
	conn, err := grpc.Dial(connectTo, grpc.WithBlock(), grpc.WithInsecure())
	if err != nil {
		return fmt.Errorf("failed to connect to ClustersService on %s: %w", connectTo, err)
	}
	log.Println("Connected to", connectTo)

	clusterService := rasv1.NewClustersServiceClient(conn)
	resp, err := clusterService.Clusters(context.Background(), &rasv1.GetClustersRequest{})
	if err != nil {
		return fmt.Errorf("failed to Clusters: %w", err)
	}

	log.Printf("Successfully Clusters %s", resp.String())
	return nil
}
```

7. Yohooo run it. Take many profit