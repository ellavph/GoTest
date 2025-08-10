package configs

import (
	"context"
	"errors"
	"log"
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
	
	log.Printf("🔍 [DEBUG] DATABASE_URL from env: %s", dbURL)
	
	pool, err := pgxpool.New(ctx, dbURL)
	if err != nil {
		log.Printf("❌ [ERROR] Failed to create connection pool: %v", err)
		return nil, err
	}
	
	log.Println("✅ [DEBUG] Connection pool created successfully")
	
	// Manter compatibilidade com código existente
	Db = pool
	
	return pool, nil
}
