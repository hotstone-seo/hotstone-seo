package controller

import (
	"fmt"
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
	e.POST("/center/addTag/:type", c.AddTag)
	e.POST("/center/addMetaTag", c.AddMetaTag)
	e.POST("/center/addTitleTag", c.AddTitleTag)
	e.POST("/center/addCanonicalTag", c.AddCanonicalTag)
	e.POST("/center/addScriptTag", c.AddScriptTag)
	e.POST("/center/addArticle", c.AddArticle)
}

func (c *CenterCntrl) AddTag(e echo.Context) (err error) {
	var (
		req     interface{}
		ctx     = e.Request().Context()
		tagType = e.Param("type")
		tag     *repository.Tag
	)
	if req = service.NewTagRequest(tagType); req == nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("%s is not a valid type", tagType))
	}
	if err = e.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if tag, err = c.CenterService.AddTag(ctx, req); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}
	return e.JSON(http.StatusCreated, tag)
}

// AddMetaTag add meta tag
func (c *CenterCntrl) AddMetaTag(ce echo.Context) (err error) {
	var (
		req            service.AddMetaTagRequest
		lastInsertedID int64
		ctx            = ce.Request().Context()
	)
	if err = ce.Bind(&req); err != nil {
		return
	}
	if lastInsertedID, err = c.CenterService.AddMetaTag(ctx, req); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}
	return ce.JSON(http.StatusCreated, GeneralResponse{
		Message: fmt.Sprintf("Success insert new meta tag #%d", lastInsertedID),
	})
}

// AddTitleTag add title tag
func (c *CenterCntrl) AddTitleTag(ce echo.Context) (err error) {
	var (
		req            service.AddTitleTagRequest
		lastInsertedID int64
		ctx            = ce.Request().Context()
	)
	if err = ce.Bind(&req); err != nil {
		return
	}
	if lastInsertedID, err = c.CenterService.AddTitleTag(ctx, req); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}
	return ce.JSON(http.StatusCreated, GeneralResponse{
		Message: fmt.Sprintf("Success insert new title tag #%d", lastInsertedID),
	})
}

// AddCanoncicalTag add canonical tag
func (c *CenterCntrl) AddCanonicalTag(ce echo.Context) (err error) {
	var (
		req            service.AddCanonicalTagRequest
		lastInsertedID int64
		ctx            = ce.Request().Context()
	)
	if err = ce.Bind(&req); err != nil {
		return
	}
	if lastInsertedID, err = c.CenterService.AddCanonicalTag(ctx, req); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}
	return ce.JSON(http.StatusCreated, GeneralResponse{
		Message: fmt.Sprintf("Success insert new canonical tag #%d", lastInsertedID),
	})
}

// AddScriptTag add script tag
func (c *CenterCntrl) AddScriptTag(ce echo.Context) (err error) {
	var (
		req            service.AddScriptTagRequest
		lastInsertedID int64
		ctx            = ce.Request().Context()
	)
	if err = ce.Bind(&req); err != nil {
		return
	}
	if lastInsertedID, err = c.CenterService.AddScriptTag(ctx, req); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}
	return ce.JSON(http.StatusCreated, GeneralResponse{
		Message: fmt.Sprintf("Success insert new canonical tag #%d", lastInsertedID),
	})
}

// AddArticle add article
func (*CenterCntrl) AddArticle(ctx echo.Context) error {
	return echo.NewHTTPError(http.StatusNotImplemented, "Not implemented")
}
