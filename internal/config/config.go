package config

import (
	"fmt"
	"os"
)

// Config holds the application configuration
type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBNAME     string
}

// LoadConfig loads configuration from environment variables
func LoadConfig() (*Config, error) {
	cfg := &Config{
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBNAME:     os.Getenv("DB_NAME"),
	}

	// Validate required fields
	if cfg.DBHost == "" || cfg.DBPort == "" || cfg.DBUser == "" || cfg.DBPassword == "" || cfg.DBNAME == "" {
		return nil, fmt.Errorf("missing required environment variables")
	}

	return cfg, nil
}

// DSN returns the Data Source Name for PostgreSQL
func (c *Config) DSN() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		c.DBHost, c.DBPort, c.DBUser, c.DBPassword, c.DBNAME)
}
