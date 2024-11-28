package main

import (
	"log"

	"codeberg.org/tomkoid/stravule/backend/internal/cache"
	"codeberg.org/tomkoid/stravule/backend/internal/config"
	"codeberg.org/tomkoid/stravule/backend/internal/database"
	"codeberg.org/tomkoid/stravule/backend/internal/server"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("env: error loading .env file")
	}

	// load config into config.Cfg
	config.LoadConfig()

	err = database.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	redisUrl := config.Cfg.RedisURL
	if redisUrl != "" {
		cache.InitRedis(redisUrl)
	}

	server.StartServer()
}
