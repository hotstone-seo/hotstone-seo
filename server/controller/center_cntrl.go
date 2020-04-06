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
	e.POST("/center/metaTag", c.AddMetaTag)
	e.PUT("/center/metaTag", c.UpdateMetaTag)
	e.POST("/center/titleTag", c.AddTitleTag)
	e.PUT("/center/titleTag", c.UpdateTitleTag)
	e.POST("/center/canonicalTag", c.AddCanonicalTag)
	e.PUT("/center/canonicalTag", c.UpdateCanonicalTag)
	e.POST("/center/scriptTag", c.AddScriptTag)
	e.PUT("/center/scriptTag", c.UpdateScriptTag)
	e.POST("/center/addArticle", c.AddArticle)
}

// AddMetaTag provides endpoint to add meta tag
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

// UpdateMetaTag provides endpoint to update meta tag
func (c *CenterCntrl) UpdateMetaTag(ce echo.Context) (err error) {
	var (
		req service.MetaTagRequest
		ctx = ce.Request().Context()
	)
	if err = ce.Bind(&req); err != nil {
		return
	}
	if req.ID <= 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}
	if err = c.CenterService.UpdateMetaTag(ctx, req); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}
	return ce.JSON(http.StatusOK, GeneralResponse{
		Message: fmt.Sprintf("Successfully update meta tag #%d", req.ID),
	})
}

// AddTitleTag provides endpoint to add title tag
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

// UpdateTitleTag provides endpoint to update title tag
func (c *CenterCntrl) UpdateTitleTag(ce echo.Context) (err error) {
	var (
		req service.TitleTagRequest
		ctx = ce.Request().Context()
	)
	if err = ce.Bind(&req); err != nil {
		return
	}
	if req.ID <= 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}
	if err = c.CenterService.UpdateTitleTag(ctx, req); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}
	return ce.JSON(http.StatusOK, GeneralResponse{
		Message: fmt.Sprintf("Successfully update title tag #%d", req.ID),
	})
}

// AddCanonicalTag provides endpoint to add canonical tag
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

// UpdateCanonicalTag provides endpoint to update canonical tag
func (c *CenterCntrl) UpdateCanonicalTag(ce echo.Context) (err error) {
	var (
		req service.CanonicalTagRequest
		ctx = ce.Request().Context()
	)
	if err = ce.Bind(&req); err != nil {
		return
	}
	if req.ID <= 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}
	if err = c.CenterService.UpdateCanonicalTag(ctx, req); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}
	return ce.JSON(http.StatusOK, GeneralResponse{
		Message: fmt.Sprintf("Successfully update canonical tag #%d", req.ID),
	})
}

// AddScriptTag provides endpoint to add script tag
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

// UpdateScriptTag provides endpoint to update script tag
func (c *CenterCntrl) UpdateScriptTag(ce echo.Context) (err error) {
	var (
		req service.ScriptTagRequest
		ctx = ce.Request().Context()
	)
	if err = ce.Bind(&req); err != nil {
		return
	}
	if req.ID <= 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}
	if err = c.CenterService.UpdateScriptTag(ctx, req); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}
	return ce.JSON(http.StatusOK, GeneralResponse{
		Message: fmt.Sprintf("Successfully update script tag #%d", req.ID),
	})
}

// AddArticle add article
func (*CenterCntrl) AddArticle(ctx echo.Context) error {
	return echo.NewHTTPError(http.StatusNotImplemented, "Not implemented")
}
