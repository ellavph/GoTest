package services

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"TestGO/internal/domain/entities"
	"TestGO/internal/domain/repositories"
	"TestGO/internal/interfaces/services"
	"TestGO/internal/infrastructure/security"
)

type userService struct {
	userRepo        repositories.UserRepository
	passwordService *security.PasswordService
}

// NewUserService cria uma nova instância do serviço de usuário
func NewUserService(
	userRepo repositories.UserRepository,
	passwordService *security.PasswordService,
) services.UserService {
	return &userService{
		userRepo:        userRepo,
		passwordService: passwordService,
	}
}

func (s *userService) GetByID(ctx context.Context, id uuid.UUID) (*entities.User, error) {
	user, err := s.userRepo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}
	return user, nil
}

func (s *userService) GetByUsername(ctx context.Context, username string) (*entities.User, error) {
	user, err := s.userRepo.GetByUsername(ctx, username)
	if err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}
	return user, nil
}

func (s *userService) Update(ctx context.Context, id uuid.UUID, req *services.UpdateUserRequest) (*entities.User, error) {
	// Buscar usuário existente
	user, err := s.userRepo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}

	// Verificar se o novo username já existe (se fornecido)
	if req.Username != "" && req.Username != user.Username {
		exists, err := s.userRepo.ExistsByUsername(ctx, req.Username)
		if err != nil {
			return nil, fmt.Errorf("failed to check username: %w", err)
		}
		if exists {
			return nil, fmt.Errorf("username already exists")
		}
	}

	// Verificar se o novo email já existe (se fornecido)
	if req.Email != "" && req.Email != user.Email {
		exists, err := s.userRepo.ExistsByEmail(ctx, req.Email)
		if err != nil {
			return nil, fmt.Errorf("failed to check email: %w", err)
		}
		if exists {
			return nil, fmt.Errorf("email already exists")
		}
	}

	// Atualizar campos
	if req.Username != "" {
		user.Username = req.Username
	}
	if req.Email != "" {
		user.Email = req.Email
	}
	if req.Name != "" {
		user.Name = req.Name
	}

	// Salvar no banco
	err = s.userRepo.Update(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("failed to update user: %w", err)
	}

	return user, nil
}

func (s *userService) Delete(ctx context.Context, id uuid.UUID) error {
	// Verificar se o usuário existe
	_, err := s.userRepo.GetByID(ctx, id)
	if err != nil {
		return fmt.Errorf("user not found: %w", err)
	}

	// Deletar usuário
	err = s.userRepo.Delete(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}

	return nil
}

func (s *userService) List(ctx context.Context, req *services.ListUsersRequest) (*services.ListUsersResponse, error) {
	// Definir valores padrão
	if req.Limit < 1 {
		req.Limit = 10
	}
	if req.Limit > 100 {
		req.Limit = 100
	}
	if req.Offset < 0 {
		req.Offset = 0
	}

	// Buscar usuários
	users, err := s.userRepo.List(ctx, req.Limit, req.Offset)
	if err != nil {
		return nil, fmt.Errorf("failed to list users: %w", err)
	}

	// Converter entities.User para services.UserResponse
	userResponses := make([]*services.UserResponse, len(users))
	for i, user := range users {
		userResponses[i] = &services.UserResponse{
			ID:        user.ID,
			Username:  user.Username,
			Email:     user.Email,
			Name:      user.Name,
			CompanyID: user.CompanyID,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		}
	}

	// Para simplicidade, vamos assumir um total fixo
	// Em uma implementação real, você faria uma query COUNT
	total := int64(len(users))

	return &services.ListUsersResponse{
		Users:  userResponses,
		Total:  total,
		Limit:  req.Limit,
		Offset: req.Offset,
	}, nil
}

func (s *userService) ChangePassword(ctx context.Context, id uuid.UUID, req *services.ChangePasswordRequest) error {
	// Buscar usuário
	user, err := s.userRepo.GetByID(ctx, id)
	if err != nil {
		return fmt.Errorf("user not found: %w", err)
	}

	// Verificar senha atual
	if !s.passwordService.CheckPassword(user.Password, req.CurrentPassword) {
		return fmt.Errorf("current password is incorrect")
	}

	// Hash da nova senha
	hashedPassword, err := s.passwordService.HashPassword(req.NewPassword)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	// Atualizar senha
	user.UpdatePassword(hashedPassword)

	// Salvar no banco
	err = s.userRepo.Update(ctx, user)
	if err != nil {
		return fmt.Errorf("failed to update password: %w", err)
	}

	return nil
}