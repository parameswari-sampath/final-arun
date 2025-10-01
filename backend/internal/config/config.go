package config

import (
	"os"
)

type Config struct {
	AppName        string
	Port           string
	AllowedOrigins string
	DatabaseURL    string
	JWTSecret      string
	Environment    string
}

func Load() *Config {
	return &Config{
		AppName:        getEnv("APP_NAME", "Final API"),
		Port:           getEnv("PORT", "3000"),
		AllowedOrigins: getEnv("ALLOWED_ORIGINS", "*"),
		DatabaseURL:    getEnv("DATABASE_URL", ""),
		JWTSecret:      getEnv("JWT_SECRET", "your-secret-key-change-in-production"),
		Environment:    getEnv("ENVIRONMENT", "development"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
