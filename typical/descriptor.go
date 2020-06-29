package typical

import (
	"github.com/hotstone-seo/hotstone-seo/internal/app"
	"github.com/hotstone-seo/hotstone-seo/internal/app/infra"
	"github.com/hotstone-seo/hotstone-seo/pkg/pgcmd"
	"github.com/hotstone-seo/hotstone-seo/pkg/rediscmd"
	"github.com/typical-go/typical-go/pkg/typdocker"
	"github.com/typical-go/typical-go/pkg/typgo"
	"github.com/typical-go/typical-go/pkg/typmock"
	"github.com/typical-go/typical-rest-server/pkg/dockerrx"
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
					{Name: "AUTH", Spec: &infra.Auth{}},
				},
			},
		},

		Compile: &typgo.StdCompile{},
		Run:     &typgo.StdRun{},
		Test:    &typgo.StdTest{},
		Clean:   &typgo.StdClean{},

		Utility: typgo.Utilities{
			&typmock.Utility{},
			&pgcmd.Utility{
				Name:         "main-db",
				HostEnv:      "PG_HOST",
				PortEnv:      "PG_PORT",
				UserEnv:      "PG_USER",
				PasswordEnv:  "PG_PASSWORD",
				DBNameEnv:    "PG_DBNAME",
				MigrationSrc: "databases/main-db/migration",
				SeedSrc:      "databases/main-db/seed",
			},
			&pgcmd.Utility{
				Name:         "analyt-db",
				HostEnv:      "ANALYT_HOST",
				PortEnv:      "ANALYT_PORT",
				UserEnv:      "ANALYT_USER",
				PasswordEnv:  "ANALYT_PASSWORD",
				DBNameEnv:    "ANALYT_DBNAME",
				MigrationSrc: "databases/analyt-db/migration",
				SeedSrc:      "databases/analyt-db/seed",
			},
			&rediscmd.Utility{
				Name:        "redis",
				HostEnv:     "REDIS_HOST",
				PortEnv:     "REDIS_PORT",
				PasswordEnv: "REDIS_PASSWORD",
			},
			&typdocker.Utility{
				Version: typdocker.V3,
				Composers: []typdocker.Composer{
					&dockerrx.RedisWithEnv{
						Name:        "redis01",
						PasswordEnv: "REDIS_PASSWORD",
						PortEnv:     "REDIS_PORT",
					},
					&dockerrx.PostgresWithEnv{
						Name:        "ht-main",
						UserEnv:     "PG_USER",
						PasswordEnv: "PG_PASSWORD",
						PortEnv:     "PG_PORT",
					},
					&dockerrx.PostgresWithEnv{
						Name:        "ht-analyt",
						Image:       "timescale/timescaledb:latest-pg11",
						UserEnv:     "ANALYT_USER",
						PasswordEnv: "ANALYT_PASSWORD",
						PortEnv:     "ANALYT_PORT",
					},
				},
			},

			typgo.NewUtility(uiUtility),
			typgo.NewUtility(jsonServer),
		},
	}
)
