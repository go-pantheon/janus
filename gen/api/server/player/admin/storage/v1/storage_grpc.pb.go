// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: player/admin/storage/v1/storage.proto

package adminv1

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
	StorageAdmin_AddItem_FullMethodName = "/player.admin.storage.v1.StorageAdmin/AddItem"
)

// StorageAdminClient is the client API for StorageAdmin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Storage admin service
// Open to the server cluster
// Provide HTTP and gRPC interfaces
type StorageAdminClient interface {
	AddItem(ctx context.Context, in *AddItemRequest, opts ...grpc.CallOption) (*AddItemResponse, error)
}

type storageAdminClient struct {
	cc grpc.ClientConnInterface
}

func NewStorageAdminClient(cc grpc.ClientConnInterface) StorageAdminClient {
	return &storageAdminClient{cc}
}

func (c *storageAdminClient) AddItem(ctx context.Context, in *AddItemRequest, opts ...grpc.CallOption) (*AddItemResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddItemResponse)
	err := c.cc.Invoke(ctx, StorageAdmin_AddItem_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StorageAdminServer is the server API for StorageAdmin service.
// All implementations must embed UnimplementedStorageAdminServer
// for forward compatibility.
//
// Storage admin service
// Open to the server cluster
// Provide HTTP and gRPC interfaces
type StorageAdminServer interface {
	AddItem(context.Context, *AddItemRequest) (*AddItemResponse, error)
	mustEmbedUnimplementedStorageAdminServer()
}

// UnimplementedStorageAdminServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedStorageAdminServer struct{}

func (UnimplementedStorageAdminServer) AddItem(context.Context, *AddItemRequest) (*AddItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddItem not implemented")
}
func (UnimplementedStorageAdminServer) mustEmbedUnimplementedStorageAdminServer() {}
func (UnimplementedStorageAdminServer) testEmbeddedByValue()                      {}

// UnsafeStorageAdminServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StorageAdminServer will
// result in compilation errors.
type UnsafeStorageAdminServer interface {
	mustEmbedUnimplementedStorageAdminServer()
}

func RegisterStorageAdminServer(s grpc.ServiceRegistrar, srv StorageAdminServer) {
	// If the following call pancis, it indicates UnimplementedStorageAdminServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&StorageAdmin_ServiceDesc, srv)
}

func _StorageAdmin_AddItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageAdminServer).AddItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StorageAdmin_AddItem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageAdminServer).AddItem(ctx, req.(*AddItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// StorageAdmin_ServiceDesc is the grpc.ServiceDesc for StorageAdmin service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StorageAdmin_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "player.admin.storage.v1.StorageAdmin",
	HandlerType: (*StorageAdminServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddItem",
			Handler:    _StorageAdmin_AddItem_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "player/admin/storage/v1/storage.proto",
}
