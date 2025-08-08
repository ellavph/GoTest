package handlers

import (
	"net/http"
	"strings"

	"TestGO/internal/interfaces/services"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type AuthHandler struct {
	authService services.AuthService
	validator   *validator.Validate
}

func NewAuthHandler(authService services.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
		validator:   validator.New(),
	}
}

type RegisterRequest struct {
	Username string `json:"username" validate:"required,min=3,max=50"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
	Name     string `json:"name" validate:"omitempty,min=2,max=100"`
}

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

// Register godoc
// @Summary Registrar novo usuário
// @Description Cria uma nova conta de usuário no sistema
// @Tags auth
// @Accept json
// @Produce json
// @Param request body RegisterRequest true "Dados do usuário"
// @Success 201 {object} map[string]interface{} "Usuário registrado com sucesso"
// @Failure 400 {object} map[string]interface{} "Dados inválidos"
// @Failure 409 {object} map[string]interface{} "Usuário já existe"
// @Failure 500 {object} map[string]interface{} "Erro interno do servidor"
// @Router /auth/register [post]
func (h *AuthHandler) Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	if err := h.validator.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	registerReq := &services.RegisterRequest{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
		Name:     req.Name,
	}

	user, err := h.authService.Register(c.Request.Context(), registerReq)
	if err != nil {
		if strings.Contains(err.Error(), "already exists") {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
		"user": gin.H{
			"id":       user.ID,
			"username": user.Username,
			"email":    user.Email,
		},
	})
}

// Login godoc
// @Summary Fazer login
// @Description Autentica um usuário e retorna tokens JWT
// @Tags auth
// @Accept json
// @Produce json
// @Param request body LoginRequest true "Credenciais de login"
// @Success 200 {object} map[string]interface{} "Login realizado com sucesso"
// @Failure 400 {object} map[string]interface{} "Dados inválidos"
// @Failure 401 {object} map[string]interface{} "Credenciais inválidas"
// @Failure 500 {object} map[string]interface{} "Erro interno do servidor"
// @Router /auth/login [post]
func (h *AuthHandler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	if err := h.validator.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	loginReq := &services.LoginRequest{
		Username: req.Username,
		Password: req.Password,
	}

	result, err := h.authService.Login(c.Request.Context(), loginReq)
	if err != nil {
		if strings.Contains(err.Error(), "invalid credentials") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to login"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":      "Login successful",
		"access_token": result.Token,
		"expires_at":   result.ExpiresAt,
	})
}

// Logout godoc
// @Summary Fazer logout
// @Description Realiza logout do usuário (remove token no frontend)
// @Tags auth
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{} "Logout realizado com sucesso"
// @Router /auth/logout [post]
func (h *AuthHandler) Logout(c *gin.Context) {
	// Na Clean Architecture, logout pode ser implementado no frontend
	// removendo o token do storage local
	c.JSON(http.StatusOK, gin.H{"message": "Logout successful"})
}

// RefreshToken godoc
// @Summary Renovar token de acesso
// @Description Renova o token de acesso usando refresh token
// @Tags auth
// @Accept json
// @Produce json
// @Param request body RefreshTokenRequest true "Refresh token"
// @Success 200 {object} map[string]interface{} "Token renovado com sucesso"
// @Failure 501 {object} map[string]interface{} "Funcionalidade não implementada"
// @Router /auth/refresh [post]
func (h *AuthHandler) RefreshToken(c *gin.Context) {
	// RefreshToken não está implementado na interface atual
	// Pode ser implementado como um caso de uso específico se necessário
	c.JSON(http.StatusNotImplemented, gin.H{"error": "Refresh token not implemented"})
}
