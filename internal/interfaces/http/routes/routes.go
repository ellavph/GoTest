package routes

import (
	"TestGO/internal/infrastructure/container"

	"github.com/gin-gonic/gin"
)

// SetupRoutes configura todas as rotas da aplicação
func SetupRoutes(router *gin.Engine, container *container.Container) {
	// Rotas de autenticação (públicas)
	authRoutes := router.Group("/api/auth")
	{
		authRoutes.POST("/register", container.AuthHandler.Register)
		authRoutes.POST("/login", container.AuthHandler.Login)
		authRoutes.POST("/logout", container.AuthMiddleware.RequireAuth(), container.AuthHandler.Logout)
		authRoutes.POST("/refresh", container.AuthHandler.RefreshToken)
	}

	// Rotas protegidas
	api := router.Group("/api")
	api.Use(container.AuthMiddleware.RequireAuth())
	{
		// Rotas de usuário
		userRoutes := api.Group("/users")
		{
			userRoutes.GET("/profile", container.UserHandler.GetProfile)
			userRoutes.PUT("/profile", container.UserHandler.UpdateProfile)
			userRoutes.DELETE("/profile", container.UserHandler.DeleteProfile)
			userRoutes.POST("/change-password", container.UserHandler.ChangePassword)
			userRoutes.GET("", container.UserHandler.ListUsers)
			userRoutes.GET("/:id", container.UserHandler.GetUserByID)
		}

		// Rotas de empresa
		companyRoutes := api.Group("/companies")
		{
			companyRoutes.POST("", container.CompanyHandler.Create)
			companyRoutes.GET("/:id", container.CompanyHandler.GetByID)
			companyRoutes.PUT("/:id", container.CompanyHandler.Update)
			companyRoutes.DELETE("/:id", container.CompanyHandler.Delete)
			companyRoutes.GET("", container.CompanyHandler.List)
		}
	}

	// Rota de health check (pública)
	router.GET("/api/v1/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"message": "API is running",
		})
	})
}
