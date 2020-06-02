package provider

import (
	"database/sql"
	"net/http"

	"github.com/hotstone-seo/hotstone-seo/internal/api/service"
	"github.com/hotstone-seo/hotstone-seo/pkg/cachekit"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/typical-go/typical-rest-server/pkg/errvalid"
	"go.uber.org/dig"
)

// Controller for provider function
type Controller struct {
	dig.In
	Service
	service.ClientKeyService
}

// SetRoute for provider
func (p *Controller) SetRoute(e *echo.Echo) {
	g := e.Group("/p")
	g.Use(p.AuthMiddleware())
	g.Use(middleware.CORSWithConfig(middleware.DefaultCORSConfig))
	g.Use(middleware.Recover())

	g.GET("/match", p.MatchRule)
	g.GET("/fetch-tags", p.FetchTag)
}

// MatchRule to match rule
func (p *Controller) MatchRule(c echo.Context) (err error) {
	ctx := c.Request().Context()
	resp, err := p.Service.Match(ctx, c.QueryParams())

	if err != nil {
		return httpError(err)
	}

	return c.JSON(http.StatusOK, resp)
}

// FetchTag to fetch the tag
func (p *Controller) FetchTag(c echo.Context) (err error) {

	pragma := cachekit.CreatePragma(c.Request())
	tags, err := p.Service.FetchTagsWithCache(
		c.Request().Context(),
		c.QueryParams(),
		pragma,
	)
	cachekit.SetHeader(c.Response(), pragma)

	if err != nil {
		return httpError(err)
	}

	return c.JSON(http.StatusOK, tags)
}

// AuthMiddleware do key-based auth
func (p *Controller) AuthMiddleware() echo.MiddlewareFunc {
	return middleware.KeyAuthWithConfig(middleware.KeyAuthConfig{
		Validator: func(clientKey string, e echo.Context) (bool, error) {
			if p.ClientKeyService.IsValidClientKey(e.Request().Context(), clientKey) {
				return true, nil
			}
			return false, nil
		},
	})
}

func httpError(err error) *echo.HTTPError {
	if err == sql.ErrNoRows {
		return echo.NewHTTPError(http.StatusNotFound)
	}

	if errvalid.Check(err) {
		return echo.NewHTTPError(
			http.StatusUnprocessableEntity,
			errvalid.Message(err),
		)
	}

	if cachekit.NotModifiedError(err) {
		return echo.NewHTTPError(http.StatusNotModified)
	}

	return echo.NewHTTPError(
		http.StatusInternalServerError,
		err.Error(),
	)
}
