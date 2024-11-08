package routes

import (
	"net/http"

	"codeberg.org/tomkoid/stravule/backend/internal/resolvers"
	"github.com/labstack/echo/v4"
)

func ListFilters(c echo.Context) error {
	sid := c.QueryParam("sid")
	canteen := c.QueryParam("canteen")

	if sid == "" || canteen == "" {
		return c.String(http.StatusBadRequest, "Missing `sid` or `canteen` query parameter")
	}

	filters := resolvers.GetFilters(sid, canteen)

	return c.JSON(http.StatusOK, filters)
}
