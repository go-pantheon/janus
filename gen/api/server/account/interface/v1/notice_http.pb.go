// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.8.4
// - protoc             (unknown)
// source: account/interface/v1/notice.proto

package interfacev1

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

const OperationNoticeInterfaceNoticeList = "/account.interface.v1.NoticeInterface/NoticeList"

type NoticeInterfaceHTTPServer interface {
	// NoticeList Notice List
	NoticeList(context.Context, *NoticeListRequest) (*NoticeListResponse, error)
}

func RegisterNoticeInterfaceHTTPServer(s *http.Server, srv NoticeInterfaceHTTPServer) {
	r := s.Route("/")
	r.GET("/v1/notice/list", _NoticeInterface_NoticeList0_HTTP_Handler(srv))
}

func _NoticeInterface_NoticeList0_HTTP_Handler(srv NoticeInterfaceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in NoticeListRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationNoticeInterfaceNoticeList)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.NoticeList(ctx, req.(*NoticeListRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*NoticeListResponse)
		return ctx.Result(200, reply)
	}
}

type NoticeInterfaceHTTPClient interface {
	NoticeList(ctx context.Context, req *NoticeListRequest, opts ...http.CallOption) (rsp *NoticeListResponse, err error)
}

type NoticeInterfaceHTTPClientImpl struct {
	cc *http.Client
}

func NewNoticeInterfaceHTTPClient(client *http.Client) NoticeInterfaceHTTPClient {
	return &NoticeInterfaceHTTPClientImpl{client}
}

func (c *NoticeInterfaceHTTPClientImpl) NoticeList(ctx context.Context, in *NoticeListRequest, opts ...http.CallOption) (*NoticeListResponse, error) {
	var out NoticeListResponse
	pattern := "/v1/notice/list"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationNoticeInterfaceNoticeList))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
