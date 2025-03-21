// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: message/room.proto

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
	RoomTCPService_RoomList_FullMethodName                 = "/message.RoomTCPService/RoomList"
	RoomTCPService_RoomDetail_FullMethodName               = "/message.RoomTCPService/RoomDetail"
	RoomTCPService_CreateRoom_FullMethodName               = "/message.RoomTCPService/CreateRoom"
	RoomTCPService_InviteToJoinRoom_FullMethodName         = "/message.RoomTCPService/InviteToJoinRoom"
	RoomTCPService_AgreeToInviteJoinRoom_FullMethodName    = "/message.RoomTCPService/AgreeToInviteJoinRoom"
	RoomTCPService_RequestToJoinRoom_FullMethodName        = "/message.RoomTCPService/RequestToJoinRoom"
	RoomTCPService_ApproveRequestToJoinRoom_FullMethodName = "/message.RoomTCPService/ApproveRequestToJoinRoom"
	RoomTCPService_KickUserFromRoom_FullMethodName         = "/message.RoomTCPService/KickUserFromRoom"
	RoomTCPService_LeaveRoom_FullMethodName                = "/message.RoomTCPService/LeaveRoom"
	RoomTCPService_CloseRoom_FullMethodName                = "/message.RoomTCPService/CloseRoom"
)

// RoomTCPServiceClient is the client API for RoomTCPService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Room service
type RoomTCPServiceClient interface {
	// Room List
	RoomList(ctx context.Context, in *CSRoomList, opts ...grpc.CallOption) (*SCRoomList, error)
	// Room detail
	RoomDetail(ctx context.Context, in *CSRoomDetail, opts ...grpc.CallOption) (*SCRoomDetail, error)
	// Create room
	CreateRoom(ctx context.Context, in *CSCreateRoom, opts ...grpc.CallOption) (*SCCreateRoom, error)
	// Invite to join room
	InviteToJoinRoom(ctx context.Context, in *CSInviteToJoinRoom, opts ...grpc.CallOption) (*SCInviteToJoinRoom, error)
	// Agree to invite to join room
	AgreeToInviteJoinRoom(ctx context.Context, in *CSAgreeToInviteJoinRoom, opts ...grpc.CallOption) (*SCAgreeToInviteJoinRoom, error)
	// Request to join room
	RequestToJoinRoom(ctx context.Context, in *CSRequestToJoinRoom, opts ...grpc.CallOption) (*SCRequestToJoinRoom, error)
	// Approve request to join room
	ApproveRequestToJoinRoom(ctx context.Context, in *CSApproveRequestToJoinRoom, opts ...grpc.CallOption) (*SCApproveRequestToJoinRoom, error)
	// Kick user from room
	KickUserFromRoom(ctx context.Context, in *CSKickUserFromRoom, opts ...grpc.CallOption) (*SCKickUserFromRoom, error)
	// Leave room
	LeaveRoom(ctx context.Context, in *CSLeaveRoom, opts ...grpc.CallOption) (*SCLeaveRoom, error)
	// Close room
	CloseRoom(ctx context.Context, in *CSCloseRoom, opts ...grpc.CallOption) (*SCCloseRoom, error)
}

type roomTCPServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRoomTCPServiceClient(cc grpc.ClientConnInterface) RoomTCPServiceClient {
	return &roomTCPServiceClient{cc}
}

