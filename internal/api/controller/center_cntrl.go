package controller

import (
	"fmt"
	"net/http"

	"github.com/hotstone-seo/hotstone-seo/internal/api/repository"
	"github.com/hotstone-seo/hotstone-seo/internal/api/service"
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
	e.POST("/center/meta-tag", c.CreateMetaTag)
	e.PUT("/center/meta-tag/:id", c.UpdateMetaTag)

	e.POST("/center/title-tag", c.AddTitleTag)
	e.PUT("/center/title-tag/:id", c.UpdateTitleTag)

	e.POST("/center/canonical-tag", c.AddCanonicalTag)
	e.PUT("/center/canonical-tag/:id", c.UpdateCanonicalTag)

	e.POST("/center/script-tag", c.AddScriptTag)
	e.PUT("/center/script-tag/:id", c.UpdateScriptTag)

	e.POST("/center/faq-page", c.AddFAQPage)
	e.PUT("/center/faq-page", c.UpdateFAQPage)

	e.POST("/center/breadcrumb-list", c.AddBreadcrumbList)
	e.PUT("/center/breadcrumb-list", c.UpdateBreadcrumbList)

	e.POST("/center/local-business", c.AddLocalBusiness)
	e.PUT("/center/local-business", c.UpdateLocalBusiness)

	e.POST("/center/addArticle", c.AddArticle)
}

// CreateMetaTag adds a new meta tag entity
func (c *CenterCntrl) CreateMetaTag(ctx echo.Context) (err error) {
	var (
		req service.MetaTagRequest
		tag *repository.Tag
	)
	if err = ctx.Bind(&req); err != nil {
		return
	}
	if tag, err = c.CenterService.AddMetaTag(ctx.Request().Context(), req); err != nil {
		return httpError(err)
	}
	return ctx.JSON(http.StatusCreated, tag)
}

// UpdateMetaTag modifies an existing meta tag entity
func (c *CenterCntrl) UpdateMetaTag(ctx echo.Context) (err error) {
	var req service.MetaTagRequest
	if err = ctx.Bind(&req); err != nil {
		return
	}
	if err = c.CenterService.UpdateMetaTag(
		ctx.Request().Context(),
		ctx.Param("id"),
		req,
	); err != nil {
		return httpError(err)
	}
	return ctx.NoContent(http.StatusOK)
}

// AddTitleTag adds a new title tag entity
func (c *CenterCntrl) AddTitleTag(ctx echo.Context) (err error) {
	var (
		req service.TitleTagRequest
		tag *repository.Tag
	)
	if err = ctx.Bind(&req); err != nil {
		return
	}
	if tag, err = c.CenterService.AddTitleTag(ctx.Request().Context(), req); err != nil {
		return httpError(err)
	}
	return ctx.JSON(http.StatusCreated, tag)
}

// UpdateTitleTag modifies an existing title tag entity
func (c *CenterCntrl) UpdateTitleTag(ctx echo.Context) (err error) {
	var req service.TitleTagRequest
	if err = ctx.Bind(&req); err != nil {
		return
	}
	if err = c.CenterService.UpdateTitleTag(
		ctx.Request().Context(),
		ctx.Param("id"),
		req,
	); err != nil {
		return httpError(err)
	}
	return ctx.NoContent(http.StatusOK)
}

// AddCanonicalTag adds a new canonical tag entity
func (c *CenterCntrl) AddCanonicalTag(ctx echo.Context) (err error) {
	var (
		req service.CanonicalTagRequest
		tag *repository.Tag
	)
	if err = ctx.Bind(&req); err != nil {
		return
	}
	if tag, err = c.CenterService.AddCanonicalTag(ctx.Request().Context(), req); err != nil {
		return httpError(err)
	}
	return ctx.JSON(http.StatusCreated, tag)
}

// UpdateCanonicalTag modifies an existing canonical tag entity
func (c *CenterCntrl) UpdateCanonicalTag(ctx echo.Context) (err error) {
	var req service.CanonicalTagRequest
	if err = ctx.Bind(&req); err != nil {
		return
	}
	if err = c.CenterService.UpdateCanonicalTag(
		ctx.Request().Context(),
		ctx.Param("id"),
		req,
	); err != nil {
		return httpError(err)
	}
	return ctx.NoContent(http.StatusOK)
}

