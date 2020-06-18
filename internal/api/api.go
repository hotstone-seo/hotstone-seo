package api

import (
	"github.com/hotstone-seo/hotstone-seo/internal/api/controller"
	"github.com/hotstone-seo/hotstone-seo/pkg/oauth2google"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"go.uber.org/dig"
)

// API side
type API struct {
	dig.In
	controller.AuthMiddleware

	Oauth2GoogleCntrl oauth2google.AuthCntrl
	controller.AuthCntrl
	controller.RuleCntrl
	controller.DataSourceCntrl
	controller.TagCntrl
	controller.CenterCntrl
	controller.MetricsCntrl
	controller.AuditTrailCntrl
	controller.StructuredDataCntrl
	controller.UserCntrl
	controller.UserRoleCntrl
	controller.ClientKeyCntrl
	controller.SettingCntrl
}

// SetRoute for API
func (a *API) SetRoute(e *echo.Echo) {
	e.POST("auth/google/login", a.AuthCntrl.Login)
	e.GET("auth/google/callback", a.Oauth2GoogleCntrl.Callback(a.AuthCntrl.Oauth2GoogleCallback))

	group := e.Group("/api")
	group.Use(a.AuthMiddleware.Middleware())
	group.Use(a.AuthMiddleware.SetTokenCtxMiddleware())
	group.Use(a.AuthMiddleware.CheckAuthModules())
	group.Use(middleware.CORSWithConfig(middleware.DefaultCORSConfig))
	group.Use(middleware.Recover())
	group.POST("/logout", a.AuthCntrl.Logout)

	a.RuleCntrl.Route(group)
	a.DataSourceCntrl.Route(group)
	a.TagCntrl.Route(group)
	a.CenterCntrl.Route(group)
	a.MetricsCntrl.Route(group)
	a.AuditTrailCntrl.Route(group)
	a.StructuredDataCntrl.Route(group)
	a.UserCntrl.Route(group)
	a.UserRoleCntrl.Route(group)
	a.ClientKeyCntrl.Route(group)
	a.SettingCntrl.Route(group)
}
