//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/vulcan-frame/vulcan-gate/app/gate/internal/client"
	"github.com/vulcan-frame/vulcan-gate/app/gate/internal/conf"
	"github.com/vulcan-frame/vulcan-gate/app/gate/internal/data"
	"github.com/vulcan-frame/vulcan-gate/app/gate/internal/intra/net/service"
	"github.com/vulcan-frame/vulcan-gate/app/gate/internal/server"
	"github.com/vulcan-frame/vulcan-gate/app/gate/internal/service/push"
	"github.com/vulcan-frame/vulcan-gate/pkg/net/health"
)

func initApp(*conf.Server, *conf.Label, *conf.Registry, *conf.Data, log.Logger, *health.Server) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, service.ProviderSet, push.ProviderSet, client.ProviderSet, newApp))
}
