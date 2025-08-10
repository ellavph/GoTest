package database_models

import (
	"time"

	"github.com/google/uuid"
)

type TestResult struct {
	ID             uuid.UUID  `gorm:"primaryKey;type:uuid" json:"id"`
	TestRunID      uuid.UUID  `gorm:"type:uuid" json:"test_run_id"`
	TestRun        *TestRun   `gorm:"foreignKey:TestRunID" json:"test_run,omitempty"`
	EndpointTestID uuid.UUID  `gorm:"type:uuid" json:"endpoint_test_id"`
	EndpointTest   *TestSuite `gorm:"foreignKey:EndpointTestID" json:"endpoint_test,omitempty"`
	Status         string     `json:"status"`
	ResponseStatus int        `json:"response_status"`
	ResponseBody   string     `json:"response_body"`
	ResponseTimeMS int        `json:"response_time_ms"`
	ErrorMessage   string     `json:"error_message"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`
}
