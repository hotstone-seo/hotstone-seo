package controller

import (
	"net/http"
	"strconv"

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

// RetrieveData retrieve data
func (p *ProviderCntrl) RetrieveData(c echo.Context) (err error) {
	var (
		req  service.RetrieveDataRequest
		resp *service.RetrieveDataResponse
		ctx  = c.Request().Context()
	)
	if err = c.Bind(&req); err != nil {
		return
	}
	if resp, err = p.ProviderService.RetrieveData(ctx, req, useCache(c)); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}
	return c.Blob(http.StatusOK, echo.MIMEApplicationJSON, resp.Data)
}

// Tags to get tag
func (p *ProviderCntrl) Tags(c echo.Context) (err error) {
	var (
		req  service.ProvideTagsRequest
		tags []*service.InterpolatedTag
		ctx  = c.Request().Context()
	)
	if err = c.Bind(&req); err != nil {
		return
	}
	if tags, err = p.ProviderService.Tags(ctx, req, useCache(c)); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}
	return c.JSON(http.StatusOK, tags)
}

// DumpRuleTree to dump rule tree
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

func useCache(c echo.Context) bool {
	xCache := c.Request().Header.Get("X-Cache")
	if xCache == "" {
		return true
	} else if xCache == "true" || xCache == "false" {
		boolVal, _ := strconv.ParseBool(xCache)
		return boolVal
	}

	return true
}
