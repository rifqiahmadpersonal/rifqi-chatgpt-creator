package models

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

const (
	ConfigKeyDefaultProxy          = "default_proxy"
	ConfigKeyDefaultPassword       = "default_password"
	ConfigKeyDefaultDomain         = "default_domain"
	ConfigKeyWorkerPoolSize        = "worker_pool_size"
	ConfigKeyMaxRetries            = "max_retries"
	ConfigKeyRegistrationTimeout   = "registration_timeout"
)

type Configuration struct {
	ID        string    `json:"id" db:"id" validate:"required,uuid"`
	Key       string    `json:"key" db:"key" validate:"required,max=100"`
	Value     string    `json:"value" db:"value"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

func (c *Configuration) Validate() error {
	validate := validator.New()
	return validate.Struct(c)
}

func (c *Configuration) GenerateID() {
	if c.ID == "" {
		c.ID = uuid.New().String()
	}
}
