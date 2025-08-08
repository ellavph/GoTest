package configs

import (
	"context"
	"errors"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

var Db *pgxpool.Pool

// LoadEnv carrega as variáveis de ambiente
func LoadEnv() error {
	return godotenv.Load()
}

// ConnectDB conecta ao banco de dados e retorna o pool de conexões
func ConnectDB(ctx context.Context) (*pgxpool.Pool, error) {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		return nil, errors.New("DATABASE_URL não definida")
	}
	
	pool, err := pgxpool.New(ctx, dbURL)
	if err != nil {
		return nil, err
	}
	
	// Manter compatibilidade com código existente
	Db = pool
	
	return pool, nil
}
