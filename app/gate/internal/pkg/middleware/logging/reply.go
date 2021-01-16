package logging

import (
	"context"
	"strconv"

	climod "github.com/vulcan-frame/vulcan-gate/gen/api/client/module"
	cliseq "github.com/vulcan-frame/vulcan-gate/gen/api/client/sequence"
	"github.com/vulcan-frame/vulcan-gate/pkg/net"
	vctx "github.com/vulcan-frame/vulcan-gate/pkg/net/context"
	"github.com/vulcan-frame/vulcan-pkg-app/profile"
	"google.golang.org/protobuf/proto"
)

func Reply(netKind net.NetKind) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			if !profile.IsDev() {
				return handler(ctx, req)
			}

			logReply(ctx, netKind, req)
			return handler(ctx, req)
		}
	}
}

func logReply(ctx context.Context, netKind net.NetKind, reply interface{}) {
	if err := proto.Unmarshal(reply.([]byte), p); err != nil {
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
	kv = append(kv, "kind", "reply",
		"net", netKind,
		"uid", uid,
		"sid", sid,
		"obj", obj,
		"mod", mod,
		"seq", seq,
		"index", index,
		"latency", 0,
	)

	log.Debugw(kv...)
}
