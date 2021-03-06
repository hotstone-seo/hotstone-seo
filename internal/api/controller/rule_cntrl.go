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

// RuleCntrl is controller to rule entity
type RuleCntrl struct {
	dig.In
	service.RuleService
}

// Route to define API Route
func (c *RuleCntrl) Route(e *echo.Group) {
	e.GET("/rules", c.Find)
	e.POST("/rules", c.Create)
	e.GET("/rules/:id", c.FindOne)
	e.PUT("/rules", c.Update)
	e.PATCH("/rules/:id", c.Patch)
	e.DELETE("/rules/:id", c.Delete)
}

// Find all rule
func (c *RuleCntrl) Find(ctx echo.Context) (err error) {
	var rules []*repository.Rule
	ctx0 := ctx.Request().Context()

	validCols := []string{"id", "name", "url_pattern", "data_source_id", "updated_at", "created_at"}
	paginationParam := repository.BuildPaginationParam(ctx.QueryParams(), validCols)
	if rules, err = c.RuleService.Find(ctx0, paginationParam); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, rules)
}

// FindOne rule
func (c *RuleCntrl) FindOne(ec echo.Context) (err error) {
	var (
		id   int64
		rule *repository.Rule
	)

	ctx := ec.Request().Context()
	if id, err = strconv.ParseInt(ec.Param("id"), 10, 64); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, "Invalid ID")
	}
	rule, err = c.RuleService.FindOne(ctx, id)
	if err == sql.ErrNoRows {
		return echo.NewHTTPError(http.StatusNotFound)
	}
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ec.JSON(http.StatusOK, rule)
}

// Create rule
func (c *RuleCntrl) Create(ctx echo.Context) (err error) {
	var rule repository.Rule
	var lastInsertID int64
	ctx0 := ctx.Request().Context()
	if err = ctx.Bind(&rule); err != nil {
		return err
	}
	if err = rule.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if lastInsertID, err = c.RuleService.Insert(ctx0, rule); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}
	rule.ID = lastInsertID
	return ctx.JSON(http.StatusCreated, rule)
}

// Delete rule
func (c *RuleCntrl) Delete(ctx echo.Context) (err error) {
	var id int64
	ctx0 := ctx.Request().Context()
	if id, err = strconv.ParseInt(ctx.Param("id"), 10, 64); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}
	if err = c.RuleService.Delete(ctx0, id); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, GeneralResponse{
		Message: fmt.Sprintf("Success delete rule #%d", id),
	})
}

// Update rule
func (c *RuleCntrl) Update(ctx echo.Context) (err error) {
	var rule repository.Rule
	ctx0 := ctx.Request().Context()
	if err = ctx.Bind(&rule); err != nil {
		return err
	}
	if rule.ID <= 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}
	if err = rule.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err = c.RuleService.Update(ctx0, rule); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, GeneralResponse{
		Message: fmt.Sprintf("Success update rule #%d", rule.ID),
	})
}

func (c *RuleCntrl) Patch(ctx echo.Context) (err error) {
	var (
		id     int64
		fields map[string]interface{}
	)
	reqCtx := ctx.Request().Context()
	if id, err = strconv.ParseInt(ctx.Param("id"), 10, 64); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, "Invalid ID")
	}
	if err = ctx.Bind(&fields); err != nil {
		return err
	}
	if err = c.RuleService.Patch(reqCtx, id, fields); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, GeneralResponse{
		Message: fmt.Sprintf("Successfully patch rule #%d", id),
	})
}
