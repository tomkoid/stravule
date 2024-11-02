package routes

import (
	"net/http"
	"strconv"

	"codeberg.org/tomkoid/stravule/backend/internal/api"
	"github.com/labstack/echo/v4"
)

func Orders(c echo.Context) error {
	var pick bool = false

	sid := c.QueryParam("sid")
	canteen := c.QueryParam("canteen")
	pickQuery := c.QueryParam("pick")

	if sid == "" || canteen == "" {
		return c.String(http.StatusBadRequest, "Missing SID or canteen query parameter")
	}

	if pickQuery != "" {
		pickTmp, err := strconv.ParseBool(pickQuery)
		if err != nil {
			return c.String(http.StatusBadRequest, "Invalid pick query parameter")
		}

		pick = pickTmp
	}

	if pick {
		res, err := api.PickOrders(sid, canteen)
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusOK, res)
	} else {
		res, err := api.Orders(sid, canteen)
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusOK, res)
	}
}