func (c *roomTCPServiceClient) RoomList(ctx context.Context, in *CSRoomList, opts ...grpc.CallOption) (*SCRoomList, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SCRoomList)
	err := c.cc.Invoke(ctx, RoomTCPService_RoomList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomTCPServiceClient) RoomDetail(ctx context.Context, in *CSRoomDetail, opts ...grpc.CallOption) (*SCRoomDetail, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SCRoomDetail)
	err := c.cc.Invoke(ctx, RoomTCPService_RoomDetail_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomTCPServiceClient) CreateRoom(ctx context.Context, in *CSCreateRoom, opts ...grpc.CallOption) (*SCCreateRoom, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SCCreateRoom)
	err := c.cc.Invoke(ctx, RoomTCPService_CreateRoom_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomTCPServiceClient) InviteToJoinRoom(ctx context.Context, in *CSInviteToJoinRoom, opts ...grpc.CallOption) (*SCInviteToJoinRoom, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SCInviteToJoinRoom)
	err := c.cc.Invoke(ctx, RoomTCPService_InviteToJoinRoom_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomTCPServiceClient) AgreeToInviteJoinRoom(ctx context.Context, in *CSAgreeToInviteJoinRoom, opts ...grpc.CallOption) (*SCAgreeToInviteJoinRoom, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SCAgreeToInviteJoinRoom)
	err := c.cc.Invoke(ctx, RoomTCPService_AgreeToInviteJoinRoom_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomTCPServiceClient) RequestToJoinRoom(ctx context.Context, in *CSRequestToJoinRoom, opts ...grpc.CallOption) (*SCRequestToJoinRoom, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SCRequestToJoinRoom)
	err := c.cc.Invoke(ctx, RoomTCPService_RequestToJoinRoom_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomTCPServiceClient) ApproveRequestToJoinRoom(ctx context.Context, in *CSApproveRequestToJoinRoom, opts ...grpc.CallOption) (*SCApproveRequestToJoinRoom, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SCApproveRequestToJoinRoom)
	err := c.cc.Invoke(ctx, RoomTCPService_ApproveRequestToJoinRoom_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomTCPServiceClient) KickUserFromRoom(ctx context.Context, in *CSKickUserFromRoom, opts ...grpc.CallOption) (*SCKickUserFromRoom, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SCKickUserFromRoom)
	err := c.cc.Invoke(ctx, RoomTCPService_KickUserFromRoom_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomTCPServiceClient) LeaveRoom(ctx context.Context, in *CSLeaveRoom, opts ...grpc.CallOption) (*SCLeaveRoom, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SCLeaveRoom)
	err := c.cc.Invoke(ctx, RoomTCPService_LeaveRoom_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomTCPServiceClient) CloseRoom(ctx context.Context, in *CSCloseRoom, opts ...grpc.CallOption) (*SCCloseRoom, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SCCloseRoom)
	err := c.cc.Invoke(ctx, RoomTCPService_CloseRoom_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RoomTCPServiceServer is the server API for RoomTCPService service.
// All implementations must embed UnimplementedRoomTCPServiceServer
// for forward compatibility.
//
// Room service
type RoomTCPServiceServer interface {
	// Room List
	RoomList(context.Context, *CSRoomList) (*SCRoomList, error)
	// Room detail
	RoomDetail(context.Context, *CSRoomDetail) (*SCRoomDetail, error)
	// Create room
	CreateRoom(context.Context, *CSCreateRoom) (*SCCreateRoom, error)
	// Invite to join room
	InviteToJoinRoom(context.Context, *CSInviteToJoinRoom) (*SCInviteToJoinRoom, error)
	// Agree to invite to join room
	AgreeToInviteJoinRoom(context.Context, *CSAgreeToInviteJoinRoom) (*SCAgreeToInviteJoinRoom, error)
	// Request to join room
	RequestToJoinRoom(context.Context, *CSRequestToJoinRoom) (*SCRequestToJoinRoom, error)
	// Approve request to join room
	ApproveRequestToJoinRoom(context.Context, *CSApproveRequestToJoinRoom) (*SCApproveRequestToJoinRoom, error)
	// Kick user from room
	KickUserFromRoom(context.Context, *CSKickUserFromRoom) (*SCKickUserFromRoom, error)
	// Leave room
	LeaveRoom(context.Context, *CSLeaveRoom) (*SCLeaveRoom, error)
	// Close room
	CloseRoom(context.Context, *CSCloseRoom) (*SCCloseRoom, error)
	mustEmbedUnimplementedRoomTCPServiceServer()
}

// UnimplementedRoomTCPServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedRoomTCPServiceServer struct{}

func (UnimplementedRoomTCPServiceServer) RoomList(context.Context, *CSRoomList) (*SCRoomList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RoomList not implemented")
}
func (UnimplementedRoomTCPServiceServer) RoomDetail(context.Context, *CSRoomDetail) (*SCRoomDetail, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RoomDetail not implemented")
}
func (UnimplementedRoomTCPServiceServer) CreateRoom(context.Context, *CSCreateRoom) (*SCCreateRoom, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRoom not implemented")
}
func (UnimplementedRoomTCPServiceServer) InviteToJoinRoom(context.Context, *CSInviteToJoinRoom) (*SCInviteToJoinRoom, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InviteToJoinRoom not implemented")
}
func (UnimplementedRoomTCPServiceServer) AgreeToInviteJoinRoom(context.Context, *CSAgreeToInviteJoinRoom) (*SCAgreeToInviteJoinRoom, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AgreeToInviteJoinRoom not implemented")
}
func (UnimplementedRoomTCPServiceServer) RequestToJoinRoom(context.Context, *CSRequestToJoinRoom) (*SCRequestToJoinRoom, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestToJoinRoom not implemented")
}
func (UnimplementedRoomTCPServiceServer) ApproveRequestToJoinRoom(context.Context, *CSApproveRequestToJoinRoom) (*SCApproveRequestToJoinRoom, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApproveRequestToJoinRoom not implemented")
}
func (UnimplementedRoomTCPServiceServer) KickUserFromRoom(context.Context, *CSKickUserFromRoom) (*SCKickUserFromRoom, error) {
	return nil, status.Errorf(codes.Unimplemented, "method KickUserFromRoom not implemented")
}
func (UnimplementedRoomTCPServiceServer) LeaveRoom(context.Context, *CSLeaveRoom) (*SCLeaveRoom, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LeaveRoom not implemented")
}
func (UnimplementedRoomTCPServiceServer) CloseRoom(context.Context, *CSCloseRoom) (*SCCloseRoom, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CloseRoom not implemented")
}
func (UnimplementedRoomTCPServiceServer) mustEmbedUnimplementedRoomTCPServiceServer() {}
func (UnimplementedRoomTCPServiceServer) testEmbeddedByValue()                        {}

// UnsafeRoomTCPServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RoomTCPServiceServer will
// result in compilation errors.
type UnsafeRoomTCPServiceServer interface {
	mustEmbedUnimplementedRoomTCPServiceServer()
}

func RegisterRoomTCPServiceServer(s grpc.ServiceRegistrar, srv RoomTCPServiceServer) {
	// If the following call pancis, it indicates UnimplementedRoomTCPServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&RoomTCPService_ServiceDesc, srv)
}

func _RoomTCPService_RoomList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CSRoomList)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomTCPServiceServer).RoomList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RoomTCPService_RoomList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomTCPServiceServer).RoomList(ctx, req.(*CSRoomList))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoomTCPService_RoomDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CSRoomDetail)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomTCPServiceServer).RoomDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RoomTCPService_RoomDetail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomTCPServiceServer).RoomDetail(ctx, req.(*CSRoomDetail))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoomTCPService_CreateRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CSCreateRoom)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomTCPServiceServer).CreateRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RoomTCPService_CreateRoom_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomTCPServiceServer).CreateRoom(ctx, req.(*CSCreateRoom))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoomTCPService_InviteToJoinRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CSInviteToJoinRoom)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomTCPServiceServer).InviteToJoinRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RoomTCPService_InviteToJoinRoom_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomTCPServiceServer).InviteToJoinRoom(ctx, req.(*CSInviteToJoinRoom))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoomTCPService_AgreeToInviteJoinRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CSAgreeToInviteJoinRoom)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomTCPServiceServer).AgreeToInviteJoinRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RoomTCPService_AgreeToInviteJoinRoom_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomTCPServiceServer).AgreeToInviteJoinRoom(ctx, req.(*CSAgreeToInviteJoinRoom))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoomTCPService_RequestToJoinRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CSRequestToJoinRoom)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomTCPServiceServer).RequestToJoinRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RoomTCPService_RequestToJoinRoom_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomTCPServiceServer).RequestToJoinRoom(ctx, req.(*CSRequestToJoinRoom))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoomTCPService_ApproveRequestToJoinRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CSApproveRequestToJoinRoom)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomTCPServiceServer).ApproveRequestToJoinRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RoomTCPService_ApproveRequestToJoinRoom_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomTCPServiceServer).ApproveRequestToJoinRoom(ctx, req.(*CSApproveRequestToJoinRoom))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoomTCPService_KickUserFromRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CSKickUserFromRoom)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomTCPServiceServer).KickUserFromRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RoomTCPService_KickUserFromRoom_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomTCPServiceServer).KickUserFromRoom(ctx, req.(*CSKickUserFromRoom))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoomTCPService_LeaveRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CSLeaveRoom)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomTCPServiceServer).LeaveRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RoomTCPService_LeaveRoom_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomTCPServiceServer).LeaveRoom(ctx, req.(*CSLeaveRoom))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoomTCPService_CloseRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CSCloseRoom)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomTCPServiceServer).CloseRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RoomTCPService_CloseRoom_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomTCPServiceServer).CloseRoom(ctx, req.(*CSCloseRoom))
	}
	return interceptor(ctx, in, info, handler)
}

