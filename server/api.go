package server

import (
	"github.com/hotstone-seo/hotstone-seo/pkg/oauth2google"
	"github.com/hotstone-seo/hotstone-seo/server/controller"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"go.uber.org/dig"
)

type api struct {
	dig.In
	oauth2google.AuthCntrl
	controller.RuleCntrl
	controller.DataSourceCntrl
	controller.TagCntrl
	controller.CenterCntrl
	controller.MetricsCntrl
	controller.AuditTrailCntrl
	controller.StructuredDataCntrl
	controller.UserCntrl
	controller.RoleTypeCntrl
	controller.ClientKeyCntrl
}

func (a *api) route(e *echo.Echo) {
	e.POST("auth/google/login", a.AuthCntrl.Login)
	e.GET("auth/google/callback", a.AuthCntrl.Callback)

	group := e.Group("/api")
	group.Use(a.AuthCntrl.Middleware())
	group.Use(a.AuthCntrl.SetTokenCtxMiddleware())
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
	a.RoleTypeCntrl.Route(group)
	a.ClientKeyCntrl.Route(group)
}
