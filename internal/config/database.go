package config

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	// Register the PostgreSQL driver.
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

type Database struct {
	*sqlx.DB
	logger *logrus.Logger
}

func NewDatabase(cfg *DatabaseConfig, logger *logrus.Logger) (*Database, error) {
	if cfg == nil {
		cfg = &DatabaseConfig{}
	}

	host := cfg.Host
	if host == "" {
		host = "postgres"
	}
	port := cfg.Port
	if port == 0 {
		port = 5432
	}
	user := cfg.User
	if user == "" {
		user = "chatgpt"
	}
	password := cfg.Password
	dbname := cfg.Name
	if dbname == "" {
		dbname = "chatgpt_creator"
	}
	sslmode := cfg.SSLMode
	if sslmode == "" {
		sslmode = "disable"
	}

	logger.Infof("Database config: host=%s port=%d", host, port)

	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbname, sslmode,
	)

	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	maxOpenConns := cfg.MaxOpenConns
	if maxOpenConns == 0 {
		maxOpenConns = 25
	}
	maxIdleConns := cfg.MaxIdleConns
	if maxIdleConns == 0 {
		maxIdleConns = 25
	}
	connMaxLifetime := cfg.ConnMaxLifetime
	if connMaxLifetime == 0 {
		connMaxLifetime = 5 * time.Minute
	}

	db.SetMaxOpenConns(maxOpenConns)
	db.SetMaxIdleConns(maxIdleConns)
	db.SetConnMaxLifetime(connMaxLifetime)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	logger.Info("Successfully connected to PostgreSQL database")

	return &Database{DB: db, logger: logger}, nil
}

func (db *Database) Close() error {
	db.logger.Info("Closing database connection")
	return db.DB.Close()
}

func (db *Database) Ping() error {
	return db.DB.Ping()
}

// RunMigrations executes SQL migration files from the given directory.
func (db *Database) RunMigrations(migrationsDir string) error {
	entries, err := os.ReadDir(migrationsDir)
	if err != nil {
		return fmt.Errorf("failed to read migrations directory: %w", err)
	}

	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".up.sql") {
			continue
		}
		path := filepath.Join(migrationsDir, entry.Name())
		content, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("failed to read migration %s: %w", entry.Name(), err)
		}
		if _, err := db.DB.Exec(string(content)); err != nil {
			// Ignore "already exists" errors for idempotency
			if strings.Contains(err.Error(), "already exists") {
				db.logger.Debugf("Migration %s: objects already exist, skipping", entry.Name())
				continue
			}
			return fmt.Errorf("failed to execute migration %s: %w", entry.Name(), err)
		}
		db.logger.Infof("Applied migration: %s", entry.Name())
	}
	return nil
}
