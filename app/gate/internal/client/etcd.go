package client

import (
	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/vulcan-frame/vulcan-gate/app/gate/internal/client/player"
	"github.com/vulcan-frame/vulcan-gate/app/gate/internal/conf"
	gate "github.com/vulcan-frame/vulcan-gate/app/gate/internal/router"
)

var ProviderSet = wire.NewSet(
	NewDiscovery,
	player.NewRouteTable, player.NewConn, player.NewClient,
	gate.NewRouteTable,
)

func NewDiscovery(conf *conf.Registry) (registry.Discovery, error) {
	client, err := etcdclient.New(etcdclient.Config{
		Endpoints: conf.Etcd.Endpoints,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "new etcdclient failed")
	}

	r := etcd.New(client)
	return r, nil
}
