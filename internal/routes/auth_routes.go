package routes

import (
	"TestGO/internal/handlers"
	"net/http"
)

func RegisterAuthRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/login", handlers.LoginHandler)
}
