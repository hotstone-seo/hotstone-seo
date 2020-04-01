package server

import (
	"github.com/go-redis/redis"
	"github.com/hotstone-seo/hotstone-seo/server/config"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/typical-go/typical-rest-server/pkg/typpostgres"
	"github.com/typical-go/typical-rest-server/pkg/typserver"

	log "github.com/sirupsen/logrus"

	"go.uber.org/dig"
)

type server struct {
	dig.In
	*typserver.Server
	*config.Config

	API      api
	Provider provider

	Postgres *typpostgres.DB
	Redis    *redis.Client
}

func startServer(s server) error {
	s.SetLogger(s.Debug)

	s.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:  "build",
		HTML5: true,
	}))

	// health check
	s.PutHealthChecker("postgres", s.Postgres.Ping)
	s.PutHealthChecker("redis", s.Redis.Ping().Err)

	s.HTTPErrorHandler = func(err error, c echo.Context) {
		s.DefaultHTTPErrorHandler(err, c)
		log.Error(err.Error())
	}

	s.API.route(s)
	s.Provider.route(s)

	return s.Start(s.Address)
}
