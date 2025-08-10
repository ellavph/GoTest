package database_models

import (
	"time"

	"github.com/google/uuid"
)

type TestRun struct {
	ID          uuid.UUID `gorm:"primaryKey;type:uuid" json:"id"`
	CompanyID   uuid.UUID `gorm:"type:uuid" json:"company_id"`
	Company     *Company  `gorm:"foreignKey:CompanyID" json:"company,omitempty"`
	StartedAt   time.Time `json:"started_at"`
	FinishedAt  time.Time `json:"finished_at"`
	Status      string    `json:"status"`
	TotalTests  int       `json:"total_tests"`
	PassedTests int       `json:"passed_tests"`
	FailedTests int       `json:"failed_tests"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
