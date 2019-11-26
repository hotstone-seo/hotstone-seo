package app

import (
	"github.com/hotstone-seo/hotstone-server/app/config"
	"github.com/labstack/echo"
	"go.uber.org/dig"
)

type server struct {
	dig.In
	*echo.Echo
	config.Config
}

func startServer(s server) error {
	return s.Start(s.Address)
}
