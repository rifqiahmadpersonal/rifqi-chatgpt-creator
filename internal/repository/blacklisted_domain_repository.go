package repository

import (
	"context"
	"database/sql"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/verssache/chatgpt-creator/internal/models"
)

type BlacklistedDomainRepository interface {
	Create(ctx context.Context, domain *models.BlacklistedDomain) error
	GetByDomain(ctx context.Context, domain string) (*models.BlacklistedDomain, error)
	List(ctx context.Context) ([]*models.BlacklistedDomain, error)
	Delete(ctx context.Context, id string) error
	Exists(ctx context.Context, domain string) (bool, error)
}

type blacklistedDomainRepository struct {
	*BaseRepository
}

func NewBlacklistedDomainRepository(db *sqlx.DB) BlacklistedDomainRepository {
	return &blacklistedDomainRepository{NewBaseRepository(db)}
}

func (r *blacklistedDomainRepository) Create(ctx context.Context, domain *models.BlacklistedDomain) error {
	query := `
		INSERT INTO blacklisted_domains (id, domain, reason, created_at)
		VALUES ($1, $2, $3, $4)
	`
	_, err := r.db.ExecContext(ctx, query,
		domain.ID,
		domain.Domain,
		domain.Reason,
		domain.CreatedAt,
	)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") || strings.Contains(err.Error(), "unique constraint") {
			return ErrDuplicate
		}
		return err
	}
	return nil
}

func (r *blacklistedDomainRepository) GetByDomain(ctx context.Context, domainName string) (*models.BlacklistedDomain, error) {
	query := `SELECT id, domain, reason, created_at FROM blacklisted_domains WHERE domain = $1`
	var domain models.BlacklistedDomain
	err := r.db.GetContext(ctx, &domain, query, domainName)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &domain, nil
}

func (r *blacklistedDomainRepository) List(ctx context.Context) ([]*models.BlacklistedDomain, error) {
	query := `SELECT id, domain, reason, created_at FROM blacklisted_domains ORDER BY created_at DESC`
	var domains []*models.BlacklistedDomain
	err := r.db.SelectContext(ctx, &domains, query)
	if err != nil {
		return nil, err
	}
	return domains, nil
}

func (r *blacklistedDomainRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM blacklisted_domains WHERE id = $1`
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

func (r *blacklistedDomainRepository) Exists(ctx context.Context, domain string) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM blacklisted_domains WHERE domain = $1)`
	var exists bool
	err := r.db.GetContext(ctx, &exists, query, domain)
	if err != nil {
		return false, err
	}
	return exists, nil
}
