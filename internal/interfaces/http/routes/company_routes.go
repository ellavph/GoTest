package routes

import (
	"TestGO/internal/interfaces/http/handlers"
	"TestGO/internal/interfaces/http/middleware"

	"github.com/gin-gonic/gin"
)

// SetupCompanyRoutes configura as rotas de empresa
func SetupCompanyRoutes(router *gin.RouterGroup, companyHandler *handlers.CompanyHandler, authMiddleware *middleware.AuthMiddleware) {
	companies := router.Group("/companies")
	companies.Use(authMiddleware.RequireAuth())
	{
		companies.POST("", companyHandler.Create)
		companies.GET("", companyHandler.List)
		companies.GET("/:id", companyHandler.GetByID)
		companies.PUT("/:id", companyHandler.Update)
		companies.DELETE("/:id", companyHandler.Delete)
	}
}
