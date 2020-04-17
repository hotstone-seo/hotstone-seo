package controller

import (
	"net/http"

	"github.com/labstack/echo"

	"github.com/hotstone-seo/hotstone-seo/server/metric"
	"github.com/hotstone-seo/hotstone-seo/server/repository"
	"go.uber.org/dig"
)

// MetricsCntrl is controller to metrics endpoint
type MetricsCntrl struct {
	dig.In
	metric.MetricsRuleMatchingService
}

// Route to define API Route
func (c *MetricsCntrl) Route(e *echo.Group) {
	e.GET("/metrics/mismatched", c.ListMismatched)
	e.GET("/metrics/hit", c.CountHit)
	e.GET("/metrics/hit/range", c.ListCountHitPerDay)
	e.GET("/metrics/unique-page", c.CountUniquePage)
}

// ListMismatched of metrics_unmatched
func (c *MetricsCntrl) ListMismatched(ctx echo.Context) (err error) {
	var metrics_mismatcheds []*metric.MetricsMismatchedCount
	ctx0 := ctx.Request().Context()

	validCols := []string{"url", "first_seen", "last_seen", "count"}
	paginationParam := repository.BuildPaginationParam(ctx.QueryParams(), validCols)

	if metrics_mismatcheds, err = c.MetricsRuleMatchingService.ListMismatchedCount(ctx0, paginationParam); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, metrics_mismatcheds)
}

func (c *MetricsCntrl) CountHit(ctx echo.Context) (err error) {
	var count int64
	ctx0 := ctx.Request().Context()
	if count, err = c.MetricsRuleMatchingService.CountMatched(ctx0, ctx.QueryParams()); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, map[string]int64{"count": count})
}

func (c *MetricsCntrl) CountUniquePage(ctx echo.Context) (err error) {
	var count int64
	ctx0 := ctx.Request().Context()
	if count, err = c.MetricsRuleMatchingService.CountUniquePage(ctx0, ctx.QueryParams()); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, map[string]int64{"count": count})
}

func (c *MetricsCntrl) ListCountHitPerDay(ctx echo.Context) (err error) {
	var counts []*metric.MetricsCountHitPerDay
	ctx0 := ctx.Request().Context()

	startDate := ctx.QueryParam("start")
	endDate := ctx.QueryParam("end")
	ruleID := ctx.QueryParam("rule_id")

	if startDate == "" || endDate == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "'start' and 'end' query params are required")
	}

	if counts, err = c.MetricsRuleMatchingService.ListCountHitPerDay(ctx0, startDate, endDate, ruleID); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, counts)
}
