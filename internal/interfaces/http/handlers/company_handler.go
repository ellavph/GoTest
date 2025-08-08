package handlers

import (
	"net/http"
	"strconv"

	"TestGO/internal/interfaces/services"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type CompanyHandler struct {
	companyService services.CompanyService
	validator      *validator.Validate
}

func NewCompanyHandler(companyService services.CompanyService) *CompanyHandler {
	return &CompanyHandler{
		companyService: companyService,
		validator:      validator.New(),
	}
}

type CreateCompanyRequest struct {
	Name    string `json:"name" validate:"required,min=2,max=100"`
	Email   string `json:"email" validate:"required,email"`
	Phone   string `json:"phone" validate:"omitempty,min=10,max=20"`
	Address string `json:"address" validate:"omitempty,max=255"`
}

type UpdateCompanyRequest struct {
	Name    string `json:"name" validate:"omitempty,min=2,max=100"`
	Email   string `json:"email" validate:"omitempty,email"`
	Phone   string `json:"phone" validate:"omitempty,min=10,max=20"`
	Address string `json:"address" validate:"omitempty,max=255"`
}

// Create godoc
// @Summary Criar nova empresa
// @Description Cria uma nova empresa no sistema
// @Tags companies
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body CreateCompanyRequest true "Dados da empresa"
// @Success 201 {object} map[string]interface{} "Empresa criada com sucesso"
// @Failure 400 {object} map[string]interface{} "Dados inválidos"
// @Failure 500 {object} map[string]interface{} "Erro interno do servidor"
// @Router /companies [post]
func (h *CompanyHandler) Create(c *gin.Context) {
	var req CreateCompanyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	if err := h.validator.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createReq := &services.CreateCompanyRequest{
		Name:    req.Name,
		Email:   req.Email,
		Phone:   req.Phone,
		Address: req.Address,
	}

	company, err := h.companyService.Create(c.Request.Context(), createReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create company"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Company created successfully",
		"company": gin.H{
			"id":      company.ID,
			"name":    company.Name,
			"email":   company.Email,
			"phone":   company.Phone,
			"address": company.Address,
		},
	})
}

// GetByID godoc
// @Summary Obter empresa por ID
// @Description Retorna os dados de uma empresa específica pelo ID
// @Tags companies
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "ID da empresa"
// @Success 200 {object} map[string]interface{} "Dados da empresa"
// @Failure 400 {object} map[string]interface{} "ID inválido"
// @Failure 404 {object} map[string]interface{} "Empresa não encontrada"
// @Router /companies/{id} [get]
func (h *CompanyHandler) GetByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid company ID"})
		return
	}

	company, err := h.companyService.GetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Company not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"company": gin.H{
			"id":      company.ID,
			"name":    company.Name,
			"email":   company.Email,
			"phone":   company.Phone,
			"address": company.Address,
		},
	})
}

// Update godoc
// @Summary Atualizar empresa
// @Description Atualiza os dados de uma empresa existente
// @Tags companies
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "ID da empresa"
// @Param request body UpdateCompanyRequest true "Dados para atualização"
// @Success 200 {object} map[string]interface{} "Empresa atualizada com sucesso"
// @Failure 400 {object} map[string]interface{} "Dados inválidos"
// @Failure 500 {object} map[string]interface{} "Erro interno do servidor"
// @Router /companies/{id} [put]
func (h *CompanyHandler) Update(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid company ID"})
		return
	}

	var req UpdateCompanyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	if err := h.validator.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updateReq := &services.UpdateCompanyRequest{
		Name:    req.Name,
		Email:   req.Email,
		Phone:   req.Phone,
		Address: req.Address,
	}

	company, err := h.companyService.Update(c.Request.Context(), id, updateReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update company"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Company updated successfully",
		"company": gin.H{
			"id":      company.ID,
			"name":    company.Name,
			"email":   company.Email,
			"phone":   company.Phone,
			"address": company.Address,
		},
	})
}

// Delete godoc
// @Summary Deletar empresa
// @Description Remove uma empresa do sistema
// @Tags companies
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "ID da empresa"
// @Success 200 {object} map[string]interface{} "Empresa deletada com sucesso"
// @Failure 400 {object} map[string]interface{} "ID inválido"
// @Failure 500 {object} map[string]interface{} "Erro interno do servidor"
// @Router /companies/{id} [delete]
func (h *CompanyHandler) Delete(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid company ID"})
		return
	}

	err = h.companyService.Delete(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete company"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Company deleted successfully"})
}

// List godoc
// @Summary Listar empresas
// @Description Retorna uma lista paginada de empresas
// @Tags companies
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param limit query int false "Limite de resultados" default(10)
// @Param offset query int false "Offset para paginação" default(0)
// @Success 200 {object} map[string]interface{} "Lista de empresas"
// @Failure 500 {object} map[string]interface{} "Erro interno do servidor"
// @Router /companies [get]
func (h *CompanyHandler) List(c *gin.Context) {
	limitStr := c.DefaultQuery("limit", "10")
	offsetStr := c.DefaultQuery("offset", "0")

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		limit = 10
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil || offset < 0 {
		offset = 0
	}

	req := &services.ListCompaniesRequest{
		Limit:  limit,
		Offset: offset,
	}

	response, err := h.companyService.List(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list companies"})
		return
	}

	c.JSON(http.StatusOK, response)
}
