package typical

import (
	"github.com/hotstone-seo/hotstone-server/app"
	"github.com/hotstone-seo/hotstone-server/pkg/grafana"
	"github.com/hotstone-seo/hotstone-server/pkg/prometheus"
	"github.com/typical-go/typical-go/pkg/typcore"
	"github.com/typical-go/typical-go/pkg/typrls"
	"github.com/typical-go/typical-rest-server/pkg/typdocker"
	"github.com/typical-go/typical-rest-server/pkg/typpostgres"
	"github.com/typical-go/typical-rest-server/pkg/typreadme"
	"github.com/typical-go/typical-rest-server/pkg/typredis"
	"github.com/typical-go/typical-rest-server/pkg/typserver"
)

// Descriptor of hotstone-server
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
			prometheus.New(),
			grafana.New(),
		)

	Descriptor = typcore.ProjectDescriptor{
		Name:    "hotstone-server",
		Version: "0.0.1",
		Package: "github.com/hotstone-seo/hotstone-server",

		App: typcore.NewApp().
			WithEntryPoint(application).
			WithProvide(
				server,
				redis,
				postgres,
			).
			WithDestroy(
				server,
				redis,
				postgres,
			).
			WithPrepare(
				redis,
				postgres,
			),

		BuildCommands: []typcore.BuildCommander{
			docker,
			readme,
			postgres,
			redis,
		},

		Configuration: typcore.NewConfiguration().
			WithConfigure(
				application,
				server,
				redis,
				postgres,
			),

		Releaser: typrls.New(),
	}
)
