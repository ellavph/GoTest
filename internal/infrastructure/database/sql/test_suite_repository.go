package sql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"TestGO/internal/domain/entities"
	"TestGO/internal/domain/repositories"
)

type testSuiteRepository struct {
	db *pgxpool.Pool
}

// NewTestSuiteRepository cria uma nova instância do repositório de test suite
func NewTestSuiteRepository(db *pgxpool.Pool) repositories.TestSuiteRepository {
	return &testSuiteRepository{db: db}
}

func (r *testSuiteRepository) Create(ctx context.Context, testSuite *entities.TestSuite) (*entities.TestSuite, error) {
	query := `
		INSERT INTO test_suites (id, company_id, name, method, url, headers, expected_status, expected_body, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, NOW(), NOW())
		RETURNING id, company_id, name, method, url, headers, expected_status, expected_body, created_at, updated_at`

	row := r.db.QueryRow(ctx, query,
		testSuite.ID,
		testSuite.CompanyID,
		testSuite.Name,
		testSuite.Method,
		testSuite.URL,
		testSuite.Headers,
		testSuite.ExpectedStatus,
		testSuite.ExpectedBody,
	)

	var created entities.TestSuite
	err := row.Scan(
		&created.ID,
		&created.CompanyID,
		&created.Name,
		&created.Method,
		&created.URL,
		&created.Headers,
		&created.ExpectedStatus,
		&created.ExpectedBody,
		&created.CreatedAt,
		&created.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create test suite: %w", err)
	}

	return &created, nil
}

func (r *testSuiteRepository) GetByID(ctx context.Context, id uuid.UUID) (*entities.TestSuite, error) {
	query := `
		SELECT id, company_id, name, method, url, headers, expected_status, expected_body, created_at, updated_at
		FROM test_suites
		WHERE id = $1`

	row := r.db.QueryRow(ctx, query, id)

	var testSuite entities.TestSuite
	err := row.Scan(
		&testSuite.ID,
		&testSuite.CompanyID,
		&testSuite.Name,
		&testSuite.Method,
		&testSuite.URL,
		&testSuite.Headers,
		&testSuite.ExpectedStatus,
		&testSuite.ExpectedBody,
		&testSuite.CreatedAt,
		&testSuite.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("test suite not found")
		}
		return nil, fmt.Errorf("failed to get test suite: %w", err)
	}

	return &testSuite, nil
}

func (r *testSuiteRepository) GetByCompanyID(ctx context.Context, companyID uuid.UUID) ([]*entities.TestSuite, error) {
	query := `
		SELECT id, company_id, name, method, url, headers, expected_status, expected_body, created_at, updated_at
		FROM test_suites
		WHERE company_id = $1
		ORDER BY created_at DESC`

	rows, err := r.db.Query(ctx, query, companyID)
	if err != nil {
		return nil, fmt.Errorf("failed to get test suites by company: %w", err)
	}
	defer rows.Close()

	var testSuites []*entities.TestSuite
	for rows.Next() {
		var testSuite entities.TestSuite
		err := rows.Scan(
			&testSuite.ID,
			&testSuite.CompanyID,
			&testSuite.Name,
			&testSuite.Method,
			&testSuite.URL,
			&testSuite.Headers,
			&testSuite.ExpectedStatus,
			&testSuite.ExpectedBody,
			&testSuite.CreatedAt,
			&testSuite.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan test suite: %w", err)
		}
		testSuites = append(testSuites, &testSuite)
	}

	return testSuites, nil
}

func (r *testSuiteRepository) Update(ctx context.Context, testSuite *entities.TestSuite) error {
	query := `
		UPDATE test_suites
		SET name = $2, method = $3, url = $4, headers = $5, expected_status = $6, expected_body = $7, updated_at = NOW()
		WHERE id = $1`

	result, err := r.db.Exec(ctx, query,
		testSuite.ID,
		testSuite.Name,
		testSuite.Method,
		testSuite.URL,
		testSuite.Headers,
		testSuite.ExpectedStatus,
		testSuite.ExpectedBody,
	)
	if err != nil {
		return fmt.Errorf("failed to update test suite: %w", err)
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("test suite not found")
	}

	return nil
}

func (r *testSuiteRepository) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM test_suites WHERE id = $1`

	result, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete test suite: %w", err)
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("test suite not found")
	}

	return nil
}

func (r *testSuiteRepository) List(ctx context.Context, limit, offset int) ([]*entities.TestSuite, error) {
	query := `
		SELECT id, company_id, name, method, url, headers, expected_status, expected_body, created_at, updated_at
		FROM test_suites
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2`

	rows, err := r.db.Query(ctx, query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to list test suites: %w", err)
	}
	defer rows.Close()

	var testSuites []*entities.TestSuite
	for rows.Next() {
		var testSuite entities.TestSuite
		err := rows.Scan(
			&testSuite.ID,
			&testSuite.CompanyID,
			&testSuite.Name,
			&testSuite.Method,
			&testSuite.URL,
			&testSuite.Headers,
			&testSuite.ExpectedStatus,
			&testSuite.ExpectedBody,
			&testSuite.CreatedAt,
			&testSuite.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan test suite: %w", err)
		}
		testSuites = append(testSuites, &testSuite)
	}

	return testSuites, nil
}