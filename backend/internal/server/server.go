package server

import (
	"os"

	sMiddleware "codeberg.org/tomkoid/stravule/backend/internal/server/middleware"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func StartServer() {
	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "[${method} ${uri} ${status}]\n",
	}))
	e.Use(middleware.BodyLimit("6M"))
	e.Use(middleware.CORS())

	if os.Getenv("BETATESTERS_ONLY") == "true" {
		e.Use(sMiddleware.Betatester)
	}

	e.HideBanner = true

	apiV1 := e.Group("/api/v1")
	initAPI(apiV1)

	e.Logger.Fatal(e.Start(":1323"))
}
