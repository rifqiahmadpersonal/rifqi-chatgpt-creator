package models

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

const (
	AccountStatusActive     = "active"
	AccountStatusInactive   = "inactive"
	AccountStatusSuspended  = "suspended"
)

type Account struct {
	ID         string     `json:"id" db:"id" validate:"required,uuid"`
	Email      string     `json:"email" db:"email" validate:"required,email,max=255"`
	Password   string     `json:"-" db:"password" validate:"required,min=12,max=255"`
	Status     string     `json:"status" db:"status" validate:"required,oneof=active inactive suspended"`
	BatchJobID *string    `json:"batch_job_id,omitempty" db:"batch_job_id" validate:"omitempty,uuid"`
	CreatedAt  time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at" db:"updated_at"`
}

func (a *Account) Validate() error {
	validate := validator.New()
	return validate.Struct(a)
}

func (a *Account) GenerateID() {
	if a.ID == "" {
		a.ID = uuid.New().String()
	}
}

type AccountFilter struct {
	Status     string
	BatchJobID string
	Email      string
	Limit      int
	Offset     int
}
