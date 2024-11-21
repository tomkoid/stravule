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
	userHash := c.QueryParam("user_hash")

	if pickQuery != "" {
		pickTmp, err := strconv.ParseBool(pickQuery)
		if err != nil {
			return c.String(http.StatusBadRequest, "Invalid pick query parameter")
		}

		pick = pickTmp
	}

	if pick {
		res, _, err := api.PickOrders(sid, canteen, userHash)
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

func SendOrders(c echo.Context) error {
	sid := c.QueryParam("sid")
	canteen := c.QueryParam("canteen")
	userHash := c.QueryParam("user_hash")

	err := api.SendOrders(api.SendOrdersRequest{
		SID:   sid,
		Cislo: canteen,
		Lang:  "CZ",
		Url:   "",
	}, userHash)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.String(http.StatusOK, "sent")
}
