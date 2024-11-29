package resolvers

import (
	"context"
	"errors"
	"strconv"

	"codeberg.org/tomkoid/stravule/backend/db"
	"codeberg.org/tomkoid/stravule/backend/internal/database"
)

func AddOrderDayException(userHash *string, orderDayException int) error {
	exists, err := userHasOrderDayException(userHash, orderDayException)
	if err != nil {
		return err
	}

	if exists {
		return errors.New("order day exception already exists")
	}

	return database.DB.AddWeekdayOrderingException(context.Background(), db.AddWeekdayOrderingExceptionParams{
		UserHash:    *userHash,
		ArrayAppend: orderDayException,
	})
}

func RemoveOrderDayException(userHash *string, orderDayException int) error {
	exists, err := userHasOrderDayException(userHash, orderDayException)
	if err != nil {
		return err
	}

	if !exists {
		return errors.New("order day exception doesn't exist")
	}

	return database.DB.RemoveWeekdayOrderingException(context.Background(), db.RemoveWeekdayOrderingExceptionParams{
		UserHash:    *userHash,
		ArrayRemove: orderDayException,
	})
}

func userHasOrderDayException(userHash *string, orderDayException int) (bool, error) {
	user, err := database.DB.GetUser(context.Background(), *userHash)
	if err != nil {
		return false, errors.New("user not found")
	}

	for _, exception := range user.OrderDaysExceptions {
		exceptionInt, err := strconv.Atoi(exception)
		if err != nil {
			return false, err
		}

		if exceptionInt == orderDayException {
			return true, nil
		}
	}

	return false, nil
}
