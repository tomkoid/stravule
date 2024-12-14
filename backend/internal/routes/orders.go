package routes

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"codeberg.org/tomkoid/stravule/backend/internal/api"
	"codeberg.org/tomkoid/stravule/backend/internal/cache"
	"codeberg.org/tomkoid/stravule/backend/internal/resolvers"
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

	cache.RDB.Del(context.Background(), fmt.Sprintf("orders:%s:%s", sid, canteen)).Val()

	return c.String(http.StatusOK, "sent")
}

func AddOrderDayException(c echo.Context) error {
	userHash := c.QueryParam("user_hash")
	weekDay := c.QueryParam("week_day")

	if weekDay == "" {
		return c.String(http.StatusBadRequest, "Missing `week_day` parameter")
	}

	weekday, err := strconv.Atoi(weekDay)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	err = resolvers.AddOrderDayException(&userHash, weekday)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.String(http.StatusOK, "added")
}

func AddNoOrderDay(c echo.Context) error {
	userHash := c.QueryParam("user_hash")
	day := c.QueryParam("day")

	if day == "" {
		return c.String(http.StatusBadRequest, "Missing `day` parameter")
	}

	err := resolvers.AddNoOrderDay(&userHash, day)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.String(http.StatusOK, "added")
}

func RemoveOrderDayException(c echo.Context) error {
	userHash := c.QueryParam("user_hash")
	weekDay := c.QueryParam("week_day")

	if weekDay == "" {
		return c.String(http.StatusBadRequest, "Missing `week_day` parameter")
	}

	weekday, err := strconv.Atoi(weekDay)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	err = resolvers.RemoveOrderDayException(&userHash, weekday)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.String(http.StatusOK, "removed")
}

func RemoveNoOrderDay(c echo.Context) error {
	userHash := c.QueryParam("user_hash")
	day := c.QueryParam("day")

	if day == "" {
		return c.String(http.StatusBadRequest, "Missing `day` parameter")
	}

	err := resolvers.RemoveNoOrderDay(&userHash, day)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.String(http.StatusOK, "removed")
}

func GetOrderDayExceptions(c echo.Context) error {
	userHash := c.QueryParam("user_hash")

	orderDayExceptions, err := resolvers.ListOrderDayExceptions(&userHash)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, orderDayExceptions)
}

func ListNoOrderDays(c echo.Context) error {
	userHash := c.QueryParam("user_hash")

	noOrderDays, err := resolvers.ListNoOrderDays(&userHash)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, noOrderDays)
}
