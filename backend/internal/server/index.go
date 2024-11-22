package server

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

type indexResponse struct {
	BetaTestersOnly bool `json:"betatestersOnly"`
}

func indexHandler(c echo.Context) error {
	res := indexResponse{
		BetaTestersOnly: os.Getenv("BETATESTERS_ONLY") == "true",
	}

	return c.JSON(http.StatusOK, res)
}
