package controller

import (
	"net/http"

	"github.com/labstack/echo"

	"github.com/hotstone-seo/hotstone-seo/app/repository"
	"github.com/hotstone-seo/hotstone-seo/app/service"
	"go.uber.org/dig"
)

// MetricsCntrl is controller to metrics endpoint
type MetricsCntrl struct {
	dig.In
	service.MetricsRuleMatchingService
}

// Route to define API Route
func (c *MetricsCntrl) Route(e *echo.Echo) {

	e.GET("metrics/mismatched", c.ListMismatched)
}

// ListMismatched of metrics_unmatched
func (c *MetricsCntrl) ListMismatched(ctx echo.Context) (err error) {
	var metrics_mismatcheds []*repository.MetricsMismatchedCount
	ctx0 := ctx.Request().Context()
	if metrics_mismatcheds, err = c.MetricsRuleMatchingService.ListMismatchedCount(ctx0); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, metrics_mismatcheds)
}
