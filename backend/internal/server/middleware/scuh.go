package middleware

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func CheckSCUH(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sid := c.QueryParam("sid")
		canteen := c.QueryParam("canteen")
		userHash := c.QueryParam("user_hash")

		if strings.TrimSpace(sid) == "" {
			return c.String(http.StatusBadRequest, "Missing `sid` query parameter")
		}

		if strings.TrimSpace(canteen) == "" {
			return c.String(http.StatusBadRequest, "Missing `canteen` query parameter")
		}

		if strings.TrimSpace(userHash) == "" {
			return c.String(http.StatusBadRequest, "Missing `user_hash` query parameter")
		}

		return next(c)
	}
}

func CheckUserHash(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		userHash := c.QueryParam("user_hash")

		if strings.TrimSpace(userHash) == "" {
			return c.String(http.StatusBadRequest, "Missing `user_hash` query parameter")
		}

		return next(c)
	}
}
