package typical

import (
	"github.com/hotstone-seo/hotstone-seo/pkg/oauth2google"
	"github.com/hotstone-seo/hotstone-seo/server"
	"github.com/typical-go/typical-go/pkg/typdocker"
	"github.com/typical-go/typical-go/pkg/typgo"
	"github.com/typical-go/typical-go/pkg/typmock"
	"github.com/typical-go/typical-rest-server/pkg/typpostgres"
	"github.com/typical-go/typical-rest-server/pkg/typredis"
)

var (
	mainDB = &typpostgres.Setting{
		DockerImage: "timescale/timescaledb:latest-pg11",
		DBName:      "hoststone",
	}
)

// Descriptor of hotstone-seo
var Descriptor = typgo.Descriptor{
	Name:    "hotstone-seo",
	Version: "0.0.1",

	EntryPoint: server.Main,

	Layouts: []string{
		"server",
		"pkg",
	},

	Configurer: typgo.Configurers{
		server.Configuration(),
		oauth2google.Configuration(),
		typredis.Configuration(),
		typpostgres.Configuration(mainDB),
	},

	Build: &typgo.StdBuild{},

	Utility: typgo.Utilities{
		typmock.Utility(),
		typpostgres.Utility(mainDB),
		typredis.Utility(),
		typgo.NewUtility(uiUtility),
		typgo.NewUtility(jsonServer),
		typdocker.Compose(
			typredis.DockerRecipeV3(),
			typpostgres.DockerRecipeV3(mainDB),
		),
	},
}
