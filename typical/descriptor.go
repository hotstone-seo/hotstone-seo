package typical

import (
	"github.com/hotstone-seo/hotstone-seo/internal/app"
	"github.com/hotstone-seo/hotstone-seo/pkg/oauth2google"
	"github.com/typical-go/typical-go/pkg/typdocker"
	"github.com/typical-go/typical-go/pkg/typgo"
	"github.com/typical-go/typical-go/pkg/typmock"
	"github.com/typical-go/typical-rest-server/pkg/typpg"
	"github.com/typical-go/typical-rest-server/pkg/typredis"
)

var (
	mainDB = typpg.Init(&typpg.Settings{
		DockerName:   "ht-main",
		DBName:       "hotstone",
		UtilityCmd:   "main-db",
		MigrationSrc: "scripts/main-db/migration",
		SeedSrc:      "scripts/main-db/seed",
	})

	analytDB = typpg.Init(&typpg.Settings{
		Ctor:         "analyt",
		ConfigName:   "ANALYT",
		DockerImage:  "timescale/timescaledb:latest-pg11",
		DockerName:   "ht-analyt",
		DBName:       "hotstone_analyt",
		UtilityCmd:   "analyt-db",
		MigrationSrc: "scripts/analyt-db/migration",
		SeedSrc:      "scripts/analyt-db/seed",
		Port:         5433,
	})

	redis = typredis.Init(&typredis.Settings{})

	// Descriptor of hotstone-seo
	Descriptor = typgo.Descriptor{
		Name:    "hotstone-seo",
		Version: "0.0.1",

		EntryPoint: app.Main,

		Layouts: []string{
			"pkg",
			"internal",
		},

		Configurer: typgo.Configurers{
			app.Configuration(),
			oauth2google.Configuration(),
			typredis.Configuration(redis),
			typpg.Configuration(mainDB),
			typpg.Configuration(analytDB),
		},

		Build: &typgo.StdBuild{},

		Utility: typgo.Utilities{
			typmock.Utility(),
			typpg.Utility(mainDB),
			typpg.Utility(analytDB),
			typredis.Utility(redis),

			typdocker.Compose(
				typredis.DockerRecipeV3(redis),
				typpg.DockerRecipeV3(mainDB),
				typpg.DockerRecipeV3(analytDB),
			),

			typgo.NewUtility(uiUtility),
			typgo.NewUtility(jsonServer),
		},
	}
)
