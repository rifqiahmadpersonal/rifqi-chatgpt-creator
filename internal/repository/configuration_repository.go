package repository

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
	"rifqi-chatgpt-creator/internal/models"
)

type ConfigurationRepository interface {
	Upsert(ctx context.Context, config *models.Configuration) error
	GetByKey(ctx context.Context, key string) (*models.Configuration, error)
	List(ctx context.Context) ([]*models.Configuration, error)
	Delete(ctx context.Context, key string) error
}

type configurationRepository struct {
	*BaseRepository
}

func NewConfigurationRepository(db *sqlx.DB) ConfigurationRepository {
	return &configurationRepository{NewBaseRepository(db)}
}

func (r *configurationRepository) Upsert(ctx context.Context, config *models.Configuration) error {
	query := `
		INSERT INTO configurations (id, key, value, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5)
		ON CONFLICT (key) DO UPDATE SET value = $3, updated_at = $5
	`
	_, err := r.db.ExecContext(ctx, query,
		config.ID,
		config.Key,
		config.Value,
		config.CreatedAt,
		config.UpdatedAt,
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *configurationRepository) GetByKey(ctx context.Context, key string) (*models.Configuration, error) {
	query := `SELECT id, key, value, created_at, updated_at FROM configurations WHERE key = $1`
	var config models.Configuration
	err := r.db.GetContext(ctx, &config, query, key)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &config, nil
}

func (r *configurationRepository) List(ctx context.Context) ([]*models.Configuration, error) {
	query := `SELECT id, key, value, created_at, updated_at FROM configurations ORDER BY key ASC`
	var configs []*models.Configuration
	err := r.db.SelectContext(ctx, &configs, query)
	if err != nil {
		return nil, err
	}
	return configs, nil
}

func (r *configurationRepository) Delete(ctx context.Context, key string) error {
	query := `DELETE FROM configurations WHERE key = $1`
	result, err := r.db.ExecContext(ctx, query, key)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return ErrNotFound
	}
	return nil
}
