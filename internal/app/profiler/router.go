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
	"github.com/typical-go/typical-rest-server/pkg/echokit"
	"go.uber.org/dig"
)

var _ echokit.Router = (*Router)(nil)

type (
	// Router for pofiler
	Router struct {
		dig.In
		MainDB   *sql.DB
		AnalytDB *sql.DB `name:"analyt"`
		Redis    *redis.Client
		urlstore.Store
	}
)

// Route to profiler
func (p *Router) Route(e echokit.Server) error {
	e.Any("application/health", p.healthCheck)
	e.GET("_prof/url-tree", p.dumpURLTree)
	e.GET("_prof/goroutine/count", p.countGoroutine)
	e.GET("_prof/goroutine/trace", p.traceGoroutine)
	return nil
}

func (p *Router) healthCheck(ec echo.Context) (err error) {
	hc := echokit.HealthCheck{
		"main-db":   p.MainDB.Ping,
		"analyt-db": p.AnalytDB.Ping,
		"redis":     p.Redis.Ping().Err,
	}

	return hc.JSON(ec)
}

func (p *Router) dumpURLTree(c echo.Context) (err error) {
	return c.String(http.StatusOK, p.Store.String())
}

func (p *Router) countGoroutine(c echo.Context) (err error) {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"count": runtime.NumGoroutine(),
	})
}

func (p *Router) traceGoroutine(c echo.Context) (err error) {
	debug, _ := strconv.Atoi(c.QueryParam("debug"))

	pprof.Lookup("goroutine").WriteTo(c.Response(), debug)
	return
}
