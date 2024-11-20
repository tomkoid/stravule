package middleware

import (
	"context"
	"net/http"

	"codeberg.org/tomkoid/stravule/backend/internal/database"
	"github.com/labstack/echo/v4"
)

func Betatester(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		userHash := c.QueryParam("user_hash")

		if userHash != "" {
			user, err := database.DB.GetUser(context.Background(), userHash)

			if err == nil && !user.IsBetaTester {
				return c.String(http.StatusUnauthorized, "This user is not a betatester.")
			}
		}

		return next(c)
	}
}
