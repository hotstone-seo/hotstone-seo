package app

import (
	"github.com/hotstone-seo/hotstone-seo/app/config"
	"github.com/hotstone-seo/hotstone-seo/app/controller"
	"github.com/juju/errors"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	log "github.com/sirupsen/logrus"

	"go.uber.org/dig"
)

type server struct {
	dig.In
	*echo.Echo
	config.Config
	controller.AuthCntrl
	controller.RuleCntrl
	controller.DataSourceCntrl
	controller.TagCntrl
	controller.ProviderCntrl
	controller.CenterCntrl
	controller.MetricsCntrl
}

func (s *server) Middleware() {
	s.GET("auth/google/login", s.AuthCntrl.AuthGoogleLogin)
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
}

func (s *server) Route() {
	// s.AuthCntrl.Route(s.Echo)
	// s.RuleCntrl.Route(s.Echo)
	// s.DataSourceCntrl.Route(s.Echo)
	// s.TagCntrl.Route(s.Echo)
	// s.ProviderCntrl.Route(s.Echo)
	// s.CenterCntrl.Route(s.Echo)
	// s.MetricsCntrl.Route(s.Echo)
}

func (s *server) ErrorHandler() {
	s.HTTPErrorHandler = func(err error, c echo.Context) {
		s.DefaultHTTPErrorHandler(err, c)
		log.Print(errors.Details(err))
	}
}

func (s *server) Start() error {
	return s.Echo.Start(s.Config.Address)
}
