package config

import (
	"log"
	"restapi_go/pkg/env"
	"strconv"
)

type config struct {
	ApiPort int
	ApiName string
}

// This object will be used throughout the application to access the configuration values.
var cfg config

func LoadConfig() config {
	env.LoadEnv()

	// Convert int vars
	apiPort, err := strconv.Atoi(env.GetEnv("API_PORT"))
	if err != nil {
		log.Fatalf("API_PORT property not set correctly: %v", err)
	}

	// Returns the configuration object with the values from the environment variables.
	cfg = config{
		ApiPort: apiPort,
		ApiName: env.GetEnv("API_NAME"),
	}

	return GetConfig()
}

func GetConfig() config {
	return cfg
}
