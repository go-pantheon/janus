package server

import (
	ketcd "github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-pantheon/fabrica-util/errors"
	"github.com/go-pantheon/janus/app/gate/internal/conf"
	"github.com/google/wire"
	etcd "go.etcd.io/etcd/client/v3"
)

var ProviderSet = wire.NewSet(NewTCPServer, NewGRPCServer, NewHTTPServer, NewRegistrar)

func NewRegistrar(conf *conf.Registry) (registry.Registrar, error) {
	client, err := etcd.New(etcd.Config{
		Endpoints: conf.Etcd.Endpoints,
		Username:  conf.Etcd.Username,
		Password:  conf.Etcd.Password,
	})
	if err != nil {
		return nil, errors.Wrap(err, "create etcd client failed")
	}

	return ketcd.New(client), nil
}
