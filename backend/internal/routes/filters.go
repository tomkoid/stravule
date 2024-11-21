package routes

import (
	"net/http"

	"codeberg.org/tomkoid/stravule/backend/internal/resolvers"
	"github.com/labstack/echo/v4"
)

func ListFilters(c echo.Context) error {
	userHash := c.QueryParam("user_hash")

	filters, err := resolvers.GetFilters(&userHash)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, filters)
}

func AddFilter(c echo.Context) error {
	userHash := c.QueryParam("user_hash")

	filter := resolvers.Filter{}
	if err := c.Bind(&filter); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	filters, err := resolvers.AddFilter(&userHash, filter)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, filters)
}

func RemoveFilter(c echo.Context) error {
	userHash := c.QueryParam("user_hash")

	filter := resolvers.Filter{}
	if err := c.Bind(&filter); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	filters, err := resolvers.RemoveFilter(&userHash, filter)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, filters)
}
