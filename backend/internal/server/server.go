package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func StartServer() {
	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "[${method} ${uri} ${status}]\n",
	}))
	e.Use(middleware.CORS())

	e.HideBanner = true

	apiV1 := e.Group("/api/v1")
	initAPI(apiV1)

	e.Logger.Fatal(e.Start(":1323"))
}
