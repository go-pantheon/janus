// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: message/dev_service.proto

package climsg

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	DevService_DevList_FullMethodName    = "/message.DevService/DevList"
	DevService_DevExecute_FullMethodName = "/message.DevService/DevExecute"
)

// DevServiceClient is the client API for DevService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Development debugging command protocol
type DevServiceClient interface {
	// Dev command list
	DevList(ctx context.Context, in *CSDevList, opts ...grpc.CallOption) (*SCDevList, error)
	// Execute Dev command
	DevExecute(ctx context.Context, in *CSDevExecute, opts ...grpc.CallOption) (*SCDevExecute, error)
}

type devServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDevServiceClient(cc grpc.ClientConnInterface) DevServiceClient {
	return &devServiceClient{cc}
}

func (c *devServiceClient) DevList(ctx context.Context, in *CSDevList, opts ...grpc.CallOption) (*SCDevList, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SCDevList)
	err := c.cc.Invoke(ctx, DevService_DevList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *devServiceClient) DevExecute(ctx context.Context, in *CSDevExecute, opts ...grpc.CallOption) (*SCDevExecute, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SCDevExecute)
	err := c.cc.Invoke(ctx, DevService_DevExecute_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DevServiceServer is the server API for DevService service.
// All implementations must embed UnimplementedDevServiceServer
// for forward compatibility.
//
// Development debugging command protocol
type DevServiceServer interface {
	// Dev command list
	DevList(context.Context, *CSDevList) (*SCDevList, error)
	// Execute Dev command
	DevExecute(context.Context, *CSDevExecute) (*SCDevExecute, error)
	mustEmbedUnimplementedDevServiceServer()
}

// UnimplementedDevServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedDevServiceServer struct{}

func (UnimplementedDevServiceServer) DevList(context.Context, *CSDevList) (*SCDevList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DevList not implemented")
}
func (UnimplementedDevServiceServer) DevExecute(context.Context, *CSDevExecute) (*SCDevExecute, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DevExecute not implemented")
}
func (UnimplementedDevServiceServer) mustEmbedUnimplementedDevServiceServer() {}
func (UnimplementedDevServiceServer) testEmbeddedByValue()                    {}

// UnsafeDevServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DevServiceServer will
// result in compilation errors.
type UnsafeDevServiceServer interface {
	mustEmbedUnimplementedDevServiceServer()
}

func RegisterDevServiceServer(s grpc.ServiceRegistrar, srv DevServiceServer) {
	// If the following call pancis, it indicates UnimplementedDevServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&DevService_ServiceDesc, srv)
}

func _DevService_DevList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CSDevList)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DevServiceServer).DevList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DevService_DevList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DevServiceServer).DevList(ctx, req.(*CSDevList))
	}
	return interceptor(ctx, in, info, handler)
}

func _DevService_DevExecute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CSDevExecute)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DevServiceServer).DevExecute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DevService_DevExecute_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DevServiceServer).DevExecute(ctx, req.(*CSDevExecute))
	}
	return interceptor(ctx, in, info, handler)
}

// DevService_ServiceDesc is the grpc.ServiceDesc for DevService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DevService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "message.DevService",
	HandlerType: (*DevServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DevList",
			Handler:    _DevService_DevList_Handler,
		},
		{
			MethodName: "DevExecute",
			Handler:    _DevService_DevExecute_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "message/dev_service.proto",
}
