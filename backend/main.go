package main

import (
	"log"

	"codeberg.org/tomkoid/stravule/backend/internal/db"
	"codeberg.org/tomkoid/stravule/backend/internal/server"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	err = db.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	server.StartServer()
}
