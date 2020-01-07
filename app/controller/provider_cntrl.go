package controller

import (
	"net/http"

	"github.com/hotstone-seo/hotstone-server/app/service"
	"github.com/labstack/echo"
	"go.uber.org/dig"
)

// ProviderCntrl is controller for provider function
type ProviderCntrl struct {
	dig.In
	service.ProviderService
}

// Route to define API Route
func (c *ProviderCntrl) Route(e *echo.Echo) {
	e.POST("provider/matchRule", c.MatchRule)
	e.POST("provider/retrieveData", c.RetrieveData)
	e.GET("provider/tags", c.Tags)
}

// MatchRule to match rule
func (p *ProviderCntrl) MatchRule(c echo.Context) (err error) {
	var (
		req  service.MatchRuleRequest
		resp *service.MatchRuleResponse
		ctx  = c.Request().Context()
	)
	if err = c.Bind(&req); err != nil {
		return err
	}
	if resp, err = p.ProviderService.MatchRule(ctx, req); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}
	return c.JSON(http.StatusOK, resp)
}

func (p *ProviderCntrl) RetrieveData(c echo.Context) (err error) {
	var (
		req  service.RetrieveDataRequest
		data interface{}
		ctx  = c.Request().Context()
	)
	if err = c.Bind(&req); err != nil {
		return err
	}
	if data, err = p.ProviderService.RetrieveData(ctx, req); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}
	return c.JSON(http.StatusOK, data)
}

func (p *ProviderCntrl) Tags(c echo.Context) (err error) {
	var (
		tags []*service.InterpolatedTag
		ctx  = c.Request().Context()
	)
	if tags, err = p.ProviderService.Tags(ctx, service.ProvideTagsRequest{
		RuleID: 0, // TODO: get from body
	}); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}
	return c.JSON(http.StatusOK, tags)
}
