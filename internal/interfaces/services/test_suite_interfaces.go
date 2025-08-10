package services

import (
	"context"
	"time"

	"TestGO/internal/domain/entities"

	"github.com/google/uuid"
)

// TestSuiteReader define operações de leitura de suítes de teste
type TestSuiteReader interface {
	GetByID(ctx context.Context, id uuid.UUID) (*entities.TestSuite, error)
	GetByCompanyID(ctx context.Context, companyID uuid.UUID) ([]*entities.TestSuite, error)
	List(ctx context.Context, req *ListTestSuitesRequest) (*ListTestSuitesResponse, error)
}

// TestSuiteWriter define operações de escrita de suítes de teste
type TestSuiteWriter interface {
	Create(ctx context.Context, req *CreateTestSuiteRequest) (*entities.TestSuite, error)
	Update(ctx context.Context, id uuid.UUID, req *UpdateTestSuiteRequest) (*entities.TestSuite, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

// TestSuiteService combina todas as operações de suíte de teste
type TestSuiteService interface {
	TestSuiteReader
	TestSuiteWriter
}

// CreateTestSuiteRequest representa uma solicitação de criação de suíte de teste
type CreateTestSuiteRequest struct {
	CompanyID      uuid.UUID `json:"company_id" validate:"required"`
	Name           string    `json:"name" validate:"required,min=3,max=100"`
	Method         string    `json:"method" validate:"required,oneof=GET POST PUT DELETE PATCH"`
	URL            string    `json:"url" validate:"required,url"`
	Headers        string    `json:"headers" validate:"omitempty"`
	ExpectedStatus int       `json:"expected_status" validate:"required,min=100,max=599"`
	ExpectedBody   string    `json:"expected_body" validate:"omitempty"`
}

// UpdateTestSuiteRequest representa uma solicitação de atualização de suíte de teste
type UpdateTestSuiteRequest struct {
	Name           string `json:"name" validate:"omitempty,min=3,max=100"`
	Method         string `json:"method" validate:"omitempty,oneof=GET POST PUT DELETE PATCH"`
	URL            string `json:"url" validate:"omitempty,url"`
	Headers        string `json:"headers" validate:"omitempty"`
	ExpectedStatus int    `json:"expected_status" validate:"omitempty,min=100,max=599"`
	ExpectedBody   string `json:"expected_body" validate:"omitempty"`
}

// ListTestSuitesRequest representa uma solicitação de listagem de suítes de teste
type ListTestSuitesRequest struct {
	CompanyID uuid.UUID `json:"company_id" validate:"omitempty"`
	Limit     int       `json:"limit" validate:"min=1,max=100"`
	Offset    int       `json:"offset" validate:"min=0"`
}

// ListTestSuitesResponse representa a resposta de listagem de suítes de teste
type ListTestSuitesResponse struct {
	TestSuites []*TestSuiteResponse `json:"test_suites"`
	Total      int64                `json:"total"`
	Limit      int                  `json:"limit"`
	Offset     int                  `json:"offset"`
}

// TestSuiteResponse representa a resposta completa de suíte de teste
type TestSuiteResponse struct {
	ID             uuid.UUID `json:"id"`
	CompanyID      uuid.UUID `json:"company_id"`
	Name           string    `json:"name"`
	Method         string    `json:"method"`
	URL            string    `json:"url"`
	Headers        string    `json:"headers"`
	ExpectedStatus int       `json:"expected_status"`
	ExpectedBody   string    `json:"expected_body"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
