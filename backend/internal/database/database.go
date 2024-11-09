package database

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

var DB *pgxpool.Pool

func InitDB() error {
	dbUrl := os.Getenv("DATABASE_URL")

	if dbUrl == "" {
		return errors.New("DATABASE_URL is not set in the .env file or as an environment variable")
	}

	ctx := context.Background()
	db, err := pgxpool.New(ctx, dbUrl)

	if err != nil {
		return errors.New(fmt.Sprintf("unable to connect to database: %v\n", err))
	}

	err = db.Ping(ctx)
	if err != nil {
		return errors.New(fmt.Sprintf("unable to ping the database: %v\n", err))
	}

	log.Printf("connected to the database successfully.")

	return nil
}
