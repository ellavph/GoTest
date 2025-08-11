package services

import (
	"context"
	"fmt"

	"TestGO/internal/domain/entities"
	"TestGO/internal/domain/repositories"
	"TestGO/internal/interfaces/services"

	"github.com/google/uuid"
)

type companyService struct {
	companyRepo repositories.CompanyRepository
}

// NewCompanyService cria uma nova instância do serviço de empresa
func NewCompanyService(companyRepo repositories.CompanyRepository) services.CompanyService {
	return &companyService{
		companyRepo: companyRepo,
	}
}

func (s *companyService) Create(ctx context.Context, req *services.CreateCompanyRequest) (*entities.Company, error) {
	// Verificar se o nome já existe
	exists, err := s.companyRepo.ExistsByName(ctx, req.Name)
	if err != nil {
		return nil, fmt.Errorf("failed to check company name: %w", err)
	}
	if exists {
		return nil, fmt.Errorf("company name already exists")
	}

	// Verificar se o email já existe
	exists, err = s.companyRepo.ExistsByEmail(ctx, req.Email)
	if err != nil {
		return nil, fmt.Errorf("failed to check company email: %w", err)
	}
	if exists {
		return nil, fmt.Errorf("company email already exists")
	}

	// Criar empresa
	company := entities.NewCompany(req.Name, req.Email, req.Phone, req.Address)

	// Salvar no banco
	err = s.companyRepo.Create(ctx, company)
	if err != nil {
		return nil, fmt.Errorf("failed to create company: %w", err)
	}

	return company, nil
}

func (s *companyService) GetByID(ctx context.Context, id uuid.UUID) (*entities.Company, error) {
	company, err := s.companyRepo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("company not found: %w", err)
	}
	return company, nil
}

func (s *companyService) Update(ctx context.Context, id uuid.UUID, req *services.UpdateCompanyRequest) (*entities.Company, error) {
	// Buscar empresa existente
	company, err := s.companyRepo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("company not found: %w", err)
	}

	// Verificar se o novo nome já existe (se fornecido)
	if req.Name != "" && req.Name != company.Name {
		exists, err := s.companyRepo.ExistsByName(ctx, req.Name)
		if err != nil {
			return nil, fmt.Errorf("failed to check company name: %w", err)
		}
		if exists {
			return nil, fmt.Errorf("company name already exists")
		}
	}

	// Verificar se o novo email já existe (se fornecido)
	if req.Email != "" && req.Email != company.Email {
		exists, err := s.companyRepo.ExistsByEmail(ctx, req.Email)
		if err != nil {
			return nil, fmt.Errorf("failed to check company email: %w", err)
		}
		if exists {
			return nil, fmt.Errorf("company email already exists")
		}
	}

	// Atualizar campos
	company.UpdateInfo(req.Name, req.Email, req.Phone, req.Address)

	// Salvar no banco
	err = s.companyRepo.Update(ctx, company)
	if err != nil {
		return nil, fmt.Errorf("failed to update company: %w", err)
	}

	return company, nil
}

func (s *companyService) Delete(ctx context.Context, id uuid.UUID) error {
	// Verificar se a empresa existe
	_, err := s.companyRepo.GetByID(ctx, id)
	if err != nil {
		return fmt.Errorf("company not found: %w", err)
	}

	// Deletar empresa
	err = s.companyRepo.Delete(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to delete company: %w", err)
	}

	return nil
}

func (s *companyService) List(ctx context.Context, req *services.ListCompaniesRequest) (*services.ListCompaniesResponse, error) {
	// Definir valores padrão
	if req.Limit < 1 {
		req.Limit = 10
	}
	if req.Limit > 100 {
		req.Limit = 100
	}
	if req.Offset < 0 {
		req.Offset = 0
	}

	// Buscar empresas
	companies, err := s.companyRepo.List(ctx, req.Limit, req.Offset)
	if err != nil {
		return nil, fmt.Errorf("failed to list companies: %w", err)
	}

	// Para simplicidade, vamos assumir um total fixo
	// Em uma implementação real, você faria uma query COUNT
	total := int64(len(companies))

	return &services.ListCompaniesResponse{
		Companies: companies,
		Total:     total,
		Limit:     req.Limit,
		Offset:    req.Offset,
	}, nil
}
