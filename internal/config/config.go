package config

import (
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	JWT      JWTConfig
	App      AppConfig
}

type ServerConfig struct {
	Port string
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

type JWTConfig struct {
	Secret               string
	AccessTokenDuration  time.Duration
	RefreshTokenDuration time.Duration
}

type AppConfig struct {
	Environment string
}

// LoadConfig membaca environment variables dan return Config struct
func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("error loading .env file: %w", err)
	}

	// Parse JWT token durations
	accessDuration, err := time.ParseDuration(getEnv("JWT_ACCESS_TOKEN_DURATION", "15m"))
	if err != nil {
		return nil, fmt.Errorf("invalid JWT_ACCESS_TOKEN_DURATION: %w", err)
	}

	refreshDuration, err := time.ParseDuration(getEnv("JWT_REFRESH_TOKEN_DURATION", "24h"))
	if err != nil {
		return nil, fmt.Errorf("invalid JWT_REFRESH_TOKEN_DURATION: %w", err)
	}

	config := &Config{
		Server: ServerConfig{
			Port: getEnv("PORT", "8080"),
		},
		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "5432"),
			User:     getEnv("DB_USER", "postgres"),
			Password: getEnv("DB_PASSWORD", ""),
			DBName:   getEnv("DB_NAME", "auth_service_db"),
			SSLMode:  getEnv("DB_SSLMODE", "disable"),
		},
		JWT: JWTConfig{
			Secret:               getEnv("JWT_SECRET", ""),
			AccessTokenDuration:  accessDuration,
			RefreshTokenDuration: refreshDuration,
		},
		App: AppConfig{
			Environment: getEnv("ENV", "development"),
		},
	}

	// Validate required fields
	if config.JWT.Secret == "" {
		return nil, fmt.Errorf("JWT_SECRET is required")
	}
	if config.Database.Password == "" {
		return nil, fmt.Errorf("DB_PASSWORD is required")
	}

	return config, nil
}

// GetDSN returns database connection string (seperti JDBC URL)
func (c *DatabaseConfig) GetDSN() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.Host, c.Port, c.User, c.Password, c.DBName, c.SSLMode,
	)
}

// Helper function untuk get env dengan default value
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
