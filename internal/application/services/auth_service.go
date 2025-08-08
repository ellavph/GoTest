package services

import (
	"context"
	"fmt"
	"time"

	"TestGO/internal/domain/entities"
	"TestGO/internal/domain/repositories"
	"TestGO/internal/infrastructure/security"
	"TestGO/internal/interfaces/services"

	"github.com/google/uuid"
)

type authService struct {
	userRepo        repositories.UserRepository
	passwordService *security.PasswordService
	jwtService      *security.JWTService
}

// NewAuthService cria uma nova instância do serviço de autenticação
func NewAuthService(
	userRepo repositories.UserRepository,
	passwordService *security.PasswordService,
	jwtService *security.JWTService,
) services.AuthService {
	return &authService{
		userRepo:        userRepo,
		passwordService: passwordService,
		jwtService:      jwtService,
	}
}

func (s *authService) Login(ctx context.Context, req *services.LoginRequest) (*services.LoginResponse, error) {
	// Buscar usuário por username
	user, err := s.userRepo.GetByUsername(ctx, req.Username)
	if err != nil {
		return nil, fmt.Errorf("invalid credentials")
	}

	// Verificar senha
	if !s.passwordService.CheckPassword(user.Password, req.Password) {
		return nil, fmt.Errorf("invalid credentials")
	}

	// Gerar token
	token, err := s.jwtService.GenerateAccessToken(user.ID, user.Username, user.Email, user.CompanyID)
	if err != nil {
		return nil, fmt.Errorf("failed to generate token: %w", err)
	}

	return &services.LoginResponse{
		Token:     token,
		ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
	}, nil
}

func (s *authService) Register(ctx context.Context, req *services.RegisterRequest) (*services.RegisterResponse, error) {
	// Verificar se username já existe
	exists, err := s.userRepo.ExistsByUsername(ctx, req.Username)
	if err != nil {
		return nil, fmt.Errorf("failed to check username existence: %w", err)
	}
	if exists {
		return nil, fmt.Errorf("username already exists")
	}

	// Verificar se email já existe
	exists, err = s.userRepo.ExistsByEmail(ctx, req.Email)
	if err != nil {
		return nil, fmt.Errorf("failed to check email existence: %w", err)
	}
	if exists {
		return nil, fmt.Errorf("email already exists")
	}

	// Hash da senha
	hashedPassword, err := s.passwordService.HashPassword(req.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	// Criar usuário
	var user *entities.User
	if req.Name != "" {
		user = entities.NewUserWithName(req.Username, req.Email, hashedPassword, req.Name)
	} else {
		user = entities.NewUser(req.Username, req.Email, hashedPassword)
	}

	// Salvar no banco
	createdUser, err := s.userRepo.Create(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return &services.RegisterResponse{
		ID:       int64(createdUser.ID.ID()),
		Username: createdUser.Username,
		Email:    createdUser.Email,
	}, nil
}

func (s *authService) ValidateJWT(tokenString string) (*uuid.UUID, error) {
	claims, err := s.jwtService.ValidateToken(tokenString)
	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	// Verificar se o token não expirou
	if time.Now().Unix() > claims.ExpiresAt {
		return nil, fmt.Errorf("token expired")
	}

	return &claims.UserID, nil
}

// Removendo métodos que não fazem parte da interface Clean Architecture
// RefreshToken e Logout serão implementados em casos de uso específicos se necessário
