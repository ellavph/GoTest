package services

import (
	"context"

	"github.com/google/uuid"
	"TestGO/internal/domain/entities"
)

// CompanyReader define operações de leitura de empresas
type CompanyReader interface {
	GetByID(ctx context.Context, id uuid.UUID) (*entities.Company, error)
	List(ctx context.Context, req *ListCompaniesRequest) (*ListCompaniesResponse, error)
}

// CompanyWriter define operações de escrita de empresas
type CompanyWriter interface {
	Create(ctx context.Context, req *CreateCompanyRequest) (*entities.Company, error)
	Update(ctx context.Context, id uuid.UUID, req *UpdateCompanyRequest) (*entities.Company, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

// CompanyService combina todas as operações de empresa
type CompanyService interface {
	CompanyReader
	CompanyWriter
}

// CreateCompanyRequest representa uma solicitação de criação de empresa
type CreateCompanyRequest struct {
	Name        string `json:"name" validate:"required,min=2,max=100"`
	Email       string `json:"email" validate:"required,email"`
	Phone       string `json:"phone,omitempty" validate:"omitempty,min=10,max=20"`
	Address     string `json:"address,omitempty" validate:"omitempty,max=200"`
	Description string `json:"description,omitempty" validate:"omitempty,max=500"`
}

// UpdateCompanyRequest representa uma solicitação de atualização de empresa
type UpdateCompanyRequest struct {
	Name        string `json:"name,omitempty" validate:"omitempty,min=2,max=100"`
	Email       string `json:"email,omitempty" validate:"omitempty,email"`
	Phone       string `json:"phone,omitempty" validate:"omitempty,min=10,max=20"`
	Address     string `json:"address,omitempty" validate:"omitempty,max=200"`
	Description string `json:"description,omitempty" validate:"omitempty,max=500"`
}

// ListCompaniesRequest representa uma solicitação de listagem de empresas
type ListCompaniesRequest struct {
	Limit  int `json:"limit" validate:"min=1,max=100"`
	Offset int `json:"offset" validate:"min=0"`
}

// ListCompaniesResponse representa a resposta de listagem de empresas
type ListCompaniesResponse struct {
	Companies []*entities.Company `json:"companies"`
	Total     int64               `json:"total"`
	Limit     int                 `json:"limit"`
	Offset    int                 `json:"offset"`
}