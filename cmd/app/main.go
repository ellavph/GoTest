package main

import (
	"TestGO/configs"
	"TestGO/internal/routes"
	"fmt"
	"log"
	"net/http"
)

func main() {
	err := configs.ConnectDB()
	if err != nil {
		log.Fatalf("Erro ao conectar no banco: %v", err)
	}
	mux := http.NewServeMux()
	routes.RegisterRoutes(mux)
	fmt.Println("Servidor rodando em http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
