package entities

import (
	"time"

	"github.com/google/uuid"
)

// TestSuite representa uma suíte de testes de API
type TestSuite struct {
	ID             uuid.UUID `json:"id" db:"id" example:"550e8400-e29b-41d4-a716-446655440000"`
	CompanyID      uuid.UUID `json:"company_id" db:"company_id" example:"550e8400-e29b-41d4-a716-446655440001"`
	Name           string    `json:"name" db:"name" example:"API Login Test"`
	Method         string    `json:"method" db:"method" example:"POST" enums:"GET,POST,PUT,DELETE,PATCH"`
	URL            string    `json:"url" db:"url" example:"https://api.example.com/login"`
	Headers        string    `json:"headers" db:"headers" example:"Content-Type: application/json"`
	ExpectedStatus int       `json:"expected_status" db:"expected_status" example:"200"`
	ExpectedBody   string    `json:"expected_body" db:"expected_body" example:"{\"success\": true}"`
	CreatedAt      time.Time `json:"created_at" db:"created_at" example:"2023-01-01T00:00:00Z"`
	UpdatedAt      time.Time `json:"updated_at" db:"updated_at" example:"2023-01-01T00:00:00Z"`
}

// NewTestSuite cria uma nova instância de TestSuite
func NewTestSuite(companyID uuid.UUID, name, method, url, headers string, expectedStatus int, expectedBody string) *TestSuite {
	return &TestSuite{
		ID:             uuid.New(),
		CompanyID:      companyID,
		Name:           name,
		Method:         method,
		URL:            url,
		Headers:        headers,
		ExpectedStatus: expectedStatus,
		ExpectedBody:   expectedBody,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}
}

// UpdateTestSuite atualiza os dados da suíte de teste
func (ts *TestSuite) UpdateTestSuite(name, method, url, headers string, expectedStatus int, expectedBody string) {
	if name != "" {
		ts.Name = name
	}
	if method != "" {
		ts.Method = method
	}
	if url != "" {
		ts.URL = url
	}
	if headers != "" {
		ts.Headers = headers
	}
	if expectedStatus > 0 {
		ts.ExpectedStatus = expectedStatus
	}
	if expectedBody != "" {
		ts.ExpectedBody = expectedBody
	}
	ts.UpdatedAt = time.Now()
}
