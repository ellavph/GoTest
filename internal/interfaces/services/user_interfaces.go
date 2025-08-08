package services

import (
	"context"
	"time"

	"github.com/google/uuid"
	"TestGO/internal/domain/entities"
)

// UserReader define operações de leitura de usuários
type UserReader interface {
	GetByID(ctx context.Context, id uuid.UUID) (*entities.User, error)
	GetByUsername(ctx context.Context, username string) (*entities.User, error)
	List(ctx context.Context, req *ListUsersRequest) (*ListUsersResponse, error)
}

// UserWriter define operações de escrita de usuários
type UserWriter interface {
	Update(ctx context.Context, id uuid.UUID, req *UpdateUserRequest) (*entities.User, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

// PasswordManager define operações de gerenciamento de senha
type PasswordManager interface {
	ChangePassword(ctx context.Context, id uuid.UUID, req *ChangePasswordRequest) error
}

// UserService combina todas as operações de usuário
type UserService interface {
	UserReader
	UserWriter
	PasswordManager
}

// UpdateUserRequest representa uma solicitação de atualização de usuário
type UpdateUserRequest struct {
	Username string `json:"username" validate:"omitempty,min=3,max=50"`
	Email    string `json:"email" validate:"omitempty,email"`
	Password string `json:"password" validate:"omitempty,min=6"`
}

// ChangePasswordRequest representa uma solicitação de mudança de senha
type ChangePasswordRequest struct {
	CurrentPassword string `json:"current_password" validate:"required"`
	NewPassword     string `json:"new_password" validate:"required,min=6"`
}

// ListUsersRequest representa uma solicitação de listagem de usuários
type ListUsersRequest struct {
	Limit  int `json:"limit" validate:"min=1,max=100"`
	Offset int `json:"offset" validate:"min=0"`
}

// ListUsersResponse representa a resposta de listagem de usuários
type ListUsersResponse struct {
	Users  []*entities.User `json:"users"`
	Total  int64            `json:"total"`
	Limit  int              `json:"limit"`
	Offset int              `json:"offset"`
}

// UserResponse representa a resposta completa de usuário
type UserResponse struct {
	ID        int64     `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CompanyID *int64    `json:"company_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}