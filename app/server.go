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
}

func startServer(s server) error {
	s.Use(middleware.Recover())

	s.RuleCntrl.Route(s.Echo)

	return s.Echo.Start(s.Config.Address)
}
