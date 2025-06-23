package player

import (
	"github.com/go-pantheon/fabrica-kit/router/routetable"
	"github.com/go-pantheon/fabrica-kit/router/routetable/redis"
	"github.com/go-pantheon/janus/app/gate/internal/data"
)

type RouteTable struct {
	routetable.RouteTable
}

func NewRouteTable(d *data.Data) *RouteTable {
	return &RouteTable{
		RouteTable: routetable.NewRouteTable(routeTableName, redis.New(d.Rdb), routetable.WithTTL(d.AppRouteTableAliveDuration)),
	}
}
