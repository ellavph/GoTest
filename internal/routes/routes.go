package routes

import (
	"TestGO/internal/handlers"
	"net/http"
)

func RegisterRoutes(mux *http.ServeMux) {
	RegisterAuthRoutes(mux)
	RegisterUserRoutes(mux)
	mux.HandleFunc("/", handlers.HomeHandler)
}
