package repositories

import (
	"context"
	"errors"
	"ezTest/configs"
	"ezTest/internal/models"
)

func GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User

	query := "SELECT id, email, password FROM users WHERE email = $1"
	row := configs.Db.QueryRow(ctx, query, email)
	err := row.Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		return nil, errors.New("usuário não encontrado")
	}

	return &user, nil
}

func GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	var user models.User

	query := "SELECT id, username, email, password FROM users WHERE username = $1"
	row := configs.Db.QueryRow(ctx, query, username)
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		return nil, errors.New("usuário não encontrado")
	}

	return &user, nil
}

func CreateUser(ctx context.Context, user *models.User) error {
	query := `
        INSERT INTO users (username, email, password)
        VALUES ($1, $2, $3)
        RETURNING id
    `
	err := configs.Db.QueryRow(ctx, query, user.Username, user.Email, user.Password).Scan(&user.ID)
	if err != nil {
		return err
	}
	return nil
}
