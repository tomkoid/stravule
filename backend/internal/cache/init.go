package cache

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

var RDB *redis.Client

func InitRedis(redisUrl string) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisUrl,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	status := rdb.Ping(context.Background())
	if status.Err() != nil {
		log.Fatalf("redis: %v\n", status.Err())
	}

	log.Println("redis: connected")

	RDB = rdb
}
