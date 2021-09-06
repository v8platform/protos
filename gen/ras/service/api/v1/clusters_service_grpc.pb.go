// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package apiv1

import (
	context "context"
	v1 "github.com/v8platform/protos/gen/ras/messages/v1"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ClustersServiceClient is the client API for ClustersService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ClustersServiceClient interface {
	GetClusters(ctx context.Context, in *v1.GetClustersRequest, opts ...grpc.CallOption) (*v1.GetClustersResponse, error)
	GetClusterInfo(ctx context.Context, in *v1.GetClusterInfoRequest, opts ...grpc.CallOption) (*v1.GetClusterInfoResponse, error)
	RegCluster(ctx context.Context, in *v1.RegClusterRequest, opts ...grpc.CallOption) (*v1.RegClusterResponse, error)
	UnregCluster(ctx context.Context, in *v1.UnregClusterRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type clustersServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewClustersServiceClient(cc grpc.ClientConnInterface) ClustersServiceClient {
	return &clustersServiceClient{cc}
}

func (c *clustersServiceClient) GetClusters(ctx context.Context, in *v1.GetClustersRequest, opts ...grpc.CallOption) (*v1.GetClustersResponse, error) {
	out := new(v1.GetClustersResponse)
	err := c.cc.Invoke(ctx, "/ras.service.api.v1.ClustersService/GetClusters", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clustersServiceClient) GetClusterInfo(ctx context.Context, in *v1.GetClusterInfoRequest, opts ...grpc.CallOption) (*v1.GetClusterInfoResponse, error) {
	out := new(v1.GetClusterInfoResponse)
	err := c.cc.Invoke(ctx, "/ras.service.api.v1.ClustersService/GetClusterInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clustersServiceClient) RegCluster(ctx context.Context, in *v1.RegClusterRequest, opts ...grpc.CallOption) (*v1.RegClusterResponse, error) {
	out := new(v1.RegClusterResponse)
	err := c.cc.Invoke(ctx, "/ras.service.api.v1.ClustersService/RegCluster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clustersServiceClient) UnregCluster(ctx context.Context, in *v1.UnregClusterRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/ras.service.api.v1.ClustersService/UnregCluster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClustersServiceServer is the server API for ClustersService service.
// All implementations should embed UnimplementedClustersServiceServer
// for forward compatibility
type ClustersServiceServer interface {
	GetClusters(context.Context, *v1.GetClustersRequest) (*v1.GetClustersResponse, error)
	GetClusterInfo(context.Context, *v1.GetClusterInfoRequest) (*v1.GetClusterInfoResponse, error)
	RegCluster(context.Context, *v1.RegClusterRequest) (*v1.RegClusterResponse, error)
	UnregCluster(context.Context, *v1.UnregClusterRequest) (*emptypb.Empty, error)
}

// UnimplementedClustersServiceServer should be embedded to have forward compatible implementations.
type UnimplementedClustersServiceServer struct {
}

func (UnimplementedClustersServiceServer) GetClusters(context.Context, *v1.GetClustersRequest) (*v1.GetClustersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClusters not implemented")
}
func (UnimplementedClustersServiceServer) GetClusterInfo(context.Context, *v1.GetClusterInfoRequest) (*v1.GetClusterInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClusterInfo not implemented")
}
func (UnimplementedClustersServiceServer) RegCluster(context.Context, *v1.RegClusterRequest) (*v1.RegClusterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegCluster not implemented")
}
func (UnimplementedClustersServiceServer) UnregCluster(context.Context, *v1.UnregClusterRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnregCluster not implemented")
}

// UnsafeClustersServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ClustersServiceServer will
// result in compilation errors.
type UnsafeClustersServiceServer interface {
	mustEmbedUnimplementedClustersServiceServer()
}

func RegisterClustersServiceServer(s grpc.ServiceRegistrar, srv ClustersServiceServer) {
	s.RegisterService(&ClustersService_ServiceDesc, srv)
}

func _ClustersService_GetClusters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.GetClustersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClustersServiceServer).GetClusters(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ras.service.api.v1.ClustersService/GetClusters",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClustersServiceServer).GetClusters(ctx, req.(*v1.GetClustersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClustersService_GetClusterInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.GetClusterInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClustersServiceServer).GetClusterInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ras.service.api.v1.ClustersService/GetClusterInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClustersServiceServer).GetClusterInfo(ctx, req.(*v1.GetClusterInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClustersService_RegCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.RegClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClustersServiceServer).RegCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ras.service.api.v1.ClustersService/RegCluster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClustersServiceServer).RegCluster(ctx, req.(*v1.RegClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClustersService_UnregCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.UnregClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClustersServiceServer).UnregCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ras.service.api.v1.ClustersService/UnregCluster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClustersServiceServer).UnregCluster(ctx, req.(*v1.UnregClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ClustersService_ServiceDesc is the grpc.ServiceDesc for ClustersService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ClustersService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ras.service.api.v1.ClustersService",
	HandlerType: (*ClustersServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetClusters",
			Handler:    _ClustersService_GetClusters_Handler,
		},
		{
			MethodName: "GetClusterInfo",
			Handler:    _ClustersService_GetClusterInfo_Handler,
		},
		{
			MethodName: "RegCluster",
			Handler:    _ClustersService_RegCluster_Handler,
		},
		{
			MethodName: "UnregCluster",
			Handler:    _ClustersService_UnregCluster_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ras/service/api/v1/clusters_service.proto",
}
