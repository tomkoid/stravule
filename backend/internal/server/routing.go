package server

import (
	"codeberg.org/tomkoid/stravule/backend/internal/routes"
	"github.com/labstack/echo/v4"
)

func initAPI(e *echo.Group) error {
	e.POST("/login", routes.Login)
	e.GET("/orders", routes.Orders)
	e.GET("/filters", routes.ListFilters)
	e.POST("/filters", routes.AddFilter)
	e.DELETE("/filters", routes.RemoveFilter)

	return nil
}
