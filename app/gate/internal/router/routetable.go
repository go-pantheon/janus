package router

import (
	"context"
	"time"

	"github.com/pkg/errors"
)

type RouteTable struct {
	routetable.RouteTable
}

func NewRouteTable() *RouteTable {
	return &RouteTable{
		RouteTable: NewRouteTable("gate"),
	}
}

func AddRouteTable(ctx context.Context, rt *RouteTable, color string, oid int64) error {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	oldAddr, err := rt.GetSet(ctx, color, oid, addr)
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

	err := rt.DelIfSame(ctx, color, uid, addr)
	if err != nil && !errors.Is(err, context.Canceled) {
		log.Errorf("[gate.RouteTable] del route table failed. color=%s oid=%d %+v", color, uid, err)
	}
	return nil
}
