package server

import (
	"github.com/go-redis/redis"
	"github.com/hotstone-seo/hotstone-seo/server/config"
	"github.com/hotstone-seo/hotstone-seo/server/controller"
	"github.com/juju/errors"
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
	config.Config
	controller.AuthCntrl
	controller.RuleCntrl
	controller.DataSourceCntrl
	controller.TagCntrl
	controller.ProviderCntrl
	controller.CenterCntrl
	controller.MetricsCntrl

	Postgres *typpostgres.DB
	Redis    *redis.Client
}

func startServer(s server) error {
	s.SetDebug(s.Debug)

	// health check
	s.PutHealthChecker("postgres", s.Postgres.Ping)
	s.PutHealthChecker("redis", s.Redis.Ping().Err)

	s.HTTPErrorHandler = func(err error, c echo.Context) {
		s.DefaultHTTPErrorHandler(err, c)
		log.Print(errors.Details(err))
	}

	s.POST("auth/google/login", s.AuthCntrl.AuthGoogleLogin)
	s.GET("auth/google/callback", s.AuthCntrl.AuthGoogleCallback)

	api := s.Group("/api")

	jwtCfg := middleware.DefaultJWTConfig
	jwtCfg.SigningKey = []byte(s.Config.JwtSecret)
	jwtCfg.TokenLookup = "cookie:secure_token"
	api.Use(middleware.JWTWithConfig(jwtCfg))

	api.Use(middleware.CORSWithConfig(middleware.DefaultCORSConfig))
	api.Use(middleware.Recover())

	api.POST("/logout", s.AuthCntrl.AuthLogout)

	s.RuleCntrl.Route(api)
	s.DataSourceCntrl.Route(api)
	s.TagCntrl.Route(api)
	s.ProviderCntrl.Route(api)
	s.CenterCntrl.Route(api)
	s.MetricsCntrl.Route(api)

	return s.Start(s.Address)
}
