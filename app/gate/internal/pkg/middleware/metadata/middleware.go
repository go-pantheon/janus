package metadata

import (
	"context"
	"strconv"

	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-pantheon/fabrica-kit/profile"
	"github.com/go-pantheon/fabrica-net/xnet"
	clipkt "github.com/go-pantheon/janus/gen/api/client/packet"
	"google.golang.org/grpc/metadata"
)

func Server() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req any) (any, error) {
			requestHeader := newRequestMetadata(req)
			replyHeader := metadata.MD{}

			ctx = transport.NewServerContext(ctx, xnet.NewTransport(
				profile.GRPCEndpoint(),
				"",
				xnet.HeaderCarrier(requestHeader),
				xnet.HeaderCarrier(replyHeader),
			))

			return handler(ctx, req)
		}
	}
}

func newRequestMetadata(pack any) metadata.MD {
	md := metadata.New(make(map[string]string, 8))

	p, ok := pack.(*clipkt.Packet)
	if !ok {
		return md
	}

	md.Set("ver", strconv.FormatInt(int64(p.Ver), 10))
	md.Set("idx", strconv.FormatInt(int64(p.Index), 10))
	md.Set("cmp", strconv.FormatBool(p.Compress))
	md.Set("mod", strconv.FormatInt(int64(p.Mod), 10))
	md.Set("seq", strconv.FormatInt(int64(p.Seq), 10))
	md.Set("obj", strconv.FormatInt(p.Obj, 10))

	return md
}
