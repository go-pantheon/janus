package logging

import (
	"context"
	"strconv"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/metadata"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-pantheon/fabrica-kit/profile"
	"github.com/go-pantheon/fabrica-kit/xcontext"
	"github.com/go-pantheon/fabrica-net/xnet"
	climod "github.com/go-pantheon/janus/gen/api/client/module"
	clipkt "github.com/go-pantheon/janus/gen/api/client/packet"
	cliseq "github.com/go-pantheon/janus/gen/api/client/sequence"
	"google.golang.org/protobuf/proto"
)

func Reply(netKind xnet.NetKind) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req any) (any, error) {
			if !profile.IsDev() {
				return handler(ctx, req)
			}

			logReply(ctx, netKind, req)

			return handler(ctx, req)
		}
	}
}

func logReply(ctx context.Context, netKind xnet.NetKind, reply any) {
	p := &clipkt.Packet{}

	if err := proto.Unmarshal(reply.(xnet.Pack), p); err != nil {
		log.Debugf("logReply: proto.Unmarshal failed: %v", err)
		return
	}

	if !profile.IsDev() {
		if p.Mod == int32(climod.ModuleID_System) && p.Seq == int32(cliseq.SystemSeq_Heartbeat) {
			return
		}
	}

	var (
		uid   string
		sid   string
		obj   string
		mod   string
		seq   string
		index string
		size  string
	)

	if md, ok := metadata.FromServerContext(ctx); ok {
		uid = md.Get(xcontext.CtxUID)
		sid = md.Get(xcontext.CtxSID)
	}

	obj = strconv.FormatInt(p.Obj, 10)
	mod = strconv.FormatInt(int64(p.Mod), 10)
	seq = strconv.FormatInt(int64(p.Seq), 10)
	index = strconv.FormatInt(int64(p.Index), 10)
	size = strconv.FormatInt(int64(len(p.Data)), 10)

	kv := make([]any, 0, 16)

	traceID, spanID := getTrace(ctx)

	kv = append(kv, "kind", "reply",
		"net", netKind,
		"trace", traceID,
		"span", spanID,
		"uid", uid,
		"sid", sid,
		"obj", obj,
		"mod", mod,
		"seq", seq,
		"index", index,
		"size", size,
		"latency", 0,
	)

	log.Debugw(kv...)
}
