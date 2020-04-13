package typical

import (
	"github.com/hotstone-seo/hotstone-seo/pkg/oauth2google"
	"github.com/hotstone-seo/hotstone-seo/server"
	"github.com/typical-go/typical-go/pkg/typapp"
	"github.com/typical-go/typical-go/pkg/typbuildtool"
	"github.com/typical-go/typical-go/pkg/typcore"
	"github.com/typical-go/typical-go/pkg/typdocker"
	"github.com/typical-go/typical-rest-server/pkg/typpostgres"
	"github.com/typical-go/typical-rest-server/pkg/typredis"
)

var _ = func() bool {
	typpostgres.DefaultDockerImage = "timescale/timescaledb:latest-pg11"
	typpostgres.DefaultDBName = "hotstone"

	return true
}()

// Descriptor of hotstone-seo
var Descriptor = typcore.Descriptor{
	Name:    "hotstone-seo",
	Version: "0.0.1",

	App: typapp.EntryPoint(server.Main, "server").
		Imports(
			server.Configuration(),
			typredis.Module(),
			typpostgres.Module(),
			oauth2google.Module(),
		),

	BuildTool: typbuildtool.
		BuildSequences(
			typbuildtool.StandardBuild(),
		).
		Utilities(
			typpostgres.Utility(),
			typredis.Utility(),
			typdocker.Compose(
				typredis.DockerRecipeV3(),
				typpostgres.DockerRecipeV3(),
			),
			typbuildtool.NewUtility(uiUtility),
			typbuildtool.NewUtility(jsonServer),
		),
}
