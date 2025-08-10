package database_models

import (
	"time"

	"github.com/google/uuid"
)

type Company struct {
	ID        uuid.UUID `gorm:"primaryKey;type:uuid" json:"id"`
	Name      string    `json:"name"`
	Location  string    `json:"location"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
