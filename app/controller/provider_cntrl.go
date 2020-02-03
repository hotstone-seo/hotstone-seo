package controller

import (
	"net/http"
	"strconv"

	"github.com/hotstone-seo/hotstone-seo/app/service"
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
	e.POST("provider/tags", c.Tags)
	e.GET("provider/rule-tree", c.DumpRuleTree)
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
		data []byte
		ctx  = c.Request().Context()
	)
	if err = c.Bind(&req); err != nil {
		return
	}
	if data, err = p.ProviderService.RetrieveData(ctx, req, isUseCache(c.Request().Header)); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}
	return c.Blob(http.StatusOK, echo.MIMEApplicationJSON, data)
}

func (p *ProviderCntrl) Tags(c echo.Context) (err error) {
	var (
		req  service.ProvideTagsRequest
		tags []*service.InterpolatedTag
		ctx  = c.Request().Context()
	)
	if err = c.Bind(&req); err != nil {
		return
	}
	if tags, err = p.ProviderService.Tags(ctx, req, isUseCache(c.Request().Header)); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}
	return c.JSON(http.StatusOK, tags)
}

func (p *ProviderCntrl) DumpRuleTree(c echo.Context) (err error) {
	var (
		req     service.ProvideTagsRequest
		strTree string
		ctx     = c.Request().Context()
	)
	if err = c.Bind(&req); err != nil {
		return
	}
	if strTree, err = p.ProviderService.DumpRuleTree(ctx); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}
	return c.String(http.StatusOK, strTree)
}

func isUseCache(headers http.Header) bool {
	xCache := headers.Get("X-Cache")
	if xCache == "" {
		return true
	} else if xCache == "true" || xCache == "false" {
		boolVal, _ := strconv.ParseBool(xCache)
		return boolVal
	}

	return true
}
