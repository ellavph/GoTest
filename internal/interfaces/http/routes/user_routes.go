package routes

import (
	"github.com/gin-gonic/gin"
	"TestGO/internal/interfaces/http/handlers"
	"TestGO/internal/interfaces/http/middleware"
)

// SetupUserRoutes configura as rotas de usuário
func SetupUserRoutes(router *gin.RouterGroup, userHandler *handlers.UserHandler, authMiddleware *middleware.AuthMiddleware) {
	users := router.Group("/users")
	users.Use(authMiddleware.RequireAuth())
	{
		users.GET("/profile", userHandler.GetProfile)
		users.PUT("/profile", userHandler.UpdateProfile)
		users.DELETE("/profile", userHandler.DeleteProfile)
		users.PUT("/password", userHandler.ChangePassword)
		users.GET("", userHandler.ListUsers)
		users.GET("/:id", userHandler.GetUserByID)
	}
}