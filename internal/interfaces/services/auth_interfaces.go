package services

import (
	"context"

	"github.com/google/uuid"
)

// Authenticator define operações de autenticação
type Authenticator interface {
	Login(ctx context.Context, req *LoginRequest) (*LoginResponse, error)
}

// TokenValidator define operações de validação de token
type TokenValidator interface {
	ValidateJWT(tokenString string) (*uuid.UUID, error)
}

// UserRegistrar define operações de registro de usuário
type UserRegistrar interface {
	Register(ctx context.Context, req *RegisterRequest) (*RegisterResponse, error)
}

// AuthService combina todas as operações de autenticação
type AuthService interface {
	Authenticator
	TokenValidator
	UserRegistrar
}

// LoginRequest representa uma solicitação de login
type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// LoginResponse representa a resposta de login
type LoginResponse struct {
	Token     string       `json:"token"`
	User      UserResponse `json:"user"`
	ExpiresAt int64        `json:"expires_at"`
}

// RegisterRequest representa uma solicitação de registro
type RegisterRequest struct {
	Username string `json:"username" validate:"required,min=3,max=50"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
	Name     string `json:"name" validate:"omitempty,min=2,max=100"`
}

// RegisterResponse representa a resposta de registro
type RegisterResponse struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
}