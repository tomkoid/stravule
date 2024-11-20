package main

import (
	"log"

	"codeberg.org/tomkoid/stravule/backend/internal/database"
	"codeberg.org/tomkoid/stravule/backend/internal/server"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("env: error loading .env file")
	}

	err = database.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	server.StartServer()
}
