package server

import (
	"github.com/labstack/echo/v4"
)

func StartServer() {
	e := echo.New()
	e.HideBanner = true

	apiV1 := e.Group("/api/v1")
	initAPI(apiV1)

	e.Logger.Fatal(e.Start(":1323"))
}
