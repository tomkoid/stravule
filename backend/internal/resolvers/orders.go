package resolvers

import (
	"context"
	"errors"

	"codeberg.org/tomkoid/stravule/backend/db"
	"codeberg.org/tomkoid/stravule/backend/internal/database"
)

func ListOrderDayExceptions(userHash *string) ([]int32, error) {
	user, err := database.DB.GetUser(context.Background(), *userHash)
	if err != nil {
		return nil, errors.New("user not found")
	}

	return user.OrderDaysExceptions, nil
}

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
		if exception == int32(orderDayException) {
			return true, nil
		}
	}

	return false, nil
}
