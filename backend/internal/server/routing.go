package server

import (
	"codeberg.org/tomkoid/stravule/backend/internal/routes"
	"github.com/labstack/echo/v4"
)

func initAPI(e *echo.Group) error {
	e.POST("/login", routes.Login)

	return nil
}
