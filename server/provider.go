package server

import (
	"github.com/hotstone-seo/hotstone-seo/internal/provider"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"go.uber.org/dig"
)

// Provider side
type Provider struct {
	dig.In
	provider.Controller
}

// SetRoute for Provider
func (p *Provider) SetRoute(e *echo.Echo) {
	group := e.Group("/p")
	group.Use(p.Controller.AuthMiddleware())
	group.Use(middleware.CORSWithConfig(middleware.DefaultCORSConfig))
	group.Use(middleware.Recover())

	p.Route(group)
}
