package server

import (
	"github.com/hotstone-seo/hotstone-seo/server/controller"
	"github.com/labstack/echo"
	"go.uber.org/dig"
)

type provider struct {
	dig.In
	controller.ProviderCntrl
}

func (p *provider) route(e *echo.Echo) {
	e.POST("p/match", p.MatchRule)
	e.GET("p/rules/:id/tags", p.FetchTag)
}
