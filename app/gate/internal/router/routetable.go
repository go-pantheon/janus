package router

import (
	"context"
	"time"

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
		RouteTable: routetable.NewMasterRouteTable(redis.New(d.Rdb), profile.ServiceName(), routetable.WithTTL(d.GatewayRouteTableAliveDuration)),
	}
}

func AddRouteTable(ctx context.Context, rt *RouteTable, color string, uid int64) error {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	return rt.Set(ctx, color, uid, profile.GRPCEndpoint())
}

// func DelRouteTable(ctx context.Context, rt *RouteTable, color string, uid int64) error {
// 	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
// 	defer cancel()

// 	return rt.DelIfSame(ctx, color, uid, profile.GRPCEndpoint())
// }
