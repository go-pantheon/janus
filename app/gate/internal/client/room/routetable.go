package room

import (
	"github.com/go-pantheon/janus/app/gate/internal/data"
	"github.com/go-pantheon/vulcan-kit/router/routetable"
	"github.com/go-pantheon/vulcan-kit/router/routetable/redis"
)

type RouteTable struct {
	routetable.RouteTable
}

func NewRouteTable(d *data.Data) *RouteTable {
	return &RouteTable{
		RouteTable: routetable.NewRouteTable("room", redis.NewRedisRouteTable(d.Rdb)),
	}
}
