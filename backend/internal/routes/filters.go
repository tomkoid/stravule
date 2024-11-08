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

	filters := resolvers.GetFilters(&sid, &canteen)

	return c.JSON(http.StatusOK, filters)
}

func AddFilter(c echo.Context) error {
	sid := c.QueryParam("sid")
	canteen := c.QueryParam("canteen")

	if sid == "" || canteen == "" {
		return c.String(http.StatusBadRequest, "Missing `sid` or `canteen` query parameter")
	}

	filter := resolvers.Filter{}
	if err := c.Bind(&filter); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	filters, err := resolvers.AddFilter(sid, canteen, filter)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, filters)
}

func RemoveFilter(c echo.Context) error {
	sid := c.QueryParam("sid")
	canteen := c.QueryParam("canteen")

	if sid == "" || canteen == "" {
		return c.String(http.StatusBadRequest, "Missing `sid` or `canteen` query parameter")
	}

	filter := resolvers.Filter{}
	if err := c.Bind(&filter); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	filters, err := resolvers.RemoveFilter(sid, canteen, filter)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, filters)
}
