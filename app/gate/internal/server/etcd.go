package server

import (
	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/vulcan-frame/vulcan-gate/app/gate/internal/conf"
)

var ProviderSet = wire.NewSet(NewTCPServer, NewGRPCServer, NewHTTPServer, NewRegistrar)

func NewRegistrar(conf *conf.Registry) (registry.Registrar, error) {
	client, err := etcdclient.New(etcdclient.Config{
		Endpoints: conf.Etcd.Endpoints,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "create etcd client failed")
	}

	return etcd.New(client), nil
}