// AddScriptTag adds a new script tag entity
func (c *CenterCntrl) AddScriptTag(ctx echo.Context) (err error) {
	var (
		req service.ScriptTagRequest
		tag *repository.Tag
	)
	if err = ctx.Bind(&req); err != nil {
		return
	}
	if tag, err = c.CenterService.AddScriptTag(ctx.Request().Context(), req); err != nil {
		return httpError(err)
	}
	return ctx.JSON(http.StatusCreated, tag)
}

// UpdateScriptTag modifies an existing script tag entity
func (c *CenterCntrl) UpdateScriptTag(ctx echo.Context) (err error) {
	var req service.ScriptTagRequest
	if err = ctx.Bind(&req); err != nil {
		return
	}
	if err = c.CenterService.UpdateScriptTag(
		ctx.Request().Context(),
		ctx.Param("id"),
		req,
	); err != nil {
		return httpError(err)
	}
	return ctx.NoContent(http.StatusOK)
}

// AddFAQPage provides endpoint to add FAQPage structured data
func (c *CenterCntrl) AddFAQPage(ce echo.Context) (err error) {
	var (
		req        service.FAQPageRequest
		structData *repository.StructuredData
		ctx        = ce.Request().Context()
	)
	if err = ce.Bind(&req); err != nil {
		return
	}
	if structData, err = c.CenterService.AddFAQPage(ctx, req); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}
	return ce.JSON(http.StatusCreated, structData)
}

// UpdateFAQPage provides endpoint to update FAQPage structured data
func (c *CenterCntrl) UpdateFAQPage(ce echo.Context) (err error) {
	var (
		req service.FAQPageRequest
		ctx = ce.Request().Context()
	)
	if err = ce.Bind(&req); err != nil {
		return
	}
	if req.ID <= 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}
	if err = c.CenterService.UpdateFAQPage(ctx, req); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}
	return ce.JSON(http.StatusOK, GeneralResponse{
		Message: fmt.Sprintf("Successfully update FAQPage structured data #%d", req.ID),
	})
}

// AddBreadcrumbList provides endpoint to add BreadcrumbList structured data
func (c *CenterCntrl) AddBreadcrumbList(ce echo.Context) (err error) {
	var (
		req        service.BreadcrumbListRequest
		structData *repository.StructuredData
		ctx        = ce.Request().Context()
	)
	if err = ce.Bind(&req); err != nil {
		return
	}
	if structData, err = c.CenterService.AddBreadcrumbList(ctx, req); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}
	return ce.JSON(http.StatusCreated, structData)
}

// UpdateBreadcrumbList provides endpoint to update BreadcrumbList structured data
func (c *CenterCntrl) UpdateBreadcrumbList(ce echo.Context) (err error) {
	var (
		req service.BreadcrumbListRequest
		ctx = ce.Request().Context()
	)
	if err = ce.Bind(&req); err != nil {
		return
	}
	if req.ID <= 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}
	if err = c.CenterService.UpdateBreadcrumbList(ctx, req); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}
	return ce.JSON(http.StatusOK, GeneralResponse{
		Message: fmt.Sprintf("Successfully update BreadcrumbList structured data #%d", req.ID),
	})
}

// AddLocalBusiness provides endpoint to add LocalBusiness structured data
func (c *CenterCntrl) AddLocalBusiness(ce echo.Context) (err error) {
	var (
		req        service.LocalBusinessRequest
		structData *repository.StructuredData
		ctx        = ce.Request().Context()
	)
	if err = ce.Bind(&req); err != nil {
		return
	}
	if structData, err = c.CenterService.AddLocalBusiness(ctx, req); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}
	return ce.JSON(http.StatusCreated, structData)
}

// UpdateLocalBusiness provides endpoint to update LocalBusiness structured data
func (c *CenterCntrl) UpdateLocalBusiness(ce echo.Context) (err error) {
	var (
		req service.LocalBusinessRequest
		ctx = ce.Request().Context()
	)
	if err = ce.Bind(&req); err != nil {
		return
	}
	if req.ID <= 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}
	if err = c.CenterService.UpdateLocalBusiness(ctx, req); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}
	return ce.JSON(http.StatusOK, GeneralResponse{
		Message: fmt.Sprintf("Successfully update LocalBusiness structured data #%d", req.ID),
	})
}

// AddArticle add article
func (*CenterCntrl) AddArticle(ctx echo.Context) error {
	return echo.NewHTTPError(http.StatusNotImplemented, "Not implemented")
}
