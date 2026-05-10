package repository

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/verssache/chatgpt-creator/internal/models"
)

type RegistrationAttemptRepository interface {
	Create(ctx context.Context, attempt *models.RegistrationAttempt) error
	GetByBatchJobID(ctx context.Context, batchJobID string) ([]*models.RegistrationAttempt, error)
	Update(ctx context.Context, attempt *models.RegistrationAttempt) error
}

type registrationAttemptRepository struct {
	*BaseRepository
}

func NewRegistrationAttemptRepository(db *sqlx.DB) RegistrationAttemptRepository {
	return &registrationAttemptRepository{NewBaseRepository(db)}
}

func (r *registrationAttemptRepository) Create(ctx context.Context, attempt *models.RegistrationAttempt) error {
	query := `
		INSERT INTO registration_attempts (id, email, status, error_message, worker_id, batch_job_id, started_at, completed_at, duration_ms)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	`
	_, err := r.db.ExecContext(ctx, query,
		attempt.ID,
		attempt.Email,
		attempt.Status,
		attempt.ErrorMessage,
		attempt.WorkerID,
		attempt.BatchJobID,
		attempt.StartedAt,
		attempt.CompletedAt,
		attempt.Duration,
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *registrationAttemptRepository) GetByBatchJobID(ctx context.Context, batchJobID string) ([]*models.RegistrationAttempt, error) {
	query := `SELECT id, email, status, error_message, worker_id, batch_job_id, started_at, completed_at, duration_ms FROM registration_attempts WHERE batch_job_id = $1 ORDER BY started_at DESC`
	var attempts []*models.RegistrationAttempt
	err := r.db.SelectContext(ctx, &attempts, query, batchJobID)
	if err != nil {
		return nil, err
	}
	return attempts, nil
}

func (r *registrationAttemptRepository) Update(ctx context.Context, attempt *models.RegistrationAttempt) error {
	query := `
		UPDATE registration_attempts
		SET email = $1, status = $2, error_message = $3, worker_id = $4, batch_job_id = $5, started_at = $6, completed_at = $7, duration_ms = $8
		WHERE id = $9
	`
	result, err := r.db.ExecContext(ctx, query,
		attempt.Email,
		attempt.Status,
		attempt.ErrorMessage,
		attempt.WorkerID,
		attempt.BatchJobID,
		attempt.StartedAt,
		attempt.CompletedAt,
		attempt.Duration,
		attempt.ID,
	)
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
