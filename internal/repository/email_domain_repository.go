package repository

import (
	"context"
	"database/sql"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/verssache/chatgpt-creator/internal/models"
)

type EmailDomainRepository interface {
	Create(ctx context.Context, domain *models.EmailDomain) error
	GetByID(ctx context.Context, id string) (*models.EmailDomain, error)
	GetByDomain(ctx context.Context, domain string) (*models.EmailDomain, error)
	List(ctx context.Context, activeOnly bool) ([]*models.EmailDomain, error)
	Update(ctx context.Context, domain *models.EmailDomain) error
	Delete(ctx context.Context, id string) error
	GetNextDomain(ctx context.Context) (*models.EmailDomain, error)
}

type emailDomainRepository struct {
	*BaseRepository
}

func NewEmailDomainRepository(db *sqlx.DB) EmailDomainRepository {
	return &emailDomainRepository{NewBaseRepository(db)}
}

func (r *emailDomainRepository) Create(ctx context.Context, domain *models.EmailDomain) error {
	query := `
		INSERT INTO email_domains (id, domain, priority, is_active, source, last_checked, health_status, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	`
	_, err := r.db.ExecContext(ctx, query,
		domain.ID,
		domain.Domain,
		domain.Priority,
		domain.IsActive,
		domain.Source,
		domain.LastChecked,
		domain.HealthStatus,
		domain.CreatedAt,
		domain.UpdatedAt,
	)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") || strings.Contains(err.Error(), "unique constraint") {
			return ErrDuplicate
		}
		return err
	}
	return nil
}

func (r *emailDomainRepository) GetByID(ctx context.Context, id string) (*models.EmailDomain, error) {
	query := `SELECT id, domain, priority, is_active, source, last_checked, health_status, created_at, updated_at FROM email_domains WHERE id = $1`
	var domain models.EmailDomain
	err := r.db.GetContext(ctx, &domain, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &domain, nil
}

func (r *emailDomainRepository) GetByDomain(ctx context.Context, domainName string) (*models.EmailDomain, error) {
	query := `SELECT id, domain, priority, is_active, source, last_checked, health_status, created_at, updated_at FROM email_domains WHERE domain = $1`
	var domain models.EmailDomain
	err := r.db.GetContext(ctx, &domain, query, domainName)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &domain, nil
}

func (r *emailDomainRepository) List(ctx context.Context, activeOnly bool) ([]*models.EmailDomain, error) {
	query := `SELECT id, domain, priority, is_active, source, last_checked, health_status, created_at, updated_at FROM email_domains`
	if activeOnly {
		query += " WHERE is_active = true"
	}
	query += " ORDER BY priority ASC, created_at DESC"

	var domains []*models.EmailDomain
	err := r.db.SelectContext(ctx, &domains, query)
	if err != nil {
		return nil, err
	}
	return domains, nil
}

func (r *emailDomainRepository) Update(ctx context.Context, domain *models.EmailDomain) error {
	query := `
		UPDATE email_domains
		SET domain = $1, priority = $2, is_active = $3, source = $4, last_checked = $5, health_status = $6, updated_at = $7
		WHERE id = $8
	`
	result, err := r.db.ExecContext(ctx, query,
		domain.Domain,
		domain.Priority,
		domain.IsActive,
		domain.Source,
		domain.LastChecked,
		domain.HealthStatus,
		domain.UpdatedAt,
		domain.ID,
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

func (r *emailDomainRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM email_domains WHERE id = $1`
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

func (r *emailDomainRepository) GetNextDomain(ctx context.Context) (*models.EmailDomain, error) {
	query := `
		SELECT id, domain, priority, is_active, source, last_checked, health_status, created_at, updated_at
		FROM email_domains
		WHERE is_active = true AND health_status = $1
		ORDER BY priority ASC, RANDOM()
		LIMIT 1
	`
	var domain models.EmailDomain
	err := r.db.GetContext(ctx, &domain, query, models.DomainHealthStatusHealthy)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &domain, nil
}
