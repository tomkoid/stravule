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

func ListNoOrderDays(userHash *string) ([]string, error) {
	user, err := database.DB.GetUser(context.Background(), *userHash)
	if err != nil {
		return nil, errors.New("user not found")
	}

	return user.NoOrderDays, nil
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

func AddNoOrderDay(userHash *string, day string) error {
	exists, err := userHasNoOrderDay(userHash, day)
	if err != nil {
		return err
	}

	if exists {
		return errors.New("no order day already exists")
	}

	return database.DB.AddNoOrderDay(context.Background(), db.AddNoOrderDayParams{
		UserHash:    *userHash,
		ArrayAppend: day,
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

func RemoveNoOrderDay(userHash *string, day string) error {
	exists, err := userHasNoOrderDay(userHash, day)
	if err != nil {
		return err
	}

	if !exists {
		return errors.New("no order day doesn't exist")
	}

	return database.DB.RemoveNoOrderDay(context.Background(), db.RemoveNoOrderDayParams{
		UserHash:    *userHash,
		ArrayRemove: day,
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

func userHasNoOrderDay(userHash *string, day string) (bool, error) {
	user, err := database.DB.GetUser(context.Background(), *userHash)
	if err != nil {
		return false, errors.New("user not found")
	}

	for _, noOrderDay := range user.NoOrderDays {
		if noOrderDay == day {
			return true, nil
		}
	}

	return false, nil
}
