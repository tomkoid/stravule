package api

import (
	"context"
	"errors"

	"codeberg.org/tomkoid/stravule/backend/internal/database"
)

func RegisterBetatester(userHash string, registerHash string) error {
	ctx := context.Background()

	// the person who already registered can add the betatester
	moderator, err := database.DB.GetUser(ctx, userHash)

	if err != nil {
		return err
	}

	if moderator.IsBetaTester {
		wantedToBeAdded, err := database.DB.GetUser(ctx, registerHash)

		if err != nil {
			return err
		}

		if wantedToBeAdded.IsBetaTester {
			return errors.New("The user is already a betatester")
		}

		return database.DB.RegisterBetatester(ctx, registerHash)
	}

	return nil
}
