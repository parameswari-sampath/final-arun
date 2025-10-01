package config

import (
	"os"
	"strconv"
)

type Config struct {
	AppName        string
	Port           string
	AllowedOrigins string
	DBHost         string
	DBPort         string
	DBUser         string
	DBPassword     string
	DBName         string
	DBMaxIdleConns int
	DBMaxOpenConns int
	DBMaxLifetime  int // in minutes
	JWTSecret      string
	Environment    string
}

func Load() *Config {
	return &Config{
		AppName:        getEnv("APP_NAME", "Final API"),
		Port:           getEnv("PORT", "3000"),
		AllowedOrigins: getEnv("ALLOWED_ORIGINS", "*"),
		DBHost:         getEnv("DB_HOST", "localhost"),
		DBPort:         getEnv("DB_PORT", "5432"),
		DBUser:         getEnv("DB_USER", "mcqadmin"),
		DBPassword:     getEnv("DB_PASSWORD", "mcqpass123"),
		DBName:         getEnv("DB_NAME", "mcqdb"),
		DBMaxIdleConns: getEnvAsInt("DB_MAX_IDLE_CONNS", 10),
		DBMaxOpenConns: getEnvAsInt("DB_MAX_OPEN_CONNS", 100),
		DBMaxLifetime:  getEnvAsInt("DB_MAX_LIFETIME", 60), // 60 minutes default
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

func getEnvAsInt(key string, defaultValue int) int {
	valueStr := os.Getenv(key)
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultValue
}
