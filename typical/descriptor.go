package typical

import (
	"github.com/hotstone-seo/hotstone-seo/internal/app"
	"github.com/hotstone-seo/hotstone-seo/internal/app/infra"
	"github.com/hotstone-seo/hotstone-seo/pkg/oauth2google"
	"github.com/typical-go/typical-go/pkg/typdocker"
	"github.com/typical-go/typical-go/pkg/typgo"
	"github.com/typical-go/typical-go/pkg/typmock"
)

var (

	// Descriptor of hotstone-seo
	Descriptor = typgo.Descriptor{
		Name:    "hotstone-seo",
		Version: "0.0.1",

		EntryPoint: app.Main,
		Layouts:    []string{"pkg", "internal"},

		Prebuild: &typgo.Prebuilds{
			&typgo.DependencyInjection{},
			&typgo.ConfigManager{
				Configs: []*typgo.Configuration{
					{Name: "APP", Spec: &infra.App{}},
					{Name: "REDIS", Spec: &infra.Redis{}},
					{Name: "PG", Spec: &infra.Pg{}},
					{Name: "ANALYT", Spec: &infra.Analyt{}},
					{Name: "OAUTH2_GOOGLE", Spec: &oauth2google.Config{}},
				},
			},
		},

		Compile: &typgo.StdCompile{},
		Run:     &typgo.StdRun{},
		Test:    &typgo.StdTest{},
		Clean:   &typgo.StdClean{},

		Utility: typgo.Utilities{
			&typmock.Utility{},
			&pgUtility{
				name:         "main-db",
				hostEnv:      "PG_HOST",
				portEnv:      "PG_PORT",
				userEnv:      "PG_USER",
				passwordEnv:  "PG_PASSWORD",
				dbnameEnv:    "PG_DBNAME",
				migrationSrc: "scripts/main-db/migration",
				seedSrc:      "scripts/main-db/seed",
			},
			&pgUtility{
				name:         "analyt-db",
				hostEnv:      "ANALYT_HOST",
				portEnv:      "ANALYT_PORT",
				userEnv:      "ANALYT_USER",
				passwordEnv:  "ANALYT_PASSWORD",
				dbnameEnv:    "ANALYT_DBNAME",
				migrationSrc: "scripts/analyt-db/migration",
				seedSrc:      "scripts/analyt-db/seed",
			},
			&redisUtility{},

			&typdocker.Utility{
				Version: typdocker.V3,
				Composers: []typdocker.Composer{
					&redisDocker{name: "redis01"},
					&pgDocker{
						name:        "ht-main",
						userEnv:     "PG_USER",
						passwordEnv: "PG_PASSWORD",
						portEnv:     "PG_PORT",
					},
					&pgDocker{
						name:        "ht-analyt",
						image:       "timescale/timescaledb:latest-pg11",
						userEnv:     "ANALYT_USER",
						passwordEnv: "ANALYT_PASSWORD",
						portEnv:     "ANALYT_PORT",
					},
				},
			},

			typgo.NewUtility(uiUtility),
			typgo.NewUtility(jsonServer),
		},
	}
)
