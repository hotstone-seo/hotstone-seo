package controller

import (
	"strconv"

	"github.com/hotstone-seo/hotstone-server/app/repository"
	"github.com/hotstone-seo/hotstone-server/app/service"
	"github.com/typical-go/typical-rest-server/pkg/utility/responsekit"
	"go.uber.org/dig"
	"gopkg.in/go-playground/validator.v9"
)

// RuleCntrl is controller to rule entity
type RuleCntrl struct {
	dig.In
	service.RuleService
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
