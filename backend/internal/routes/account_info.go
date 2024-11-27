package routes

import (
	"net/http"

	"codeberg.org/tomkoid/stravule/backend/internal/api"
	"github.com/labstack/echo/v4"
)

func AccountInfo(c echo.Context) error {
	sid := c.QueryParam("sid")
	canteen := c.QueryParam("canteen")
	// userHash := c.QueryParam("user_hash")

	accountInfo, err := api.GetAccountInfo(sid, canteen)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, accountInfo)
}
