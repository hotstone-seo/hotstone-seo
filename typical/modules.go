package typical

import (
	"github.com/hotstone-seo/hotstone-seo/app"
	"github.com/typical-go/typical-rest-server/pkg/typdocker"
	"github.com/typical-go/typical-rest-server/pkg/typpostgres"
	"github.com/typical-go/typical-rest-server/pkg/typreadme"
	"github.com/typical-go/typical-rest-server/pkg/typredis"
	"github.com/typical-go/typical-rest-server/pkg/typserver"
)

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
