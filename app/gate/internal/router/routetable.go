package router

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-pantheon/fabrica-kit/profile"
	"github.com/go-pantheon/fabrica-kit/router/routetable"
	"github.com/go-pantheon/fabrica-kit/router/routetable/redis"
	"github.com/go-pantheon/janus/app/gate/internal/data"
)

type RouteTable struct {
	routetable.RouteTable
}

func NewRouteTable(d *data.Data) *RouteTable {
	return &RouteTable{
		RouteTable: routetable.NewRouteTable("gate", redis.New(d.Rdb)),
	}
}

func AddRouteTable(ctx context.Context, rt *RouteTable, color string, uid int64) error {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	oldAddr, err := rt.GetSet(ctx, color, uid, profile.GRPCEndpoint())
	if err != nil {
		return err
	}

	if len(oldAddr) > 0 && oldAddr != profile.GRPCEndpoint() {
		log.Warnf("[gate.RouteTable] found old route table on add. uid=%d color=%s addr=%s", uid, color, oldAddr)
	}

	return nil
}

func DelRouteTable(ctx context.Context, rt *RouteTable, color string, uid int64) error {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	return rt.DelIfSame(ctx, color, uid, profile.GRPCEndpoint())
}
