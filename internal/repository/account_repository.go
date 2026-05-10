package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/verssache/chatgpt-creator/internal/models"
)

type AccountRepository interface {
	Create(ctx context.Context, account *models.Account) error
	GetByID(ctx context.Context, id string) (*models.Account, error)
	GetByEmail(ctx context.Context, email string) (*models.Account, error)
	List(ctx context.Context, filter *models.AccountFilter) ([]*models.Account, error)
	Update(ctx context.Context, account *models.Account) error
	Delete(ctx context.Context, id string) error
	Count(ctx context.Context) (int64, error)
	CountByStatus(ctx context.Context, status string) (int64, error)
}

type accountRepository struct {
	*BaseRepository
}

func NewAccountRepository(db *sqlx.DB) AccountRepository {
	return &accountRepository{NewBaseRepository(db)}
}

func (r *accountRepository) Create(ctx context.Context, account *models.Account) error {
	query := `
		INSERT INTO accounts (id, email, password, status, batch_job_id, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`
	_, err := r.db.ExecContext(ctx, query,
		account.ID,
		account.Email,
		account.Password,
		account.Status,
		account.BatchJobID,
		account.CreatedAt,
		account.UpdatedAt,
	)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") || strings.Contains(err.Error(), "unique constraint") {
			return ErrDuplicate
		}
		return err
	}
	return nil
}

func (r *accountRepository) GetByID(ctx context.Context, id string) (*models.Account, error) {
	query := `SELECT id, email, password, status, batch_job_id, created_at, updated_at FROM accounts WHERE id = $1`
	var account models.Account
	err := r.db.GetContext(ctx, &account, query, id)
	if err != nil {
		if isNoRowsError(err) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &account, nil
}

func (r *accountRepository) GetByEmail(ctx context.Context, email string) (*models.Account, error) {
	query := `SELECT id, email, password, status, batch_job_id, created_at, updated_at FROM accounts WHERE email = $1`
	var account models.Account
	err := r.db.GetContext(ctx, &account, query, email)
	if err != nil {
		if isNoRowsError(err) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &account, nil
}

func (r *accountRepository) List(ctx context.Context, filter *models.AccountFilter) ([]*models.Account, error) {
	query := `SELECT id, email, password, status, batch_job_id, created_at, updated_at FROM accounts WHERE 1=1`
	args := []interface{}{}
	argPos := 1

	if filter != nil {
		if filter.Status != "" {
			query += fmt.Sprintf(" AND status = $%d", argPos)
			args = append(args, filter.Status)
			argPos++
		}
		if filter.BatchJobID != "" {
			query += fmt.Sprintf(" AND batch_job_id = $%d", argPos)
			args = append(args, filter.BatchJobID)
			argPos++
		}
		if filter.Email != "" {
			query += fmt.Sprintf(" AND email ILIKE $%d", argPos)
			args = append(args, "%"+filter.Email+"%")
			argPos++
		}
	}

	query += " ORDER BY created_at DESC"

	if filter != nil {
		if filter.Limit > 0 {
			query += fmt.Sprintf(" LIMIT $%d", argPos)
			args = append(args, filter.Limit)
			argPos++
		}
		if filter.Offset > 0 {
			query += fmt.Sprintf(" OFFSET $%d", argPos)
			args = append(args, filter.Offset)
		}
	}

	var accounts []*models.Account
	err := r.db.SelectContext(ctx, &accounts, query, args...)
	if err != nil {
		return nil, err
	}
	return accounts, nil
}

func (r *accountRepository) Update(ctx context.Context, account *models.Account) error {
	query := `
		UPDATE accounts
		SET email = $1, password = $2, status = $3, batch_job_id = $4, updated_at = $5
		WHERE id = $6
	`
	result, err := r.db.ExecContext(ctx, query,
		account.Email,
		account.Password,
		account.Status,
		account.BatchJobID,
		account.UpdatedAt,
		account.ID,
	)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") || strings.Contains(err.Error(), "unique constraint") {
			return ErrDuplicate
		}
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

func (r *accountRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM accounts WHERE id = $1`
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

func (r *accountRepository) Count(ctx context.Context) (int64, error) {
	query := `SELECT COUNT(*) FROM accounts`
	var count int64
	err := r.db.GetContext(ctx, &count, query)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *accountRepository) CountByStatus(ctx context.Context, status string) (int64, error) {
	query := `SELECT COUNT(*) FROM accounts WHERE status = $1`
	var count int64
	err := r.db.GetContext(ctx, &count, query, status)
	if err != nil {
		return 0, err
	}
	return count, nil
}
