package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"TestGO/internal/interfaces/services"
)

type AuthMiddleware struct {
	authService services.AuthService
}

// NewAuthMiddleware cria uma nova instância do middleware de autenticação
func NewAuthMiddleware(authService services.AuthService) *AuthMiddleware {
	return &AuthMiddleware{
		authService: authService,
	}
}

// RequireAuth middleware que requer autenticação
func (m *AuthMiddleware) RequireAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extrair token do header Authorization
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Authorization header required",
			})
			c.Abort()
			return
		}

		// Verificar formato Bearer token
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid authorization header format",
			})
			c.Abort()
			return
		}

		token := tokenParts[1]

		// Validar token
		userID, err := m.authService.ValidateJWT(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid or expired token",
			})
			c.Abort()
			return
		}

		// Adicionar user_id ao contexto
		c.Set("user_id", userID)

		c.Next()
	}
}

// OptionalAuth middleware que permite autenticação opcional
func (m *AuthMiddleware) OptionalAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extrair token do header Authorization
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.Next()
			return
		}

		// Verificar formato Bearer token
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			c.Next()
			return
		}

		token := tokenParts[1]

		// Validar token
		userID, err := m.authService.ValidateJWT(token)
		if err != nil {
			c.Next()
			return
		}

		// Adicionar user_id ao contexto se válido
		c.Set("user_id", userID)

		c.Next()
	}
}