package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

// Config holds the database configuration
type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

// NewConfigFromEnv creates a Config from environment variables with defaults
func NewConfigFromEnv() Config {
	return Config{
		Host:     getEnv("DB_HOST", "postgres"),
		Port:     getEnv("DB_PORT", "5432"),
		User:     getEnv("DB_USER", "postgres"),
		Password: getEnv("DB_PASSWORD", "postgres"),
		DBName:   getEnv("DB_NAME", "postgres"),
	}
}

// Connect establishes a connection to the PostgreSQL database
func Connect(cfg Config) (*sql.DB, error) {
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return db, nil
}

// MustConnect connects to the database or panics on failure
func MustConnect(cfg Config) *sql.DB {
	db, err := Connect(cfg)
	if err != nil {
		panic(err)
	}
	return db
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
