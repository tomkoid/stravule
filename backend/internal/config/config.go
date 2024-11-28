package config

import "os"

type Config struct {
	Host            string
	Port            string
	DatabaseURL     string
	RedisURL        string
	BetatestersOnly bool
}

var Cfg Config

func LoadConfig() {
	Cfg = GetConfig()
}

func GetConfig() Config {
	return Config{
		Host:            os.Getenv("HOST"),
		DatabaseURL:     os.Getenv("DATABASE_URL"),
		RedisURL:        os.Getenv("REDIS_URL"),
		BetatestersOnly: os.Getenv("BETATESTERS_ONLY") == "true",
	}
}
