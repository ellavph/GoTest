package configs

import (
	"context"
	"errors"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv" // Import the godotenv library
)

var Db *pgxpool.Pool

func ConnectDB() error {
	err := godotenv.Load()
	if err != nil {
		return errors.New("Erro ao carregar .env")
	}

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		return errors.New("DATABASE_URL não definida")
	}
	pool, err := pgxpool.New(context.Background(), dbURL)
	if err != nil {
		return err
	}
	Db = pool
	return nil
}
