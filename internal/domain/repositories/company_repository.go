package repositories

import (
	"context"

	"github.com/google/uuid"
	"TestGO/internal/domain/entities"
)

// CompanyRepository define as operações de persistência para empresas
type CompanyRepository interface {
	Create(ctx context.Context, company *entities.Company) error
	GetByID(ctx context.Context, id uuid.UUID) (*entities.Company, error)
	GetByName(ctx context.Context, name string) (*entities.Company, error)
	GetByEmail(ctx context.Context, email string) (*entities.Company, error)
	Update(ctx context.Context, company *entities.Company) error
	Delete(ctx context.Context, id uuid.UUID) error
	List(ctx context.Context, limit, offset int) ([]*entities.Company, error)
	ExistsByName(ctx context.Context, name string) (bool, error)
	ExistsByEmail(ctx context.Context, email string) (bool, error)
}