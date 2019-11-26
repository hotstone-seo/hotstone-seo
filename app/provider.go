package app

import (
	"github.com/hotstone-seo/hotstone-server/app/config"
	"github.com/labstack/echo"
	"go.uber.org/dig"
)

type provider struct {
	dig.In
	*echo.Echo
	config.Config
}

func startProvider(p provider) error {
	return p.Start(p.Address)
}
