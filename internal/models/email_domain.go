package models

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

const (
	DomainSourceGenerator = "generator"
	DomainSourceCustom    = "custom"

	DomainHealthStatusHealthy   = "healthy"
	DomainHealthStatusUnhealthy = "unhealthy"
	DomainHealthStatusUnknown   = "unknown"
)

type EmailDomain struct {
	ID           string     `json:"id" db:"id" validate:"required,uuid"`
	Domain       string     `json:"domain" db:"domain" validate:"required,hostname,max=255"`
	Priority     int        `json:"priority" db:"priority" validate:"min=1,max=100"`
	IsActive     bool       `json:"is_active" db:"is_active"`
	Source       string     `json:"source" db:"source" validate:"required,oneof=generator custom"`
	LastChecked  *time.Time `json:"last_checked,omitempty" db:"last_checked"`
	HealthStatus string     `json:"health_status" db:"health_status" validate:"oneof=healthy unhealthy unknown"`
	CreatedAt    time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at" db:"updated_at"`
}

func (e *EmailDomain) Validate() error {
	validate := validator.New()
	return validate.Struct(e)
}

func (e *EmailDomain) GenerateID() {
	if e.ID == "" {
		e.ID = uuid.New().String()
	}
}
