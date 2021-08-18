// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// EndpointServiceClient is the client API for EndpointService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EndpointServiceClient interface {
	Open(ctx context.Context, in *EndpointOpenRequest, opts ...grpc.CallOption) (*EndpointOpenResponse, error)
	Request(ctx context.Context, in *EndpointRequest, opts ...grpc.CallOption) (*EndpointResponse, error)
}

type endpointServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEndpointServiceClient(cc grpc.ClientConnInterface) EndpointServiceClient {
	return &endpointServiceClient{cc}
}

func (c *endpointServiceClient) Open(ctx context.Context, in *EndpointOpenRequest, opts ...grpc.CallOption) (*EndpointOpenResponse, error) {
	out := new(EndpointOpenResponse)
	err := c.cc.Invoke(ctx, "/ras.api.v1.EndpointService/Open", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *endpointServiceClient) Request(ctx context.Context, in *EndpointRequest, opts ...grpc.CallOption) (*EndpointResponse, error) {
	out := new(EndpointResponse)
	err := c.cc.Invoke(ctx, "/ras.api.v1.EndpointService/Request", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EndpointServiceServer is the server API for EndpointService service.
// All implementations should embed UnimplementedEndpointServiceServer
// for forward compatibility
type EndpointServiceServer interface {
	Open(context.Context, *EndpointOpenRequest) (*EndpointOpenResponse, error)
	Request(context.Context, *EndpointRequest) (*EndpointResponse, error)
}

// UnimplementedEndpointServiceServer should be embedded to have forward compatible implementations.
type UnimplementedEndpointServiceServer struct {
}

func (UnimplementedEndpointServiceServer) Open(context.Context, *EndpointOpenRequest) (*EndpointOpenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Open not implemented")
}
func (UnimplementedEndpointServiceServer) Request(context.Context, *EndpointRequest) (*EndpointResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Request not implemented")
}

// UnsafeEndpointServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EndpointServiceServer will
// result in compilation errors.
type UnsafeEndpointServiceServer interface {
	mustEmbedUnimplementedEndpointServiceServer()
}

func RegisterEndpointServiceServer(s grpc.ServiceRegistrar, srv EndpointServiceServer) {
	s.RegisterService(&EndpointService_ServiceDesc, srv)
}

func _EndpointService_Open_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EndpointOpenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EndpointServiceServer).Open(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ras.api.v1.EndpointService/Open",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EndpointServiceServer).Open(ctx, req.(*EndpointOpenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EndpointService_Request_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EndpointRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EndpointServiceServer).Request(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ras.api.v1.EndpointService/Request",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EndpointServiceServer).Request(ctx, req.(*EndpointRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// EndpointService_ServiceDesc is the grpc.ServiceDesc for EndpointService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EndpointService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ras.api.v1.EndpointService",
	HandlerType: (*EndpointServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Open",
			Handler:    _EndpointService_Open_Handler,
		},
		{
			MethodName: "Request",
			Handler:    _EndpointService_Request_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v8platform/ras/api/v1/endpoint.proto",
}