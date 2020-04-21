package controller

import (
	"net/http"

	"github.com/hotstone-seo/hotstone-seo/pkg/errkit"

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

	if errkit.IsValidationErr(err) {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, resp)
}

// FetchTag to fetch the tag
func (p *ProviderCntrl) FetchTag(c echo.Context) (err error) {
	var (
		tags []*service.ITag
	)

	ctx := c.Request().Context()
	pragma := cachekit.CreatePragma(c.Request())

	tags, err = p.ProviderService.FetchTagsWithCache(ctx, c.QueryParams(), pragma)

	if errkit.IsValidationErr(err) {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	if cachekit.NotModifiedError(err) {
		cachekit.SetHeader(c.Response(), pragma)
		return echo.NewHTTPError(http.StatusNotModified)
	}

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	cachekit.SetHeader(c.Response(), pragma)
	return c.JSON(http.StatusOK, tags)
}
