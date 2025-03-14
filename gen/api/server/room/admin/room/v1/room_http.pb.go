// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.8.4
// - protoc             (unknown)
// source: room/admin/room/v1/room.proto

package adminv1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationRoomAdminGetById = "/room.admin.room.v1.RoomAdmin/GetById"

type RoomAdminHTTPServer interface {
	GetById(context.Context, *GetByIdRequest) (*GetByIdResponse, error)
}

func RegisterRoomAdminHTTPServer(s *http.Server, srv RoomAdminHTTPServer) {
	r := s.Route("/")
	r.GET("/admin/room/id", _RoomAdmin_GetById0_HTTP_Handler(srv))
}

func _RoomAdmin_GetById0_HTTP_Handler(srv RoomAdminHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetByIdRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationRoomAdminGetById)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetById(ctx, req.(*GetByIdRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetByIdResponse)
		return ctx.Result(200, reply)
	}
}

type RoomAdminHTTPClient interface {
	GetById(ctx context.Context, req *GetByIdRequest, opts ...http.CallOption) (rsp *GetByIdResponse, err error)
}

type RoomAdminHTTPClientImpl struct {
	cc *http.Client
}

func NewRoomAdminHTTPClient(client *http.Client) RoomAdminHTTPClient {
	return &RoomAdminHTTPClientImpl{client}
}

func (c *RoomAdminHTTPClientImpl) GetById(ctx context.Context, in *GetByIdRequest, opts ...http.CallOption) (*GetByIdResponse, error) {
	var out GetByIdResponse
	pattern := "/admin/room/id"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationRoomAdminGetById))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
