package routes

import (
	"net/http"

	"codeberg.org/tomkoid/stravule/backend/internal/api"
	"github.com/labstack/echo/v4"
)

func RegisterBetatester(c echo.Context) error {
	userHash := c.QueryParam("user_hash")
	registerHash := c.QueryParam("register_hash")

	if registerHash == "" {
		return c.String(http.StatusBadRequest, "Missing `register_hash` parameter")
	}

	err := api.RegisterBetatester(userHash, registerHash)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, "User successfully registered as betatester")
}
