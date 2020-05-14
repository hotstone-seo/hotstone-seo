package controller

import (
	"net/http"

	"github.com/labstack/echo"

	"github.com/hotstone-seo/hotstone-seo/metric"
	"github.com/hotstone-seo/hotstone-seo/server/repository"
	"github.com/hotstone-seo/hotstone-seo/server/service"
	"go.uber.org/dig"
)

// MetricsCntrl is controller to metrics endpoint
type MetricsCntrl struct {
	dig.In
	service.MetricService
}

// Route to define API Route
func (c *MetricsCntrl) Route(e *echo.Group) {
	e.GET("/metrics/mismatched", c.ListMismatched)
	e.GET("/metrics/hit", c.CountHit)
	e.GET("/metrics/hit/range", c.ListCountHitPerDay)
	e.GET("/metrics/unique-page", c.CountUniquePage)
}

// ListMismatched of metrics_unmatched
func (c *MetricsCntrl) ListMismatched(ec echo.Context) (err error) {
	var report []*metric.MismatchReport
	ctx := ec.Request().Context()

	validCols := []string{"url", "first_seen", "last_seen", "count"}
	paginationParam := repository.BuildPaginationParam(ec.QueryParams(), validCols)

	if report, err = c.MetricService.MismatchReports(ctx, paginationParam); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return ec.JSON(http.StatusOK, report)
}

func (c *MetricsCntrl) CountHit(ec echo.Context) (err error) {
	var count int64
	ctx := ec.Request().Context()
	if count, err = c.MetricService.CountMatched(ctx, ec.QueryParams()); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return ec.JSON(http.StatusOK, map[string]int64{"count": count})
}

func (c *MetricsCntrl) CountUniquePage(ec echo.Context) (err error) {
	var count int64
	ctx := ec.Request().Context()
	if count, err = c.MetricService.CountUniquePage(ctx, ec.QueryParams()); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return ec.JSON(http.StatusOK, map[string]int64{"count": count})
}

func (c *MetricsCntrl) ListCountHitPerDay(ec echo.Context) (err error) {
	var counts []*metric.DailyReport

	ctx := ec.Request().Context()
	startDate := ec.QueryParam("start")
	endDate := ec.QueryParam("end")
	ruleID := ec.QueryParam("rule_id")

	if startDate == "" || endDate == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "'start' and 'end' query params are required")
	}

	if counts, err = c.MetricService.DailyReports(ctx, startDate, endDate, ruleID); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return ec.JSON(http.StatusOK, counts)
}
