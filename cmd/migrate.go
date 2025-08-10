package main

import (
	"log"
	"os"
	"os/exec"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Erro ao carregar .env: %v", err)
	}

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL não definido no .env")
	}

	// Usar goose CLI diretamente que funciona melhor com Neon
	cmd := exec.Command("goose", "-dir", "migrations", "postgres", dbURL, "up")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	
	if err := cmd.Run(); err != nil {
		log.Fatalf("Erro ao aplicar migrations: %v", err)
	}

	log.Println("Migrations aplicadas com sucesso!")
}
