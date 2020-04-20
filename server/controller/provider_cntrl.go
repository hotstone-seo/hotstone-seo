package controller

import (
	"net/http"
	"strconv"

	"github.com/hotstone-seo/hotstone-seo/pkg/cachekit"
	"github.com/hotstone-seo/hotstone-seo/server/service"
	"github.com/labstack/echo"
	"go.uber.org/dig"
)

// ProviderCntrl is controller for provider function
type ProviderCntrl struct {
	dig.In
	service.ProviderService
}

// MatchRule to match rule
func (p *ProviderCntrl) MatchRule(c echo.Context) (err error) {
	var (
		req  service.MatchRequest
		resp *service.MatchResponse
		ctx  = c.Request().Context()
	)
	if err = c.Bind(&req); err != nil {
		return err
	}
	if resp, err = p.ProviderService.Match(ctx, req); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}
	return c.JSON(http.StatusOK, resp)
}

// FetchTag to fetch the tag
func (p *ProviderCntrl) FetchTag(c echo.Context) (err error) {
	var (
		id     int64
		locale string
		tags   []*service.ITag
	)
	ctx := c.Request().Context()

	if rawID := c.Param("id"); rawID != "" {
		if id, err = strconv.ParseInt(rawID, 10, 64); err != nil {
			return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
		}
	} else {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, "Missing url param for `ID`")
	}

	if locale = c.QueryParam("locale"); locale == "" {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, "Missing query param for `Locale`")
	}

	pragma := cachekit.CreatePragma(c.Request())
	tags, err = p.ProviderService.FetchTagsWithCache(ctx, id, locale, pragma)
	cachekit.SetHeader(c.Response(), pragma)
	if cachekit.NotModifiedError(err) {
		return echo.NewHTTPError(http.StatusNotModified)
	}
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, tags)
}

// NOTE: Keep it as example
// RetrieveData retrieve data
// func (p *ProviderCntrl) RetrieveData(c echo.Context) (err error) {
// 	var (
// 		req  service.RetrieveDataRequest
// 		resp *service.RetrieveDataResponse
// 		ctx  = c.Request().Context()
// 	)
// 	if err = c.Bind(&req); err != nil {
// 		return
// 	}

// 	pragma := cachekit.CreatePragma(c.Request())

// 	resp, err = p.ProviderService.RetrieveData(ctx, req, pragma)

// 	header := c.Response().Header()
// 	for key, value := range pragma.ResponseHeaders() {
// 		header.Set(key, value)
// 	}

// 	if cachekit.NotModifiedError(err) {
// 		return echo.NewHTTPError(http.StatusNotModified, err.Error())
// 	}

// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
// 	}

// 	header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

// 	c.Response().WriteHeader(http.StatusOK)
// 	_, err = c.Response().Write(resp.Data)
// 	return
// }

// // Tags to get tag
// func (p *ProviderCntrl) Tags(c echo.Context) (err error) {
// 	var (
// 		req  service.ProvideTagsRequest
// 		tags []*service.InterpolatedTag
// 		ctx  = c.Request().Context()
// 	)
// 	if err = c.Bind(&req); err != nil {
// 		return
// 	}

// 	pragma := cachekit.CreatePragma(c.Request())
// 	if tags, err = p.ProviderService.Tags(ctx, req, pragma); err != nil {
// 		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
// 	}
// 	return c.JSON(http.StatusOK, tags)
// }
