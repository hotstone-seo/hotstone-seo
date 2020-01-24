package typical

import (
	"github.com/hotstone-seo/hotstone-seo/app"
	"github.com/typical-go/typical-go/pkg/typcore"
	"github.com/typical-go/typical-rest-server/pkg/typdocker"
	"github.com/typical-go/typical-rest-server/pkg/typpostgres"
	"github.com/typical-go/typical-rest-server/pkg/typreadme"
	"github.com/typical-go/typical-rest-server/pkg/typredis"
	"github.com/typical-go/typical-rest-server/pkg/typserver"
)

// Descriptor of hotstone-seo
var Descriptor = typcore.Descriptor{
	Name:    "hotstone-seo",
	Version: "0.0.1",
	Package: "github.com/hotstone-seo/hotstone-seo",

	App: typcore.NewApp(application).
		WithDependency(
			server,
			redis,
			postgres,
		).
		WithPrepare(
			redis,
			postgres,
		),

	Build: typcore.NewBuild().
		WithCommands(
			docker,
			readme,
			postgres,
			redis,
		),

	Configuration: typcore.NewConfiguration().
		WithConfigure(
			application,
			server,
			redis,
			postgres,
		),
}

var (
	application = app.New()
	readme      = typreadme.New()
	server      = typserver.New()
	redis       = typredis.New()
	postgres    = typpostgres.New().WithDBName("hotstone")

	docker = typdocker.New().
		WithComposers(
			redis,
			postgres,
			// prometheus.New(),
			// grafana.New(),
		)
)
