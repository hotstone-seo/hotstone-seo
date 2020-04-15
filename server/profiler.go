package server

import (
	"net/http"

	"github.com/go-redis/redis"
	"github.com/hotstone-seo/hotstone-seo/server/service"
	"github.com/labstack/echo"
	"github.com/typical-go/typical-rest-server/pkg/serverkit"
	"github.com/typical-go/typical-rest-server/pkg/typpostgres"
	"go.uber.org/dig"
)

type profiler struct {
	dig.In
	Pg    *typpostgres.DB
	Redis *redis.Client
	service.URLService
}

func (p *profiler) route(e *echo.Echo) {
	e.Any("application/health", p.healthCheck)
	e.GET("prof/url-tree", p.dumpURLTree)
}

func (p *profiler) healthCheck(ec echo.Context) (err error) {
	healthcheck := serverkit.NewHealthCheck()
	healthcheck.Put("postgres", p.Pg.Ping)
	healthcheck.Put("redis", p.Redis.Ping().Err)

	status, message := healthcheck.Process()
	return ec.JSON(status, message)
}

func (p *profiler) dumpURLTree(c echo.Context) (err error) {
	return c.String(http.StatusOK, p.URLService.DumpTree())
}