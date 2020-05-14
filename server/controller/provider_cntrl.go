package controller

import (
	"net/http"

	"github.com/hotstone-seo/hotstone-seo/pkg/cachekit"
	"github.com/hotstone-seo/hotstone-seo/server/service"
	"github.com/labstack/echo"
	"go.uber.org/dig"
)

// ProviderCntrl is controller for provider function
type ProviderCntrl struct {
	dig.In
	service.ProviderService
}

// MatchRule to match rule
func (p *ProviderCntrl) MatchRule(c echo.Context) (err error) {

	ctx := c.Request().Context()
	resp, err := p.ProviderService.Match(ctx, c.QueryParams())

	if err != nil {
		return httpError(err)
	}

	return c.JSON(http.StatusOK, resp)
}

// FetchTag to fetch the tag
func (p *ProviderCntrl) FetchTag(c echo.Context) (err error) {

	pragma := cachekit.CreatePragma(c.Request())
	tags, err := p.ProviderService.FetchTagsWithCache(
		c.Request().Context(),
		c.QueryParams(),
		pragma,
	)
	cachekit.SetHeader(c.Response(), pragma)

	if err != nil {
		return httpError(err)
	}

	return c.JSON(http.StatusOK, tags)
}
