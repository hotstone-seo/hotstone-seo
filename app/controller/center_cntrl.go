package controller

import (
	"net/http"

	"github.com/labstack/echo"
	"go.uber.org/dig"
)

// CenterCntrl is controller to rule entity
type CenterCntrl struct {
	dig.In
}

// Route to define API Route
func (c *CenterCntrl) Route(e *echo.Echo) {
	e.POST("center/addMetaTag", c.AddMetaTag)
	e.POST("center/addTitleTag", c.AddTitleTag)
	e.POST("center/addCanonicalTag", c.AddCanoncicalTag)
	e.POST("center/addScriptTag", c.AddScriptTag)
	e.POST("center/addArticle", c.AddArticle)
}

// AddMetaTag add meta tag
func (*CenterCntrl) AddMetaTag(ctx echo.Context) error {
	return echo.NewHTTPError(http.StatusNotImplemented, "Not implemented")
}

// AddTitleTag add title tag
func (*CenterCntrl) AddTitleTag(ctx echo.Context) error {
	return echo.NewHTTPError(http.StatusNotImplemented, "Not implemented")
}

// AddCanoncicalTag add canonical tag
func (*CenterCntrl) AddCanoncicalTag(ctx echo.Context) error {
	return echo.NewHTTPError(http.StatusNotImplemented, "Not implemented")
}

// AddScriptTag add script tag
func (*CenterCntrl) AddScriptTag(ctx echo.Context) error {
	return echo.NewHTTPError(http.StatusNotImplemented, "Not implemented")
}

// AddArticle add article
func (*CenterCntrl) AddArticle(ctx echo.Context) error {
	return echo.NewHTTPError(http.StatusNotImplemented, "Not implemented")
}
