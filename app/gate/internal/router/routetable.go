package router

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"github.com/vulcan-frame/vulcan-gate/app/gate/internal/data"
	"github.com/vulcan-frame/vulcan-pkg-app/profile"
	"github.com/vulcan-frame/vulcan-pkg-app/router/routetable"
	"github.com/vulcan-frame/vulcan-pkg-app/router/routetable/redis"
)

type RouteTable struct {
	routetable.RouteTable
}

func NewRouteTable(d *data.Data) *RouteTable {
	return &RouteTable{
		RouteTable: routetable.NewRouteTable("gate", redis.NewRouteTable(d.Rdb)),
	}
}

func AddRouteTable(ctx context.Context, rt *RouteTable, color string, oid int64) error {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	oldAddr, err := rt.GetSet(ctx, color, oid, profile.GRPCEndpoint())
	if err != nil {
		return errors.WithMessagef(err, "add route table failed. color=%s oid=%d", color, oid)
	}
	if len(oldAddr) > 0 {
		log.Debugf("[gate.RouteTable] found old route table on add. color=%s oid=%d addr=%s", color, oid, oldAddr)
	}
	return nil
}

func DelRouteTable(ctx context.Context, rt *RouteTable, color string, uid int64) error {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	err := rt.DelIfSame(ctx, color, uid, profile.GRPCEndpoint())
	if err != nil && !errors.Is(err, context.Canceled) {
		log.Errorf("[gate.RouteTable] del route table failed. color=%s oid=%d %+v", color, uid, err)
	}
	return nil
}
