package routes

import (
	"net/http"

	"codeberg.org/tomkoid/stravule/backend/internal/api"
	"github.com/labstack/echo/v4"
)

func Orders(c echo.Context) error {
	sid := c.QueryParam("sid")
	canteen := c.QueryParam("canteen")

	if sid == "" || canteen == "" {
		return c.String(http.StatusBadRequest, "Missing SID or canteen query parameter")
	}

	res, err := api.Orders(sid, canteen)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, res)
}
