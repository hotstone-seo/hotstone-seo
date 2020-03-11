package typical

import (
	"github.com/hotstone-seo/hotstone-seo/pkg/gsociallogin"
	"github.com/hotstone-seo/hotstone-seo/server"
	"github.com/typical-go/typical-go/pkg/typdocker"
	"github.com/typical-go/typical-go/pkg/typreadme"
	"github.com/typical-go/typical-rest-server/pkg/typpostgres"
	"github.com/typical-go/typical-rest-server/pkg/typredis"
)

var (
	serverApp = server.New()

	readme = typreadme.New()

	redis = typredis.New()

	postgres = typpostgres.New().
			WithDBName("hotstone").
			WithDockerImage("timescale/timescaledb:latest-pg11")

	socialLogin = gsociallogin.New()

	docker = typdocker.New().
		WithComposers(
			redis,
			postgres,
			// prometheus.New(),
			// grafana.New(),
		)
)
