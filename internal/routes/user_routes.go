package routes

import (
	"TestGO/internal/handlers"
	"net/http"
)

func RegisterUserRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/register", handlers.RegisterHandler)
}
