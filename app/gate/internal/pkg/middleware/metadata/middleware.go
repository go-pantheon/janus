package metadata

import (
	"context"
	"strconv"

	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	clipkt "github.com/go-pantheon/janus/gen/api/client/packet"
	"github.com/go-pantheon/vulcan-kit/profile"
	"github.com/go-pantheon/vulcan-net"
	"google.golang.org/grpc/metadata"
)

func Server() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			requestHeader := newRequestMetadata(req)
			replyHeader := metadata.MD{}
			ctx = transport.NewServerContext(ctx, net.NewTransport(
				profile.GRPCEndpoint(),
				"",
				net.HeaderCarrier(requestHeader),
				net.HeaderCarrier(replyHeader),
			))
			return handler(ctx, req)
		}
	}
}

func newRequestMetadata(pack interface{}) metadata.MD {
	md := metadata.New(make(map[string]string, 8))
	p, ok := pack.(*clipkt.Packet)
	if !ok {
		return md
	}

	md.Set("ver", strconv.FormatInt(int64(p.Ver), 10))
	md.Set("index", strconv.FormatInt(int64(p.Index), 10))
	md.Set("cp", strconv.FormatBool(p.Compress))
	md.Set("mod", strconv.FormatInt(int64(p.Mod), 10))
	md.Set("seq", strconv.FormatInt(int64(p.Seq), 10))
	md.Set("obj", strconv.FormatInt(p.Obj, 10))
	return md
}
