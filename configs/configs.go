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

// LoadEnv carrega as variáveis de ambiente do arquivo .env
func LoadEnv() {
	// Forçar o carregamento do arquivo .env, sobrescrevendo variáveis existentes
	err := godotenv.Overload()
	if err != nil {
		log.Println("Aviso: Não foi possível carregar o arquivo .env:", err)
	} else {
		log.Println("Variáveis de ambiente carregadas do arquivo .env")
	}
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
