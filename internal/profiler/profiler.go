package profiler

import (
	"database/sql"
	"net/http"
	"runtime"
	"runtime/pprof"
	"strconv"

	"github.com/go-redis/redis"
	"github.com/hotstone-seo/hotstone-seo/internal/urlstore"
	"github.com/labstack/echo"
	"github.com/typical-go/typical-rest-server/pkg/serverkit"
	"go.uber.org/dig"
)

// Profiler app
type Profiler struct {
	dig.In
	Pg    *sql.DB
	Redis *redis.Client
	urlstore.Store
}

// SetRoute for profiler
func (p *Profiler) SetRoute(e *echo.Echo) {
	e.Any("application/health", p.healthCheck)
	e.GET("_prof/url-tree", p.dumpURLTree)
	e.GET("_prof/goroutine/count", p.countGoroutine)
	e.GET("_prof/goroutine/trace", p.traceGoroutine)
}

func (p *Profiler) healthCheck(ec echo.Context) (err error) {
	healthcheck := serverkit.NewHealthCheck()
	healthcheck.Put("postgres", p.Pg.Ping)
	healthcheck.Put("redis", p.Redis.Ping().Err)

	status, message := healthcheck.Process()
	return ec.JSON(status, message)
}

func (p *Profiler) dumpURLTree(c echo.Context) (err error) {
	return c.String(http.StatusOK, p.Store.String())
}

func (p *Profiler) countGoroutine(c echo.Context) (err error) {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"count": runtime.NumGoroutine(),
	})
}

func (p *Profiler) traceGoroutine(c echo.Context) (err error) {
	debug, _ := strconv.Atoi(c.QueryParam("debug"))

	pprof.Lookup("goroutine").WriteTo(c.Response(), debug)
	return
}