// RoomTCPService_ServiceDesc is the grpc.ServiceDesc for RoomTCPService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RoomTCPService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "message.RoomTCPService",
	HandlerType: (*RoomTCPServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RoomList",
			Handler:    _RoomTCPService_RoomList_Handler,
		},
		{
			MethodName: "RoomDetail",
			Handler:    _RoomTCPService_RoomDetail_Handler,
		},
		{
			MethodName: "CreateRoom",
			Handler:    _RoomTCPService_CreateRoom_Handler,
		},
		{
			MethodName: "InviteToJoinRoom",
			Handler:    _RoomTCPService_InviteToJoinRoom_Handler,
		},
		{
			MethodName: "AgreeToInviteJoinRoom",
			Handler:    _RoomTCPService_AgreeToInviteJoinRoom_Handler,
		},
		{
			MethodName: "RequestToJoinRoom",
			Handler:    _RoomTCPService_RequestToJoinRoom_Handler,
		},
		{
			MethodName: "ApproveRequestToJoinRoom",
			Handler:    _RoomTCPService_ApproveRequestToJoinRoom_Handler,
		},
		{
			MethodName: "KickUserFromRoom",
			Handler:    _RoomTCPService_KickUserFromRoom_Handler,
		},
		{
			MethodName: "LeaveRoom",
			Handler:    _RoomTCPService_LeaveRoom_Handler,
		},
		{
			MethodName: "CloseRoom",
			Handler:    _RoomTCPService_CloseRoom_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "message/room.proto",
}
