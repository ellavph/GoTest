package routes

import (
	"TestGO/internal/interfaces/http/handlers"

	"github.com/gin-gonic/gin"
)

// SetupAuthRoutes configura as rotas de autenticação
func SetupAuthRoutes(router *gin.RouterGroup, authHandler *handlers.AuthHandler) {
	auth := router.Group("/auth")
	{
		auth.POST("/register", authHandler.Register)
		auth.POST("/login", authHandler.Login)
		auth.POST("/logout", authHandler.Logout)
		auth.POST("/refresh", authHandler.RefreshToken)
	}
}
