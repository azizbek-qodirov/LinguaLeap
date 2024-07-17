package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	HTTP_PORT             string
	LEARNING_SERVICE_PORT string
	PROGRESS_SERVICE_PORT string
}

func Load() Config {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}

	config := Config{}

	config.HTTP_PORT = cast.ToString(coalesce("HTTP_USER_PORT", ":8080"))
	config.LEARNING_SERVICE_PORT = cast.ToString(coalesce("LEARNING_SERVICE_PORT", ":50051"))
	config.PROGRESS_SERVICE_PORT = cast.ToString(coalesce("LEARNING_SERVICE_PORT", ":50052"))

	return config
}

func coalesce(key string, defaultValue interface{}) interface{} {
	val, exists := os.LookupEnv(key)
	if exists {
		return val
	}
	return defaultValue
}
