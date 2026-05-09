package models

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type BlacklistedDomain struct {
	ID        string    `json:"id" db:"id" validate:"required,uuid"`
	Domain    string    `json:"domain" db:"domain" validate:"required,hostname,max=255"`
	Reason    string    `json:"reason" db:"reason" validate:"max=500"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

func (b *BlacklistedDomain) Validate() error {
	validate := validator.New()
	return validate.Struct(b)
}

func (b *BlacklistedDomain) GenerateID() {
	if b.ID == "" {
		b.ID = uuid.New().String()
	}
}
