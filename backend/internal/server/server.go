package server

import (
	"codeberg.org/tomkoid/stravule/backend/internal/config"
	sMiddleware "codeberg.org/tomkoid/stravule/backend/internal/server/middleware"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func StartServer() {
	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "[${remote_ip}] ${method} ${status} ${path} ${latency_human} IN:${bytes_in}B OUT:${bytes_out}B\n",
	}))
	e.Use(middleware.BodyLimit("6M"))
	e.Use(middleware.CORS())

	if config.Cfg.BetatestersOnly {
		e.Use(sMiddleware.Betatester)
	}

	e.HideBanner = true

	apiV1 := e.Group("/api/v1")
	initAPI(apiV1)

	e.GET("/", indexHandler)

	host := ":1323"
	if config.Cfg.Host != "" {
		host = config.Cfg.Host
	}

	e.Logger.Fatal(e.Start(host))
}
