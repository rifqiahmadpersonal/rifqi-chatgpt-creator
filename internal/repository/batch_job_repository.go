package repository

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
	"rifqi-chatgpt-creator/internal/models"
)

type BatchJobRepository interface {
	Create(ctx context.Context, job *models.BatchJob) error
	GetByID(ctx context.Context, id string) (*models.BatchJob, error)
	List(ctx context.Context, status string) ([]*models.BatchJob, error)
	Update(ctx context.Context, job *models.BatchJob) error
	IncrementSuccess(ctx context.Context, id string) error
	IncrementFailure(ctx context.Context, id string) error
}

type batchJobRepository struct {
	*BaseRepository
}

func NewBatchJobRepository(db *sqlx.DB) BatchJobRepository {
	return &batchJobRepository{NewBaseRepository(db)}
}

func (r *batchJobRepository) Create(ctx context.Context, job *models.BatchJob) error {
	query := `
		INSERT INTO batch_jobs (id, target_count, success_count, failure_count, status, max_workers, default_password, proxy, created_at, completed_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	`
	_, err := r.db.ExecContext(ctx, query,
		job.ID,
		job.TargetCount,
		job.SuccessCount,
		job.FailureCount,
		job.Status,
		job.MaxWorkers,
		job.DefaultPassword,
		job.Proxy,
		job.CreatedAt,
		job.CompletedAt,
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *batchJobRepository) GetByID(ctx context.Context, id string) (*models.BatchJob, error) {
	query := `SELECT id, target_count, success_count, failure_count, status, max_workers, default_password, proxy, created_at, completed_at FROM batch_jobs WHERE id = $1`
	var job models.BatchJob
	err := r.db.GetContext(ctx, &job, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &job, nil
}

func (r *batchJobRepository) List(ctx context.Context, status string) ([]*models.BatchJob, error) {
	query := `SELECT id, target_count, success_count, failure_count, status, max_workers, default_password, proxy, created_at, completed_at FROM batch_jobs`
	args := []interface{}{}

	if status != "" {
		query += " WHERE status = $1"
		args = append(args, status)
	}

	query += " ORDER BY created_at DESC"

	var jobs []*models.BatchJob
	err := r.db.SelectContext(ctx, &jobs, query, args...)
	if err != nil {
		return nil, err
	}
	return jobs, nil
}

func (r *batchJobRepository) Update(ctx context.Context, job *models.BatchJob) error {
	query := `
		UPDATE batch_jobs
		SET target_count = $1, success_count = $2, failure_count = $3, status = $4, max_workers = $5, default_password = $6, proxy = $7, completed_at = $8
		WHERE id = $9
	`
	result, err := r.db.ExecContext(ctx, query,
		job.TargetCount,
		job.SuccessCount,
		job.FailureCount,
		job.Status,
		job.MaxWorkers,
		job.DefaultPassword,
		job.Proxy,
		job.CompletedAt,
		job.ID,
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

func (r *batchJobRepository) IncrementSuccess(ctx context.Context, id string) error {
	query := `UPDATE batch_jobs SET success_count = success_count + 1 WHERE id = $1`
	result, err := r.db.ExecContext(ctx, query, id)
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

func (r *batchJobRepository) IncrementFailure(ctx context.Context, id string) error {
	query := `UPDATE batch_jobs SET failure_count = failure_count + 1 WHERE id = $1`
	result, err := r.db.ExecContext(ctx, query, id)
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
