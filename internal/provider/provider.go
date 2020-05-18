package provider

import (
	"github.com/hotstone-seo/hotstone-seo/server/controller"
	"github.com/labstack/echo"
	"go.uber.org/dig"
)

// Provider side
type Provider struct {
	dig.In
	controller.ProviderCntrl
}

// SetRoute provider
func (p *Provider) SetRoute(e *echo.Echo) {
	e.GET("p/match", p.MatchRule)
	e.GET("p/fetch-tags", p.FetchTag)
}
