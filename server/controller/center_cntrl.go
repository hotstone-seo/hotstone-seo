package controller

import (
	"net/http"

	"github.com/hotstone-seo/hotstone-seo/server/repository"
	"github.com/hotstone-seo/hotstone-seo/server/service"
	"github.com/labstack/echo"
	"go.uber.org/dig"
)

// CenterCntrl is controller to rule entity
type CenterCntrl struct {
	dig.In
	service.CenterService
}

// Route to define API Route
func (c *CenterCntrl) Route(e *echo.Group) {
	e.POST("/center/addMetaTag", c.AddMetaTag)
	e.POST("/center/addTitleTag", c.AddTitleTag)
	e.POST("/center/addCanonicalTag", c.AddCanonicalTag)
	e.POST("/center/addScriptTag", c.AddScriptTag)
	e.POST("/center/addArticle", c.AddArticle)
}

// AddMetaTag add meta tag
func (c *CenterCntrl) AddMetaTag(ce echo.Context) (err error) {
	var (
		req service.MetaTagRequest
		tag *repository.Tag
		ctx = ce.Request().Context()
	)
	if err = ce.Bind(&req); err != nil {
		return
	}
	if tag, err = c.CenterService.AddMetaTag(ctx, req); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}
	return ce.JSON(http.StatusCreated, tag)
}

// AddTitleTag add title tag
func (c *CenterCntrl) AddTitleTag(ce echo.Context) (err error) {
	var (
		req service.TitleTagRequest
		tag *repository.Tag
		ctx = ce.Request().Context()
	)
	if err = ce.Bind(&req); err != nil {
		return
	}
	if tag, err = c.CenterService.AddTitleTag(ctx, req); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}
	return ce.JSON(http.StatusCreated, tag)
}

// AddCanoncicalTag add canonical tag
func (c *CenterCntrl) AddCanonicalTag(ce echo.Context) (err error) {
	var (
		req service.CanonicalTagRequest
		tag *repository.Tag
		ctx = ce.Request().Context()
	)
	if err = ce.Bind(&req); err != nil {
		return
	}
	if tag, err = c.CenterService.AddCanonicalTag(ctx, req); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}
	return ce.JSON(http.StatusCreated, tag)
}

// AddScriptTag add script tag
func (c *CenterCntrl) AddScriptTag(ce echo.Context) (err error) {
	var (
		req service.ScriptTagRequest
		tag *repository.Tag
		ctx = ce.Request().Context()
	)
	if err = ce.Bind(&req); err != nil {
		return
	}
	if tag, err = c.CenterService.AddScriptTag(ctx, req); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}
	return ce.JSON(http.StatusCreated, tag)
}

// AddArticle add article
func (*CenterCntrl) AddArticle(ctx echo.Context) error {
	return echo.NewHTTPError(http.StatusNotImplemented, "Not implemented")
}
