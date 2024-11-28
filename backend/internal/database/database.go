package database

import (
	"context"
	"errors"
	"fmt"

	"log"

	"codeberg.org/tomkoid/stravule/backend/db"
	"codeberg.org/tomkoid/stravule/backend/internal/config"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *db.Queries

func InitDB() error {
	dbUrl := config.Cfg.DatabaseURL

	// migrations
	m, err := migrate.New(
		"file://misc/migrations",
		dbUrl)
	if err != nil {
		log.Fatalf("db: can't apply migrations: %v\n", err)
	}

	if err := m.Up(); err != nil {
		log.Printf("db: %s\n", err)
	}

	if dbUrl == "" {
		return errors.New("DATABASE_URL is not set in the .env file or as an environment variable")
	}

	ctx := context.Background()
	freshdb, err := pgxpool.New(ctx, dbUrl)

	if err != nil {
		return errors.New(fmt.Sprintf("unable to connect to database: %v\n", err))
	}

	err = freshdb.Ping(ctx)
	if err != nil {
		return errors.New(fmt.Sprintf("unable to ping the database: %v\n", err))
	}

	log.Printf("db: connected to the database successfully.")

	DB = db.New(freshdb)

	if DB == nil {
		log.Fatal("db: database is nil")
	}

	return nil
}
