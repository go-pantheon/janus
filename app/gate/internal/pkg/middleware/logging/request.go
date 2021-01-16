package logging

import (
	"context"
	"strconv"

	"github.com/vulcan-frame/vulcan-gate/app/gate/internal/pkg/pool"
	climod "github.com/vulcan-frame/vulcan-gate/gen/api/client/module"
	cliseq "github.com/vulcan-frame/vulcan-gate/gen/api/client/sequence"
	"github.com/vulcan-frame/vulcan-gate/pkg/net"
	vctx "github.com/vulcan-frame/vulcan-gate/pkg/net/context"
	"github.com/vulcan-frame/vulcan-pkg-app/profile"
	"google.golang.org/protobuf/proto"
)

func Request(netKind net.NetKind) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			if !profile.IsDev() {
				return handler(ctx, req)
			}

			logRequest(ctx, netKind, req)
			return handler(ctx, req)
		}
	}
}

func logRequest(ctx context.Context, netKind net.NetKind, req interface{}) {
	p := pool.GetPacket()
	defer pool.PutPacket(p)

	if req, ok := req.([]byte); ok {
		if err := proto.Unmarshal(req, p); err != nil {
			log.Errorf("log request packet unmarshal failed. %s", err.Error())
			return
		}

		var (
			uid   string
			sid   string
			obj   string
			mod   string
			seq   string
			index string
		)

		if md, ok := metadata.FromServerContext(ctx); ok {
			uid = md.Get(vctx.CtxUID)
			sid = md.Get(vctx.CtxSID)
		}
		obj = strconv.FormatInt(p.Obj, 10)
		mod = strconv.FormatInt(int64(p.Mod), 10)
		seq = strconv.FormatInt(int64(p.Seq), 10)
		index = strconv.FormatInt(int64(p.Index), 10)

		kv := make([]interface{}, 0, 8*2)
		kv = append(kv, "kind", "req",
			"net", netKind,
			"uid", uid,
			"sid", sid,
			"obj", obj,
			"mod", mod,
			"seq", seq,
			"index", index,
		)

		log.Debugw(kv...)
	}
}
