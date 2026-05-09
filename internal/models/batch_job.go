package models

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

const (
	BatchJobStatusPending   = "pending"
	BatchJobStatusRunning   = "running"
	BatchJobStatusCompleted = "completed"
	BatchJobStatusCancelled = "cancelled"
	BatchJobStatusFailed    = "failed"
)

type BatchJob struct {
	ID              string     `json:"id" db:"id" validate:"required,uuid"`
	TargetCount     int        `json:"target_count" db:"target_count" validate:"required,min=1"`
	SuccessCount    int        `json:"success_count" db:"success_count" validate:"min=0"`
	FailureCount    int        `json:"failure_count" db:"failure_count" validate:"min=0"`
	Status          string     `json:"status" db:"status" validate:"required,oneof=pending running completed cancelled failed"`
	MaxWorkers      int        `json:"max_workers" db:"max_workers" validate:"required,min=1,max=20"`
	DefaultPassword string     `json:"default_password,omitempty" db:"default_password"`
	Proxy           string     `json:"proxy,omitempty" db:"proxy"`
	CreatedAt       time.Time  `json:"created_at" db:"created_at"`
	CompletedAt     *time.Time `json:"completed_at,omitempty" db:"completed_at"`
}

func (b *BatchJob) Validate() error {
	validate := validator.New()
	return validate.Struct(b)
}

func (b *BatchJob) GenerateID() {
	if b.ID == "" {
		b.ID = uuid.New().String()
	}
}
