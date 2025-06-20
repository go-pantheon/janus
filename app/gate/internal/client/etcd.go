package client

import (
	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-pantheon/janus/app/gate/internal/client/player"
	"github.com/go-pantheon/janus/app/gate/internal/client/room"
	"github.com/go-pantheon/janus/app/gate/internal/conf"
	gate "github.com/go-pantheon/janus/app/gate/internal/router"
	"github.com/google/wire"
	"github.com/pkg/errors"
	etcdclient "go.etcd.io/etcd/client/v3"
)

var ProviderSet = wire.NewSet(
	NewDiscovery,
	player.NewRouteTable, player.NewConn, player.NewClient,
	room.NewRouteTable, room.NewConn, room.NewClient,
	gate.NewRouteTable,
)

func NewDiscovery(conf *conf.Registry) (registry.Discovery, error) {
	client, err := etcdclient.New(etcdclient.Config{
		Endpoints: conf.Etcd.Endpoints,
		Username:  conf.Etcd.Username,
		Password:  conf.Etcd.Password,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "new etcdclient failed")
	}

	return etcd.New(client), nil
}
