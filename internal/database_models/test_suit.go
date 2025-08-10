package database_models

import (
	"time"

	"github.com/google/uuid"
)

type TestSuite struct {
	ID             uuid.UUID `gorm:"primaryKey;type:uuid" json:"id"`
	CompanyID      uuid.UUID `gorm:"type:uuid" json:"company_id"`
	Company        *Company  `gorm:"foreignKey:CompanyID" json:"company,omitempty"`
	Name           string    `json:"name"`
	Method         string    `json:"method"`
	URL            string    `json:"url"`
	Headers        string    `json:"headers"`
	ExpectedStatus int       `json:"expected_status"`
	ExpectedBody   string    `json:"expected_body"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
