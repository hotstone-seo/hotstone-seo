package server

import (
	"context"
	"database/sql"
	"time"

	logrusmiddleware "github.com/bakatz/echo-logrusmiddleware"
	"github.com/go-redis/redis"
	"github.com/hotstone-seo/hotstone-seo/internal/config"
	"github.com/hotstone-seo/hotstone-seo/internal/profiler"
	"github.com/hotstone-seo/hotstone-seo/internal/provider"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	log "github.com/sirupsen/logrus"

	"go.uber.org/dig"
)

type server struct {
	dig.In
	*config.Config

	API      API
	Provider provider.Controller
	Profiler profiler.Profiler

	Postgres *sql.DB
	Redis    *redis.Client
}

func startServer(s server) error {
	e := echo.New()
	defer shutdown(e)

	e.HideBanner = true
	initLogger(e, s.Debug)
	initErrHandler(e)

	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:  "build",
		HTML5: true,
	}))

	s.API.SetRoute(e)
	s.Provider.SetRoute(e)
	s.Profiler.SetRoute(e)

	return e.Start(s.Address)
}

func initErrHandler(e *echo.Echo) {
	e.HTTPErrorHandler = func(err error, c echo.Context) {
		e.DefaultHTTPErrorHandler(err, c)
		log.Error(err.Error())
	}
}

func initLogger(e *echo.Echo, debug bool) {
	e.Logger = logrusmiddleware.Logger{Logger: log.StandardLogger()}
	e.Debug = debug
	if debug {
		log.SetLevel(log.DebugLevel)
		log.SetFormatter(&log.TextFormatter{})
		e.Use(logrusmiddleware.HookWithConfig(logrusmiddleware.Config{
			IncludeRequestBodies:  true,
			IncludeResponseBodies: true,
		}))
	} else {
		log.SetLevel(log.WarnLevel)
		log.SetFormatter(&log.JSONFormatter{})
		e.Use(logrusmiddleware.HookWithConfig(logrusmiddleware.Config{}))
	}
}

func shutdown(e *echo.Echo) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	e.Shutdown(ctx)
}
