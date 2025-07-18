//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-pantheon/fabrica-net/http/health"
	"github.com/go-pantheon/janus/app/gate/internal/broadcast"
	"github.com/go-pantheon/janus/app/gate/internal/client"
	"github.com/go-pantheon/janus/app/gate/internal/conf"
	"github.com/go-pantheon/janus/app/gate/internal/data"
	"github.com/go-pantheon/janus/app/gate/internal/intra/net/service"
	"github.com/go-pantheon/janus/app/gate/internal/server"
	"github.com/go-pantheon/janus/app/gate/internal/service/push"
	"github.com/google/wire"
)

func initApp(*conf.Server, *conf.Label, *conf.Registry, *conf.Data, log.Logger, *health.Server) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, service.ProviderSet, push.ProviderSet, client.ProviderSet, broadcast.ProviderSet, newApp))
}
