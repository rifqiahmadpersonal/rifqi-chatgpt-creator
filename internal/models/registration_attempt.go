package models

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

const (
	AttemptStatusSuccess    = "success"
	AttemptStatusFailed     = "failed"
	AttemptStatusInProgress = "in_progress"
)

type RegistrationAttempt struct {
	ID           string     `json:"id" db:"id" validate:"required,uuid"`
	Email        string     `json:"email" db:"email" validate:"required,email"`
	Status       string     `json:"status" db:"status" validate:"required,oneof=success failed in_progress"`
	ErrorMessage *string    `json:"error_message,omitempty" db:"error_message"`
	WorkerID     int        `json:"worker_id" db:"worker_id" validate:"min=1"`
	BatchJobID   string     `json:"batch_job_id" db:"batch_job_id" validate:"required,uuid"`
	StartedAt    time.Time  `json:"started_at" db:"started_at"`
	CompletedAt  *time.Time `json:"completed_at,omitempty" db:"completed_at"`
	Duration     *int64     `json:"duration_ms,omitempty" db:"duration_ms"`
}

func (r *RegistrationAttempt) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}

func (r *RegistrationAttempt) GenerateID() {
	if r.ID == "" {
		r.ID = uuid.New().String()
	}
}
