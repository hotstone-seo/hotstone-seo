package controller

import (
	"fmt"
	"net/http"

	"github.com/hotstone-seo/hotstone-seo/internal/api/repository"
	"github.com/hotstone-seo/hotstone-seo/internal/api/service"
	"github.com/hotstone-seo/hotstone-seo/internal/app/infra"
	"github.com/labstack/echo"
	"go.uber.org/dig"
)

// TagCntrl is the Controller to manage Tag entity
type TagCntrl struct {
	dig.In
	service.TagService
	*infra.App
}

// Route is a method to define exposed paths on a Controller
func (c *TagCntrl) Route(e *echo.Group) {
	e.GET("/tags", c.Find)
	e.GET("/tags/:id", c.FindOne)
	e.POST("/tags", c.Create)
	e.PUT("/tags/:id", c.Update)
	e.DELETE("/tags/:id", c.Delete)
}

// Find returns a list of Tags based on query params provided
func (c *TagCntrl) Find(ctx echo.Context) (err error) {
	var (
		tags    []*repository.Tag
		reqCtx  = ctx.Request().Context()
		filters = ctx.QueryParams()
	)
	if tags, err = c.TagService.Find(reqCtx, filters); err != nil {
		return httpError(err)
	}
	return ctx.JSON(http.StatusOK, tags)
}

// FindOne returns a single Tag entity based on id provided in the path
func (c *TagCntrl) FindOne(ctx echo.Context) (err error) {
	var tag *repository.Tag
	if tag, err = c.TagService.FindOne(
		ctx.Request().Context(),
		ctx.Param("id"),
	); err != nil {
		return httpError(err)
	}
	return ctx.JSON(http.StatusOK, tag)
}

// Create adds a new Tag entity
func (c *TagCntrl) Create(ctx echo.Context) (err error) {
	var (
		tag repository.Tag
		id  int64
	)
	if err = ctx.Bind(&tag); err != nil {
		return err
	}
	if id, err = c.TagService.Create(ctx.Request().Context(), tag); err != nil {
		return httpError(err)
	}
	ctx.Response().Header().Set(echo.HeaderLocation, fmt.Sprintf("/tags/%d", id))
	return ctx.NoContent(http.StatusCreated)
}

// Update modifies an existing Tag entity
func (c *TagCntrl) Update(ctx echo.Context) (err error) {
	var tag repository.Tag
	if err = ctx.Bind(&tag); err != nil {
		return err
	}
	if err = c.TagService.Update(
		ctx.Request().Context(),
		ctx.Param("id"),
		tag,
	); err != nil {
		return httpError(err)
	}
	return ctx.NoContent(http.StatusOK)
}

// Delete removes an existing Tag entity
func (c *TagCntrl) Delete(ctx echo.Context) (err error) {
	if err = c.TagService.Delete(
		ctx.Request().Context(),
		ctx.Param("id"),
	); err != nil {
		return httpError(err)
	}
	return ctx.NoContent(http.StatusOK)
}
