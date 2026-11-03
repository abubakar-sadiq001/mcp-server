package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	APIURL    string // http://localhost:8080/api/v1
	AuthToken string // Bearer <token>
	LogLevel  string
	Transport string // stdio -> not implemented yet, http
}

func Load() (*Config, error) {
	_ = godotenv.Load()

	authToken := os.Getenv("AUTH_TOKEN")
	if authToken == "" {
		authToken = os.Getenv("JWT_TOKEN")
	}

	config := &Config{
		APIURL:    getEnv("API_URL", "http://localhost:8080/api/v1"),
		AuthToken: authToken,
		LogLevel:  getEnv("LOG_LEVEL", "debug"),
		Transport: getEnv("TRANSPORT", "stdio"),
	}

	return config, nil
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
