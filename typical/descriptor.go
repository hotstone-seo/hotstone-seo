package typical

import (
	"github.com/hotstone-seo/hotstone-seo/pkg/gsociallogin"
	"github.com/hotstone-seo/hotstone-seo/server"
	"github.com/typical-go/typical-go/pkg/typapp"
	"github.com/typical-go/typical-go/pkg/typbuildtool"
	"github.com/typical-go/typical-go/pkg/typcfg"
	"github.com/typical-go/typical-go/pkg/typcore"
	"github.com/typical-go/typical-go/pkg/typdocker"
	"github.com/typical-go/typical-go/pkg/typreadme"
	"github.com/typical-go/typical-rest-server/pkg/typpostgres"
	"github.com/typical-go/typical-rest-server/pkg/typredis"
)

var (
	serverApp = server.New()
	redis     = typredis.New()
	postgres  = typpostgres.New().
			WithDBName("hotstone").
			WithDockerImage("timescale/timescaledb:latest-pg11")
	socialLogin = gsociallogin.New()
)

// Descriptor of hotstone-seo
var Descriptor = typcore.Descriptor{
	Name:    "hotstone-seo",
	Version: "0.0.1",

	App: typapp.Create(serverApp).
		WithModules(
			redis,
			postgres,
			socialLogin,
		),

	BuildTool: typbuildtool.
		Create(
			typbuildtool.StandardBuild(),
		).
		WithCommanders(
			postgres,
			typdocker.
				Create(
					redis,
					postgres,
					// prometheus.New(),
					// grafana.New(),
				),
			typreadme.Create(),
		),

	ConfigManager: typcfg.
		Create(
			serverApp,
			redis,
			postgres,
			socialLogin,
		),
}
