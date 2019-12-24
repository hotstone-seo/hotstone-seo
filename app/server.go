package app

import (
	"github.com/hotstone-seo/hotstone-server/app/config"
	"github.com/hotstone-seo/hotstone-server/app/controller"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"go.uber.org/dig"
)

type server struct {
	dig.In
	*echo.Echo
	config.Config
	controller.RuleCntrl
	controller.LocaleCntrl
	controller.DataSourceCntrl
	controller.TagCntrl
}

func (s *server) Middleware() {
	s.Use(middleware.CORSWithConfig(middleware.DefaultCORSConfig))
	s.Use(middleware.Recover())
}

func (s *server) Route() {
	s.RuleCntrl.Route(s.Echo)
	s.LocaleCntrl.Route(s.Echo)
	s.DataSourceCntrl.Route(s.Echo)
	s.TagCntrl.Route(s.Echo)
}

func (s *server) Start() error {
	return s.Echo.Start(s.Config.Address)
}
