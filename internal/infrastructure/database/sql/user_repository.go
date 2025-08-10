package sql

import (
	"context"
	"database/sql"
	"fmt"

	"TestGO/internal/domain/entities"
	"TestGO/internal/domain/repositories"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type userRepository struct {
	db *pgxpool.Pool
}

// NewUserRepository cria uma nova instância do repositório de usuário
func NewUserRepository(db *pgxpool.Pool) repositories.UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(ctx context.Context, user *entities.User) (*entities.User, error) {
	query := `
		INSERT INTO users (id, username, email, password, name, company_id, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id, username, email, password, name, company_id, created_at, updated_at
	`

	createdUser := &entities.User{}
	err := r.db.QueryRow(ctx, query,
		user.ID,
		user.Username,
		user.Email,
		user.Password,
		user.Name,
		user.CompanyID,
		user.CreatedAt,
		user.UpdatedAt,
	).Scan(
		&createdUser.ID,
		&createdUser.Username,
		&createdUser.Email,
		&createdUser.Password,
		&createdUser.Name,
		&createdUser.CompanyID,
		&createdUser.CreatedAt,
		&createdUser.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return createdUser, nil
}

func (r *userRepository) GetByID(ctx context.Context, id uuid.UUID) (*entities.User, error) {
	query := `
		SELECT id, username, email, password, name, company_id, created_at, updated_at
		FROM users
		WHERE id = $1
	`

	user := &entities.User{}
	err := r.db.QueryRow(ctx, query, id).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.Name,
		&user.CompanyID,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, err
	}

	return user, nil
}

func (r *userRepository) GetByUsername(ctx context.Context, username string) (*entities.User, error) {
	query := `
		SELECT id, username, email, password, name, company_id, created_at, updated_at
		FROM users
		WHERE username = $1
	`

	user := &entities.User{}
	err := r.db.QueryRow(ctx, query, username).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.Name,
		&user.CompanyID,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, err
	}

	return user, nil
}

func (r *userRepository) GetByEmail(ctx context.Context, email string) (*entities.User, error) {
	query := `
		SELECT id, username, email, password, name, company_id, created_at, updated_at
		FROM users
		WHERE email = $1
	`

	user := &entities.User{}
	err := r.db.QueryRow(ctx, query, email).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.Name,
		&user.CompanyID,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, err
	}

	return user, nil
}

func (r *userRepository) Update(ctx context.Context, user *entities.User) error {
	query := `
		UPDATE users
		SET username = $2, email = $3, password = $4, name = $5, company_id = $6, updated_at = $7
		WHERE id = $1
	`

	_, err := r.db.Exec(ctx, query,
		user.ID,
		user.Username,
		user.Email,
		user.Password,
		user.Name,
		user.CompanyID,
		user.UpdatedAt,
	)

	return err
}

func (r *userRepository) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := r.db.Exec(ctx, query, id)
	return err
}

func (r *userRepository) List(ctx context.Context, limit, offset int) ([]*entities.User, error) {
	query := `
		SELECT id, username, email, password, name, company_id, created_at, updated_at
		FROM users
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2
	`

	rows, err := r.db.Query(ctx, query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*entities.User
	for rows.Next() {
		user := &entities.User{}
		err := rows.Scan(
			&user.ID,
			&user.Username,
			&user.Email,
			&user.Password,
			&user.Name,
			&user.CompanyID,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (r *userRepository) ExistsByUsername(ctx context.Context, username string) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM users WHERE username = $1)`
	var exists bool
	err := r.db.QueryRow(ctx, query, username).Scan(&exists)
	return exists, err
}

func (r *userRepository) ExistsByEmail(ctx context.Context, email string) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)`
	var exists bool
	err := r.db.QueryRow(ctx, query, email).Scan(&exists)
	return exists, err
}
