package database

import (
	"TestGO/internal/domain/entities"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type testSuitRepository struct {
	db *pgxpool.Pool
}

func NewTestSuiteRepository(db *pgxpool.Pool) *testSuitRepository {
	return &testSuitRepository{db: db}
}

func (r *testSuitRepository) Create(ctx context.Context, testSuite *entities.TestSuite) (*entities.TestSuite, error) {
	query := `INSERT INTO test_suites (id, company_id, name, method, url, headers, expected_status, expected_body, created_at, updated_at)
			VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
			RETURNING id, company_id, name, method, url, headers, expected_status, expected_body, created_at, updated_at
	`

	createdTestSuite := &entities.TestSuite{}
	err := r.db.QueryRow(ctx, query,
		testSuite.ID,
		testSuite.CompanyID,
		testSuite.Name,
		testSuite.Method,
		testSuite.URL,
		testSuite.Headers,
		testSuite.ExpectedStatus,
		testSuite.ExpectedBody,
		testSuite.CreatedAt,
		testSuite.UpdatedAt,
	).Scan(
		&createdTestSuite.ID,
		&createdTestSuite.CompanyID,
		&createdTestSuite.Name,
		&createdTestSuite.Method,
		&createdTestSuite.URL,
		&createdTestSuite.Headers,
		&createdTestSuite.ExpectedStatus,
		&createdTestSuite.ExpectedBody,
		&createdTestSuite.CreatedAt,
		&createdTestSuite.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return createdTestSuite, nil
}
