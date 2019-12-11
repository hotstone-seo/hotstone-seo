package controller

import (
	"net/http"
	"strconv"

	"github.com/hotstone-seo/hotstone-server/app/repository"
	"github.com/hotstone-seo/hotstone-server/app/service"
	"github.com/labstack/echo"
	"github.com/typical-go/typical-rest-server/pkg/utility/responsekit"
	"go.uber.org/dig"
	"gopkg.in/go-playground/validator.v9"
)

// RuleCntrl is controller to rule entity
type RuleCntrl struct {
	dig.In
	service.RuleService
}

// Route to define API Route
func (c *RuleCntrl) Route(e *echo.Echo) {
	e.GET("rules", c.List)
	e.POST("rules", c.Create)
	e.GET("rules/:id", c.Get)
	e.PUT("rules", c.Update)
	e.DELETE("rules/:id", c.Delete)
}

// List of rule
func (c *RuleCntrl) List(ctx echo.Context) (err error) {
	var rules []*repository.Rule
	ctx0 := ctx.Request().Context()
	if rules, err = c.RuleService.List(ctx0); err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, rules)
}

// Get rule
func (c *RuleCntrl) Get(ctx echo.Context) (err error) {
	var id int64
	var rule *repository.Rule
	ctx0 := ctx.Request().Context()
	if id, err = strconv.ParseInt(ctx.Param("id"), 10, 64); err != nil {
		return responsekit.InvalidID(ctx, err)
	}
	if rule, err = c.RuleService.Find(ctx0, id); err != nil {
		return err
	}
	if rule == nil {
		return responsekit.NotFound(ctx, id)
	}
	return ctx.JSON(http.StatusOK, rule)
}

// Create rule
func (c *RuleCntrl) Create(ctx echo.Context) (err error) {
	var rule repository.Rule
	var lastInsertID int64
	ctx0 := ctx.Request().Context()
	if err = ctx.Bind(&rule); err != nil {
		return err
	}
	if err = validator.New().Struct(rule); err != nil {
		return responsekit.InvalidRequest(ctx, err)
	}
	if lastInsertID, err = c.RuleService.Insert(ctx0, rule); err != nil {
		return err
	}
	return responsekit.InsertSuccess(ctx, lastInsertID)
}

// Delete rule
func (c *RuleCntrl) Delete(ctx echo.Context) (err error) {
	var id int64
	ctx0 := ctx.Request().Context()
	if id, err = strconv.ParseInt(ctx.Param("id"), 10, 64); err != nil {
		return responsekit.InvalidID(ctx, err)
	}
	if err = c.RuleService.Delete(ctx0, id); err != nil {
		return err
	}
	return responsekit.DeleteSuccess(ctx, id)
}

// Update rule
func (c *RuleCntrl) Update(ctx echo.Context) (err error) {
	var rule repository.Rule
	ctx0 := ctx.Request().Context()
	if err = ctx.Bind(&rule); err != nil {
		return err
	}
	if rule.ID <= 0 {
		return responsekit.InvalidID(ctx, err)
	}
	if err = validator.New().Struct(rule); err != nil {
		return responsekit.InvalidRequest(ctx, err)
	}
	if err = c.RuleService.Update(ctx0, rule); err != nil {
		return err
	}
	return responsekit.UpdateSuccess(ctx, rule.ID)
}
