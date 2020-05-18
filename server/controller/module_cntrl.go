package controller

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/hotstone-seo/hotstone-seo/server/repository"
	"github.com/hotstone-seo/hotstone-seo/server/service"
	"github.com/labstack/echo"
	"go.uber.org/dig"
)

// ModuleCntrl is controller to module entity
type ModuleCntrl struct {
	dig.In
	service.ModuleService
}

// Route to define API Route
func (c *ModuleCntrl) Route(e *echo.Group) {
	e.GET("/modules", c.Find)
	e.GET("/modules/:id", c.FindOne)
}

// Find all modules
func (c *ModuleCntrl) Find(ctx echo.Context) (err error) {
	var modules []*repository.Module
	ctx0 := ctx.Request().Context()

	validCols := []string{"id", "name", "path", "pattern", "label", "updated_at", "created_at"}
	paginationParam := repository.BuildPaginationParam(ctx.QueryParams(), validCols)
	if modules, err = c.ModuleService.Find(ctx0, paginationParam); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, modules)
}

// FindOne module
func (c *ModuleCntrl) FindOne(ec echo.Context) (err error) {
	var (
		id     int64
		module *repository.Module
	)

	ctx := ec.Request().Context()
	if id, err = strconv.ParseInt(ec.Param("id"), 10, 64); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, "Invalid ID")
	}
	module, err = c.ModuleService.FindOne(ctx, id)
	if err == sql.ErrNoRows {
		return echo.NewHTTPError(http.StatusNotFound)
	}
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ec.JSON(http.StatusOK, module)
}
