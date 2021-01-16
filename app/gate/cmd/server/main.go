package main

import (
	"flag"
	"path/filepath"

	"github.com/vulcan-frame/vulcan-gate/app/gate/internal/conf"
	"github.com/vulcan-frame/vulcan-gate/app/gate/internal/pkg/security"
	tcp "github.com/vulcan-frame/vulcan-gate/pkg/net/tcp/server"
	"github.com/vulcan-frame/vulcan-pkg-app/profile"
)

var (
	flagConf string
)

func init() {
	flag.StringVar(&flagConf, "conf", "app/gate/configs", "config path, eg: -conf config.yaml")
}

func newApp(logger log.Logger, ts *tcp.Server, hs *http.Server, gs *grpc.Server, health *health.Server,
	label *conf.Label, rr registry.Registrar,
) *kratos.App {
	md := map[string]string{
		profile.SERVICE: label.Service,
		profile.PROFILE: label.Profile,
		profile.VERSION: label.Version,
		profile.COLOR:   label.Color,
	}

	url, err := gs.Endpoint()
	if err != nil {
		panic(err)
	}

	profile.Init(label.Profile, label.Color, label.Zone, label.Version, label.Node, url)

	return kratos.New(
		kratos.Server(health, ts, hs, gs),
		kratos.Registrar(rr),
	)
}

func main() {
	flag.Parse()

	flagConf, err := filepath.Abs(flagConf)
	if err != nil {
		panic(err)
	}

	c := config.New(
		config.WithSource(
			env.NewSource(profile.ORG_PREFIX),
			file.NewSource(flagConf),
		),
	)
	if err := c.Load(); err != nil {
		panic(err)
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}

	var rc conf.Registry
	if err := c.Scan(&rc); err != nil {
		panic(err)
	}

	var sc conf.Secret
	if err := c.Scan(&sc); err != nil {
		panic(err)
	}
	if err := security.Init(sc.AesKey, sc.PrivateKey); err != nil {
		panic(err)
	}

	app, cleanup, err := initApp(bc.Server, bc.Label, &rc, bc.Data, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	log.Infof("[%s] is running", bc.Label.Service)

	if err = app.Run(); err != nil {
		panic(err)
	}
}
