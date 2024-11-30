package server

import (
	"codeberg.org/tomkoid/stravule/backend/internal/config"
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
	scuh.GET("/account_info", routes.AccountInfo)
	scuh.POST("/save_orders", routes.SendOrders)
	uh.GET("/filters", routes.ListFilters)
	uh.POST("/filters", routes.AddFilter)
	uh.PUT("/filters_weight", routes.UpdateFilterWeight)
	uh.DELETE("/filters", routes.RemoveFilter)
	uh.GET("/order_day_exceptions", routes.GetOrderDayExceptions)
	uh.POST("/order_day_exception", routes.AddOrderDayException)
	uh.DELETE("/order_day_exception", routes.RemoveOrderDayException)

	// betatester stuff
	if config.Cfg.BetatestersOnly {
		uh.POST("/betatester", routes.RegisterBetatester)
	}

	return nil
}
