package config

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

type Database struct {
	*sqlx.DB
	logger *logrus.Logger
}

func NewDatabase(cfg *DatabaseConfig, logger *logrus.Logger) (*Database, error) {
	host := getEnvOrDefault("DB_HOST", "postgres")
	port := getEnvOrDefaultInt("DB_PORT", 5432)
	user := getEnvOrDefault("DB_USER", "chatgpt")
	password := getEnvOrDefault("DB_PASSWORD", "chatgpt_secret_change_me")
	dbname := getEnvOrDefault("DB_NAME", "chatgpt_creator")
	sslmode := getEnvOrDefault("DB_SSL_MODE", "disable")

	logger.Infof("Database config: host=%s port=%d", host, port)

	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbname, sslmode,
	)

	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	maxOpenConns := getEnvOrDefaultInt("DB_MAX_OPEN_CONNS", 25)
	maxIdleConns := getEnvOrDefaultInt("DB_MAX_IDLE_CONNS", 25)
	connMaxLifetime := getEnvOrDefault("DB_CONN_MAX_LIFETIME", "5m")

	db.SetMaxOpenConns(maxOpenConns)
	db.SetMaxIdleConns(maxIdleConns)
	
	// Parse duration
	duration, _ := time.ParseDuration(connMaxLifetime)
	if duration > 0 {
		db.SetConnMaxLifetime(duration)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	logger.Info("Successfully connected to PostgreSQL database")

	return &Database{DB: db, logger: logger}, nil
}

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvOrDefaultInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intVal, err := strconv.Atoi(value); err == nil {
			return intVal
		}
	}
	return defaultValue
}

func (db *Database) Close() error {
	db.logger.Info("Closing database connection")
	return db.DB.Close()
}

func (db *Database) Ping() error {
	return db.DB.Ping()
}
