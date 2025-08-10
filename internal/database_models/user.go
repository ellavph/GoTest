package database_models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"primaryKey;type:uuid" json:"id"`
	Username  string    `gorm:"unique" json:"username"`
	Email     string    `gorm:"unique" json:"email"`
	Password  string    `json:"-"`
	Name      string    `json:"name"`
	CompanyID uuid.UUID `gorm:"type:uuid" json:"company_id"`
	Company   *Company  `gorm:"foreignKey:CompanyID" json:"company,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
