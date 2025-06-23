package data

import (
	"time"

	cache "github.com/go-pantheon/fabrica-util/data/redis"
	"github.com/go-pantheon/janus/app/gate/internal/conf"
	"github.com/google/wire"
	"github.com/redis/go-redis/v9"
)

var ProviderSet = wire.NewSet(NewData)

type Data struct {
	Rdb                            redis.UniversalClient
	GatewayRouteTableAliveDuration time.Duration
	AppRouteTableAliveDuration     time.Duration
}

func NewData(c *conf.Data) (d *Data, cleanup func(), err error) {
	var rdb redis.UniversalClient

	if c.Redis.Cluster {
		rdb, cleanup, err = cache.NewCluster(&redis.ClusterOptions{
			Addrs:        []string{c.Redis.Addr},
			Password:     c.Redis.Password,
			DialTimeout:  c.Redis.DialTimeout.AsDuration(),
			WriteTimeout: c.Redis.WriteTimeout.AsDuration(),
			ReadTimeout:  c.Redis.ReadTimeout.AsDuration(),
		})
	} else {
		rdb, cleanup, err = cache.NewStandalone(&redis.Options{
			Addr:         c.Redis.Addr,
			Password:     c.Redis.Password,
			DialTimeout:  c.Redis.DialTimeout.AsDuration(),
			WriteTimeout: c.Redis.WriteTimeout.AsDuration(),
			ReadTimeout:  c.Redis.ReadTimeout.AsDuration(),
		})
	}

	if err != nil {
		return nil, nil, err
	}

	return &Data{
		Rdb:                            rdb,
		GatewayRouteTableAliveDuration: c.GatewayRouteTableAliveDuration.AsDuration(),
		AppRouteTableAliveDuration:     c.AppRouteTableAliveDuration.AsDuration(),
	}, cleanup, nil
}
