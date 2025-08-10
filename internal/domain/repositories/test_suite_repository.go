package repositories

import (
	"context"

	"TestGO/internal/domain/entities"

	"github.com/google/uuid"
)

type TestSuiteRepository interface {
	Create(ctx context.Context, testSuite *entities.TestSuite) (*entities.TestSuite, error)
	GetByID(ctx context.Context, id uuid.UUID) (*entities.TestSuite, error)
	GetByCompanyID(ctx context.Context, companyID uuid.UUID) ([]*entities.TestSuite, error)
	Update(ctx context.Context, testSuite *entities.TestSuite) error
	Delete(ctx context.Context, id uuid.UUID) error
	List(ctx context.Context, limit, offset int) ([]*entities.TestSuite, error)
}
