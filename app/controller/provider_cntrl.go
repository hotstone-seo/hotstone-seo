package controller

import (
	"net/http"

	"github.com/hotstone-seo/hotstone-server/app/repository"
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
func (c *ProviderCntrl) MatchRule(ctx echo.Context) (err error) {
	var (
		req  service.MatchRuleRequest
		rule *repository.Rule
	)
	if err = ctx.Bind(&req); err != nil {
		return err
	}
	if rule, err = c.ProviderService.MatchRule(req); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}
	return ctx.JSON(http.StatusOK, rule)
}

func (c *ProviderCntrl) RetrieveData(ctx echo.Context) (err error) {
	var (
		req  service.RetrieveDataRequest
		data interface{}
	)
	if err = ctx.Bind(&req); err != nil {
		return err
	}
	if data, err = c.ProviderService.RetrieveData(req); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}
	return ctx.JSON(http.StatusOK, data)
}

func (c *ProviderCntrl) Tags(ctx echo.Context) (err error) {
	var tags []*repository.Tag
	ruleID := ctx.QueryParam("ruleID")
	if tags, err = c.ProviderService.Tags(ruleID, nil); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}
	return ctx.JSON(http.StatusOK, tags)
}
