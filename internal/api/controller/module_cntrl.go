package controller

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/hotstone-seo/hotstone-seo/internal/api/repository"
	"github.com/hotstone-seo/hotstone-seo/internal/api/service"
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
	e.POST("/modules", c.Create)
	e.GET("/modules/:id", c.FindOne)
	e.PUT("/modules", c.Update)
	e.DELETE("/modules/:id", c.Delete)
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

// Create module
func (c *ModuleCntrl) Create(ctx echo.Context) (err error) {
	var (
		req          service.ModuleRequest
		module       repository.Module
		lastInsertID int64
	)
	ctx0 := ctx.Request().Context()
	if err = ctx.Bind(&req); err != nil {
		return
	}
	if lastInsertID, err = c.ModuleService.Insert(ctx0, req); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}
	module.ID = lastInsertID
	module.Name = req.Name
	return ctx.JSON(http.StatusCreated, module)
}

// Delete module
func (c *ModuleCntrl) Delete(ctx echo.Context) (err error) {
	var id int64
	ctx0 := ctx.Request().Context()
	if id, err = strconv.ParseInt(ctx.Param("id"), 10, 64); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}
	if err = c.ModuleService.Delete(ctx0, id); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, GeneralResponse{
		Message: fmt.Sprintf("Success delete module #%d", id),
	})
}

// Update module
func (c *ModuleCntrl) Update(ctx echo.Context) (err error) {
	var (
		req service.ModuleRequest
	)
	ctx0 := ctx.Request().Context()
	if err = ctx.Bind(&req); err != nil {
		return err
	}
	if req.ID <= 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}
	if err = c.ModuleService.Update(ctx0, req); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, GeneralResponse{
		Message: fmt.Sprintf("Success update module #%d", req.ID),
	})
}
