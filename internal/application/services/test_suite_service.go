package services

import (
	"context"

	"TestGO/internal/domain/entities"
	"TestGO/internal/domain/repositories"
	"TestGO/internal/interfaces/services"

	"github.com/google/uuid"
)

type testSuiteService struct {
	testSuiteRepo repositories.TestSuiteRepository
}

func NewTestSuiteService(testSuiteRepo repositories.TestSuiteRepository) services.TestSuiteService {
	return &testSuiteService{
		testSuiteRepo: testSuiteRepo,
	}
}

func (s *testSuiteService) Create(ctx context.Context, req *services.CreateTestSuiteRequest) (*entities.TestSuite, error) {
	return nil, nil
}

func (s *testSuiteService) Update(ctx context.Context, id uuid.UUID, req *services.UpdateTestSuiteRequest) (*entities.TestSuite, error) {
	return nil, nil
}

func (s *testSuiteService) Delete(ctx context.Context, id uuid.UUID) error {
	return nil
}
func (s *testSuiteService) GetByID(ctx context.Context, id uuid.UUID) (*entities.TestSuite, error) {
	return nil, nil
}

func (s *testSuiteService) GetByCompanyID(ctx context.Context, companyID uuid.UUID) ([]*entities.TestSuite, error) {
	return nil, nil
}

func (s *testSuiteService) List(ctx context.Context, req *services.ListTestSuitesRequest) (*services.ListTestSuitesResponse, error) {
	return nil, nil
}
