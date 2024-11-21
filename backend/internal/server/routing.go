package server

import (
	"codeberg.org/tomkoid/stravule/backend/internal/routes"
	"codeberg.org/tomkoid/stravule/backend/internal/server/middleware"
	"github.com/labstack/echo/v4"
)

func initAPI(e *echo.Group) error {
	// sid, canteen, user hash middleware
	scuh := e.Group("")
	scuh.Use(middleware.CheckSCUH)

	uh := e.Group("")
	uh.Use(middleware.CheckUserHash)

	e.POST("/login", routes.Login)
	scuh.GET("/orders", routes.Orders)
	scuh.POST("/save_orders", routes.SendOrders)
	uh.GET("/filters", routes.ListFilters)
	uh.POST("/filters", routes.AddFilter)
	uh.DELETE("/filters", routes.RemoveFilter)

	return nil
}
