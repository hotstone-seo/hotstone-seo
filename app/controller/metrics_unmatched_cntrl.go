package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"

	"github.com/hotstone-seo/hotstone-seo/app/repository"
	"github.com/hotstone-seo/hotstone-seo/app/service"
	"go.uber.org/dig"
	"gopkg.in/go-playground/validator.v9"
)

// MetricsUnmatchedCntrl is controller to metrics_unmatched entity
type MetricsUnmatchedCntrl struct {
	dig.In
	service.MetricsUnmatchedService
}

// Route to define API Route
func (c *MetricsUnmatchedCntrl) Route(e *echo.Echo) {
	e.GET("metrics_unmatched", c.List)
}

// Create metrics_unmatched
func (c *MetricsUnmatchedCntrl) Create(ctx echo.Context) (err error) {
	var metrics_unmatched repository.MetricsUnmatched
	var lastInsertID int64
	ctx0 := ctx.Request().Context()
	if err = ctx.Bind(&metrics_unmatched); err != nil {
		return err
	}
	if err = validator.New().Struct(metrics_unmatched); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if lastInsertID, err = c.MetricsUnmatchedService.Insert(ctx0, metrics_unmatched); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}
	return ctx.JSON(http.StatusCreated, GeneralResponse{
		Message: fmt.Sprintf("Success insert new metrics_unmatched #%d", lastInsertID),
	})
}

// List of metrics_unmatched
func (c *MetricsUnmatchedCntrl) List(ctx echo.Context) (err error) {
	var metrics_unmatcheds []*repository.MetricsMismatchedCount
	ctx0 := ctx.Request().Context()
	if metrics_unmatcheds, err = c.MetricsUnmatchedService.ListCount(ctx0); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, metrics_unmatcheds)
}

// Get metrics_unmatched
func (c *MetricsUnmatchedCntrl) Get(ctx echo.Context) (err error) {
	var id int64
	var metrics_unmatched *repository.MetricsUnmatched
	ctx0 := ctx.Request().Context()
	if id, err = strconv.ParseInt(ctx.Param("id"), 10, 64); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}
	if metrics_unmatched, err = c.MetricsUnmatchedService.Find(ctx0, id); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if metrics_unmatched == nil {
		return echo.NewHTTPError(http.StatusNotFound, fmt.Sprintf("MetricsUnmatched#%d not found", id))
	}
	return ctx.JSON(http.StatusOK, metrics_unmatched)
}

// Delete metrics_unmatched
func (c *MetricsUnmatchedCntrl) Delete(ctx echo.Context) (err error) {
	var id int64
	ctx0 := ctx.Request().Context()
	if id, err = strconv.ParseInt(ctx.Param("id"), 10, 64); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}
	if err = c.MetricsUnmatchedService.Delete(ctx0, id); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, GeneralResponse{
		Message: fmt.Sprintf("Success delete metrics_unmatched #%d", id),
	})
}

// Update metrics_unmatched
func (c *MetricsUnmatchedCntrl) Update(ctx echo.Context) (err error) {
	var metrics_unmatched repository.MetricsUnmatched
	ctx0 := ctx.Request().Context()
	if err = ctx.Bind(&metrics_unmatched); err != nil {
		return err
	}
	if metrics_unmatched.ID <= 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}
	if err = validator.New().Struct(metrics_unmatched); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err = c.MetricsUnmatchedService.Update(ctx0, metrics_unmatched); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, GeneralResponse{
		Message: fmt.Sprintf("Success update metrics_unmatched #%d", metrics_unmatched.ID),
	})
}
