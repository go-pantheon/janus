package player

import (
	"github.com/vulcan-frame/vulcan-gate/app/gate/internal/data"
	"github.com/vulcan-frame/vulcan-pkg-app/router/routetable"
	"github.com/vulcan-frame/vulcan-pkg-app/router/routetable/redis"
)

type RouteTable struct {
	routetable.RouteTable
}

func NewRouteTable(d *data.Data) *RouteTable {
	return &RouteTable{
		RouteTable: routetable.NewRouteTable("player", redis.NewRouteTable(d.Rdb)),
	}
}
