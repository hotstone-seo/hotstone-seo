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
	s.Use(middleware.CORSWithConfig(middleware.DefaultCORSConfig))
	s.Use(middleware.Recover())
}

func (s *server) Route() {
	s.AuthCntrl.Route(s.Echo)
	s.RuleCntrl.Route(s.Echo)
	s.DataSourceCntrl.Route(s.Echo)
	s.TagCntrl.Route(s.Echo)
	s.ProviderCntrl.Route(s.Echo)
	s.CenterCntrl.Route(s.Echo)
	s.MetricsCntrl.Route(s.Echo)
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
