package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Loads environment variables from a '.env' file.
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func GetEnv(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		log.Fatalf("Environment variable %s is required but not set", key)
	}
	return value
}
