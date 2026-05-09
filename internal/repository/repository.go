package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"
)

var (
	ErrNotFound  = errors.New("record not found")
	ErrDuplicate = errors.New("duplicate record")
	ErrInvalidID = errors.New("invalid id")
)

type Repository interface {
	Ping() error
}

type BaseRepository struct {
	db *sqlx.DB
}

func NewBaseRepository(db *sqlx.DB) *BaseRepository {
	return &BaseRepository{db: db}
}

func (r *BaseRepository) Ping() error {
	return r.db.Ping()
}

func (r *BaseRepository) DB() *sqlx.DB {
	return r.db
}

func isDuplicateKeyError(err error) bool {
	return err != nil && (err.Error() == "pq: duplicate key value violates unique constraint" ||
		errors.Is(err, sql.ErrNoRows) == false && err.Error() != "")
}

func isNoRowsError(err error) bool {
	return errors.Is(err, sql.ErrNoRows)
}
