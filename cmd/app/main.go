package main

import (
	"TestGO/configs"
	"TestGO/internal/routes"
	"fmt"
	"log"
	"net/http"

	"github.com/rs/cors" // Importe o pacote
)

func main() {
	err := configs.ConnectDB()
	if err != nil {
		log.Fatalf("Erro ao conectar no banco: %v", err)
	}

	router := http.NewServeMux()
	routes.RegisterRoutes(router)

	mainMux := http.NewServeMux()
	mainMux.Handle("/api/", http.StripPrefix("/api", router))

	// Configuração do CORS
	c := cors.New(cors.Options{
		// Lista de origens permitidas (seu frontend React)
		AllowedOrigins: []string{"http://localhost:5173"},

		// Métodos HTTP que seu frontend pode usar
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},

		// Cabeçalhos que o frontend pode enviar
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},

		// Permite que o frontend envie credenciais (como cookies ou tokens de autorização)
		AllowCredentials: true,

		// Tempo de cache para a preflight request (em segundos)
		MaxAge: 300,
	})

	// Envolve o seu roteador principal com o handler do CORS
	handler := c.Handler(mainMux)

	fmt.Println("Servidor rodando em http://localhost:8080")
	fmt.Println("Rotas da API disponíveis em /api/")

	// Inicia o servidor com o handler que inclui o CORS
	log.Fatal(http.ListenAndServe(":8080", handler))
}
