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

type companyRepository struct {
	db *pgxpool.Pool
}

// NewCompanyRepository cria uma nova instância do repositório de empresa
func NewCompanyRepository(db *pgxpool.Pool) repositories.CompanyRepository {
	return &companyRepository{db: db}
}

func (r *companyRepository) Create(ctx context.Context, company *entities.Company) error {
	query := `
		INSERT INTO companies (id, name, email, phone, address, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`
	
	_, err := r.db.Exec(ctx, query,
		company.ID,
		company.Name,
		company.Email,
		company.Phone,
		company.Address,
		company.CreatedAt,
		company.UpdatedAt,
	)
	
	return err
}

func (r *companyRepository) GetByID(ctx context.Context, id uuid.UUID) (*entities.Company, error) {
	query := `
		SELECT id, name, email, phone, address, created_at, updated_at
		FROM companies
		WHERE id = $1
	`
	
	company := &entities.Company{}
	err := r.db.QueryRow(ctx, query, id).Scan(
		&company.ID,
		&company.Name,
		&company.Email,
		&company.Phone,
		&company.Address,
		&company.CreatedAt,
		&company.UpdatedAt,
	)
	
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("company not found")
		}
		return nil, err
	}
	
	return company, nil
}

func (r *companyRepository) GetByName(ctx context.Context, name string) (*entities.Company, error) {
	query := `
		SELECT id, name, email, phone, address, created_at, updated_at
		FROM companies
		WHERE name = $1
	`
	
	company := &entities.Company{}
	err := r.db.QueryRow(ctx, query, name).Scan(
		&company.ID,
		&company.Name,
		&company.Email,
		&company.Phone,
		&company.Address,
		&company.CreatedAt,
		&company.UpdatedAt,
	)
	
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("company not found")
		}
		return nil, err
	}
	
	return company, nil
}

func (r *companyRepository) GetByEmail(ctx context.Context, email string) (*entities.Company, error) {
	query := `
		SELECT id, name, email, phone, address, created_at, updated_at
		FROM companies
		WHERE email = $1
	`
	
	company := &entities.Company{}
	err := r.db.QueryRow(ctx, query, email).Scan(
		&company.ID,
		&company.Name,
		&company.Email,
		&company.Phone,
		&company.Address,
		&company.CreatedAt,
		&company.UpdatedAt,
	)
	
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("company not found")
		}
		return nil, err
	}
	
	return company, nil
}

func (r *companyRepository) Update(ctx context.Context, company *entities.Company) error {
	query := `
		UPDATE companies
		SET name = $2, email = $3, phone = $4, address = $5, updated_at = $6
		WHERE id = $1
	`
	
	_, err := r.db.Exec(ctx, query,
		company.ID,
		company.Name,
		company.Email,
		company.Phone,
		company.Address,
		company.UpdatedAt,
	)
	
	return err
}

func (r *companyRepository) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM companies WHERE id = $1`
	_, err := r.db.Exec(ctx, query, id)
	return err
}

func (r *companyRepository) List(ctx context.Context, limit, offset int) ([]*entities.Company, error) {
	query := `
		SELECT id, name, email, phone, address, created_at, updated_at
		FROM companies
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2
	`
	
	rows, err := r.db.Query(ctx, query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	var companies []*entities.Company
	for rows.Next() {
		company := &entities.Company{}
		err := rows.Scan(
			&company.ID,
			&company.Name,
			&company.Email,
			&company.Phone,
			&company.Address,
			&company.CreatedAt,
			&company.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		companies = append(companies, company)
	}
	
	return companies, nil
}

func (r *companyRepository) ExistsByName(ctx context.Context, name string) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM companies WHERE name = $1)`
	var exists bool
	err := r.db.QueryRow(ctx, query, name).Scan(&exists)
	return exists, err
}

func (r *companyRepository) ExistsByEmail(ctx context.Context, email string) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM companies WHERE email = $1)`
	var exists bool
	err := r.db.QueryRow(ctx, query, email).Scan(&exists)
	return exists, err
}