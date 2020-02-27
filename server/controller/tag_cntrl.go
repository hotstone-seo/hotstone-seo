package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/hotstone-seo/hotstone-seo/server/config"
	"github.com/hotstone-seo/hotstone-seo/server/repository"
	"github.com/hotstone-seo/hotstone-seo/server/service"
	"github.com/labstack/echo"
	"go.uber.org/dig"
	"gopkg.in/go-playground/validator.v9"
)

// TagCntrl is controller to tag entity
type TagCntrl struct {
	dig.In
	service.TagService
	config.Config
}

// Route to define API Route
func (c *TagCntrl) Route(e *echo.Group) {
	e.GET("/tags", c.Find)
	e.POST("/tags", c.Create)
	e.GET("/tags/:id", c.FindOne)
	e.PUT("/tags", c.Update)
	e.DELETE("/tags/:id", c.Delete)
}

// Create tag
func (c *TagCntrl) Create(ctx echo.Context) (err error) {
	var tag repository.Tag
	var lastInsertID int64
	ctx0 := ctx.Request().Context()
	if err = ctx.Bind(&tag); err != nil {
		return err
	}
	if err = validator.New().Struct(tag); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if lastInsertID, err = c.TagService.Insert(ctx0, tag); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}
	return ctx.JSON(http.StatusCreated, GeneralResponse{
		Message: fmt.Sprintf("Success insert new tag #%d", lastInsertID),
	})
}

// Find all tag
func (c *TagCntrl) Find(ctx echo.Context) (err error) {
	var (
		tags   []*repository.Tag
		filter repository.TagFilter
	)
	ctx0 := ctx.Request().Context()
	if ruleParam := ctx.QueryParam("rule_id"); ruleParam != "" {
		var ruleID int64
		if ruleID, err = strconv.ParseInt(ruleParam, 10, 64); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid Rule ID")
		}
		filter.RuleID = ruleID
	}
	filter.Locale = ctx.QueryParam("locale")
	if filter.Locale == "" {
		filter.Locale = c.Config.DefaultLocale
	}
	if tags, err = c.TagService.Find(ctx0, filter); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, tags)
}

// FindOne tag
func (c *TagCntrl) FindOne(ctx echo.Context) (err error) {
	var id int64
	var tag *repository.Tag
	ctx0 := ctx.Request().Context()
	if id, err = strconv.ParseInt(ctx.Param("id"), 10, 64); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}
	if tag, err = c.TagService.FindOne(ctx0, id); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if tag == nil {
		return echo.NewHTTPError(http.StatusNotFound, fmt.Sprintf("Tag#%d not found", id))
	}
	return ctx.JSON(http.StatusOK, tag)
}

// Delete tag
func (c *TagCntrl) Delete(ctx echo.Context) (err error) {
	var id int64
	ctx0 := ctx.Request().Context()
	if id, err = strconv.ParseInt(ctx.Param("id"), 10, 64); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}
	if err = c.TagService.Delete(ctx0, id); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, GeneralResponse{
		Message: fmt.Sprintf("Success delete tag #%d", id),
	})
}

// Update tag
func (c *TagCntrl) Update(ctx echo.Context) (err error) {
	var tag repository.Tag
	ctx0 := ctx.Request().Context()
	if err = ctx.Bind(&tag); err != nil {
		return err
	}
	if tag.ID <= 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}
	if err = validator.New().Struct(tag); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err = c.TagService.Update(ctx0, tag); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, GeneralResponse{
		Message: fmt.Sprintf("Success update tag #%d", tag.ID),
	})
}
