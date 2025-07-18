package main

import (
	"flag"
	"path/filepath"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/env"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-pantheon/fabrica-kit/metrics"
	"github.com/go-pantheon/fabrica-kit/profile"
	"github.com/go-pantheon/fabrica-kit/trace"
	"github.com/go-pantheon/fabrica-kit/xlog"
	"github.com/go-pantheon/fabrica-net/http/health"
	kcp "github.com/go-pantheon/fabrica-net/kcp/server"
	tcp "github.com/go-pantheon/fabrica-net/tcp/server"
	ws "github.com/go-pantheon/fabrica-net/websocket/server"
	"github.com/go-pantheon/fabrica-util/compress"
	"github.com/go-pantheon/fabrica-util/xtime"
	"github.com/go-pantheon/janus/app/gate/internal/broadcast"
	"github.com/go-pantheon/janus/app/gate/internal/conf"
	"github.com/go-pantheon/janus/app/gate/internal/pkg/security"
)

var (
	flagConf string
)

func init() {
	flag.StringVar(&flagConf, "conf", "app/gate/configs", "config path, eg: -conf config.yaml")
}

func newApp(logger log.Logger, b *broadcast.Broadcaster,
	ts *tcp.Server, ws *ws.Server, ks *kcp.Server, hs *http.Server, gs *grpc.Server, health *health.Server,
	label *conf.Label, rr registry.Registrar,
) *kratos.App {
	md := map[string]string{
		profile.ServiceKey: label.Service,
		profile.ProfileKey: label.Profile,
		profile.VersionKey: label.Version,
		profile.ColorKey:   label.Color,
	}

	url, err := gs.Endpoint()
	if err != nil {
		panic(err)
	}

	profile.SetGRPCEndpoint(url)

	if err := b.Start(); err != nil {
		panic(err)
	}

	return kratos.New(
		kratos.Name(label.Service),
		kratos.Version(label.Version),
		kratos.Metadata(md),
		kratos.Logger(logger),
		kratos.Server(health, ts, ws, ks, hs, gs),
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
			env.NewSource(profile.OrgPrefix),
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

	profile.Init(bc.Label.Service, bc.Label.Profile, bc.Label.Color, bc.Label.Zone, bc.Label.Version, bc.Label.Node)

	if err := xtime.InitSimple(bc.Label.Language); err != nil {
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

	if err := security.Init(&sc); err != nil {
		panic(err)
	}

	if err := trace.Init(bc.Trace.Endpoint, bc.Label.Service, bc.Label.Profile, bc.Label.Color); err != nil {
		panic(err)
	}

	metrics.Init(bc.Label.Service)
	compress.Init(bc.Compress.Weak, bc.Compress.Strong)

	logger := xlog.Init(bc.Log.Type, bc.Log.Level, bc.Label.Profile, bc.Label.Color, bc.Label.Service, bc.Label.Version, bc.Label.Node)

	app, cleanup, err := initApp(bc.Server, bc.Label, &rc, bc.Data, logger, health.NewServer(bc.Server.Health))
	if err != nil {
		panic(err)
	}
	defer cleanup()

	log.Infof("[%s] is running", bc.Label.Service)

	if err = app.Run(); err != nil {
		panic(err)
	}
}
