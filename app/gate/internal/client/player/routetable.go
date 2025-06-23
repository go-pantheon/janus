package player

import (
	"github.com/go-pantheon/fabrica-kit/router/routetable"
	"github.com/go-pantheon/fabrica-kit/router/routetable/redis"
	"github.com/go-pantheon/janus/app/gate/internal/data"
)

type RouteTable struct {
	routetable.MasterRouteTable
}

func NewRouteTable(d *data.Data) *RouteTable {
	return &RouteTable{
		MasterRouteTable: routetable.NewMasterRouteTable(redis.New(d.Rdb), serviceName, routetable.WithTTL(d.AppRouteTableAliveDuration)),
	}
}
