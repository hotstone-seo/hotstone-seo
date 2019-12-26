package controller

import (
	"fmt"
	"net/http"

	"github.com/hotstone-seo/hotstone-server/app/service"
	"github.com/labstack/echo"
	"go.uber.org/dig"
)

// CenterCntrl is controller to rule entity
type CenterCntrl struct {
	dig.In
	service.CenterService
}

// Route to define API Route
func (c *CenterCntrl) Route(e *echo.Echo) {
	e.POST("center/addMetaTag", c.AddMetaTag)
	e.POST("center/addTitleTag", c.AddTitleTag)
	e.POST("center/addCanonicalTag", c.AddCanonicalTag)
	e.POST("center/addScriptTag", c.AddScriptTag)
	e.POST("center/addArticle", c.AddArticle)
}

// AddMetaTag add meta tag
func (c *CenterCntrl) AddMetaTag(ctx echo.Context) (err error) {
	var (
		req            service.AddMetaTagRequest
		lastInsertedID int64
	)
	if err = ctx.Bind(&req); err != nil {
		return
	}
	if lastInsertedID, err = c.CenterService.AddMetaTag(req); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}
	return ctx.JSON(http.StatusCreated, GeneralResponse{
		Message: fmt.Sprintf("Success insert new meta tag #%d", lastInsertedID),
	})
}

// AddTitleTag add title tag
func (c *CenterCntrl) AddTitleTag(ctx echo.Context) (err error) {
	var (
		req            service.AddTitleTagRequest
		lastInsertedID int64
	)
	if err = ctx.Bind(&req); err != nil {
		return
	}
	if lastInsertedID, err = c.CenterService.AddTitleTag(req); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}
	return ctx.JSON(http.StatusCreated, GeneralResponse{
		Message: fmt.Sprintf("Success insert new title tag #%d", lastInsertedID),
	})
}

// AddCanoncicalTag add canonical tag
func (c *CenterCntrl) AddCanonicalTag(ctx echo.Context) (err error) {
	var (
		req            service.AddCanonicalTagRequest
		lastInsertedID int64
	)
	if err = ctx.Bind(&req); err != nil {
		return
	}
	if lastInsertedID, err = c.CenterService.AddCanonicalTag(req); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}
	return ctx.JSON(http.StatusCreated, GeneralResponse{
		Message: fmt.Sprintf("Success insert new canonical tag #%d", lastInsertedID),
	})
}

// AddScriptTag add script tag
func (c *CenterCntrl) AddScriptTag(ctx echo.Context) (err error) {
	var (
		req            service.AddScriptTagRequest
		lastInsertedID int64
	)
	if err = ctx.Bind(&req); err != nil {
		return
	}
	if lastInsertedID, err = c.CenterService.AddScriptTag(req); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}
	return ctx.JSON(http.StatusCreated, GeneralResponse{
		Message: fmt.Sprintf("Success insert new canonical tag #%d", lastInsertedID),
	})
}

// AddArticle add article
func (*CenterCntrl) AddArticle(ctx echo.Context) error {
	return echo.NewHTTPError(http.StatusNotImplemented, "Not implemented")
}
