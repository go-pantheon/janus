package player

import (
	"github.com/go-pantheon/vulcan-kit/router/routetable"
	"github.com/go-pantheon/vulcan-kit/router/routetable/redis"
	"github.com/vulcan-frame/vulcan-gate/app/gate/internal/data"
)

type RouteTable struct {
	routetable.RouteTable
}

func NewRouteTable(d *data.Data) *RouteTable {
	return &RouteTable{
		RouteTable: routetable.NewRouteTable("player", redis.NewRedisRouteTable(d.Rdb)),
	}
}
