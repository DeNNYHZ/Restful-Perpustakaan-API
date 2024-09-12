package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

// Config holds the application configuration
type Config struct {
	// Server configuration
	ServerAddress string
	ServerPort    int

	// Database configuration
	DBHost     string
	DBPort     int
	DBUser     string
	DBPassword string
	DBName     string

	// JWT configuration
	JWTSecretKey      string
	JWTExpirationTime int

	// Email configuration (for notifications)
	EmailHost     string
	EmailPort     int
	EmailUser     string
	EmailPassword string
	EmailFrom     string

	// External service configuration (example)
	ExternalServiceAPIKey  string
	ExternalServiceBaseURL string
	DatabaseURL            string
}

// LoadConfig loads configuration from environment variables or a .env file
func LoadConfig() (*Config, error) {
	// Load environment variables from .env file (if it exists)
	err := godotenv.Load()
	if err != nil {
		// .env file is optional, so we can ignore the error if it doesn't exist
		fmt.Println("No .env file found, loading configuration from environment variables")
	}

	// Get configuration values from environment variables
	cfg := &Config{
		ServerAddress: getEnv("SERVER_ADDRESS", "localhost"),
		ServerPort:    getEnvAsInt("SERVER_PORT", 8080),

		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnvAsInt("DB_PORT", 5432),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", "admin"),
		DBName:     getEnv("DB_NAME", "perpustakaan_db"),

		JWTSecretKey:      getEnv("JWT_SECRET_KEY", "your_secret_key"),
		JWTExpirationTime: getEnvAsInt("JWT_EXPIRATION_TIME", 3600), // 1 hour in seconds

		// Email configuration
		EmailHost:     getEnv("EMAIL_HOST", "smtp.example.com"),
		EmailPort:     getEnvAsInt("EMAIL_PORT", 587),
		EmailUser:     getEnv("EMAIL_USER", "your_email_username"),
		EmailPassword: getEnv("EMAIL_PASSWORD", "your_email_password"),
		EmailFrom:     getEnv("EMAIL_FROM", "your_email_address"),

		// External service configuration
		ExternalServiceAPIKey:  getEnv("EXTERNAL_SERVICE_API_KEY", "your_api_key"),
		ExternalServiceBaseURL: getEnv("EXTERNAL_SERVICE_BASE_URL", "https://api.example.com"),
	}

	return cfg, nil
}

// getEnv gets an environment variable or returns a default value if it's not set
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

// getEnvAsInt gets an environment variable as an integer or returns a default value if it's not set or invalid
func getEnvAsInt(key string, defaultValue int) int {
	valueStr := os.Getenv(key)
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultValue
}
