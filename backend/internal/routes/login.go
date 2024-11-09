package routes

import (
	"net/http"

	"codeberg.org/tomkoid/stravule/backend/internal/api"
	"github.com/labstack/echo/v4"
)

func Login(c echo.Context) error {
	username := c.QueryParam("username")
	password := c.QueryParam("password")
	canteenParam := c.QueryParam("canteen")

	if username == "" || password == "" || canteenParam == "" {
		return c.String(http.StatusBadRequest, "Missing parameters")
	}

	loginData, err := api.Login(username, password, canteenParam)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, loginData)
}
