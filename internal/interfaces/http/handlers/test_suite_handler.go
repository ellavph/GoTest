package handlers

import (
	"net/http"
	"strconv"

	"TestGO/internal/interfaces/services"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type TestSuiteHandler struct {
	testSuiteService services.TestSuiteService
	validator        *validator.Validate
}

func NewTestSuiteHandler(testSuiteService services.TestSuiteService) *TestSuiteHandler {
	return &TestSuiteHandler{
		testSuiteService: testSuiteService,
		validator:        validator.New(),
	}
}

// Request structs para o handler
type CreateTestSuiteRequest struct {
	CompanyID      string `json:"company_id" validate:"required,uuid"`
	Name           string `json:"name" validate:"required,min=3,max=100"`
	Method         string `json:"method" validate:"required,oneof=GET POST PUT DELETE PATCH"`
	URL            string `json:"url" validate:"required,url"`
	Headers        string `json:"headers"`
	ExpectedStatus int    `json:"expected_status" validate:"required,min=100,max=599"`
	ExpectedBody   string `json:"expected_body"`
}

type UpdateTestSuiteRequest struct {
	Name           string `json:"name" validate:"omitempty,min=3,max=100"`
	Method         string `json:"method" validate:"omitempty,oneof=GET POST PUT DELETE PATCH"`
	URL            string `json:"url" validate:"omitempty,url"`
	Headers        string `json:"headers"`
	ExpectedStatus int    `json:"expected_status" validate:"omitempty,min=100,max=599"`
	ExpectedBody   string `json:"expected_body"`
}

// Create godoc
// @Summary Criar nova suíte de teste
// @Description Cria uma nova suíte de teste no sistema
// @Tags test-suites
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body CreateTestSuiteRequest true "Dados da suíte de teste"
// @Success 201 {object} entities.TestSuite "Suíte de teste criada com sucesso"
// @Failure 400 {object} map[string]interface{} "Dados inválidos"
// @Failure 500 {object} map[string]interface{} "Erro interno do servidor"
// @Router /test-suites [post]
func (h *TestSuiteHandler) Create(c *gin.Context) {
	var req CreateTestSuiteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	if err := h.validator.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Converter string para UUID
	companyID, err := uuid.Parse(req.CompanyID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid company ID"})
		return
	}

	createReq := &services.CreateTestSuiteRequest{
		CompanyID:      companyID,
		Name:           req.Name,
		Method:         req.Method,
		URL:            req.URL,
		Headers:        req.Headers,
		ExpectedStatus: req.ExpectedStatus,
		ExpectedBody:   req.ExpectedBody,
	}

	testSuite, err := h.testSuiteService.Create(c.Request.Context(), createReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create test suite"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":              testSuite.ID,
		"company_id":      testSuite.CompanyID,
		"name":            testSuite.Name,
		"method":          testSuite.Method,
		"url":             testSuite.URL,
		"headers":         testSuite.Headers,
		"expected_status": testSuite.ExpectedStatus,
		"expected_body":   testSuite.ExpectedBody,
		"created_at":      testSuite.CreatedAt,
		"updated_at":      testSuite.UpdatedAt,
	})
}

// GetByID godoc
// @Summary Obter suíte de teste por ID
// @Description Retorna os dados de uma suíte de teste específica pelo ID
// @Tags test-suites
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "ID da suíte de teste"
// @Success 200 {object} entities.TestSuite "Suíte de teste encontrada"
// @Failure 400 {object} map[string]interface{} "ID inválido"
// @Failure 404 {object} map[string]interface{} "Suíte de teste não encontrada"
// @Router /test-suites/{id} [get]
func (h *TestSuiteHandler) GetByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid test suite ID"})
		return
	}

	testSuite, err := h.testSuiteService.GetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Test suite not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":              testSuite.ID,
		"company_id":      testSuite.CompanyID,
		"name":            testSuite.Name,
		"method":          testSuite.Method,
		"url":             testSuite.URL,
		"headers":         testSuite.Headers,
		"expected_status": testSuite.ExpectedStatus,
		"expected_body":   testSuite.ExpectedBody,
		"created_at":      testSuite.CreatedAt,
		"updated_at":      testSuite.UpdatedAt,
	})
}

// Update godoc
// @Summary Atualizar suíte de teste
// @Description Atualiza os dados de uma suíte de teste existente
// @Tags test-suites
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "ID da suíte de teste"
// @Param request body UpdateTestSuiteRequest true "Dados para atualização"
// @Success 200 {object} entities.TestSuite "Suíte de teste atualizada com sucesso"
// @Failure 400 {object} map[string]interface{} "Dados inválidos"
// @Failure 500 {object} map[string]interface{} "Erro interno do servidor"
// @Router /test-suites/{id} [put]
func (h *TestSuiteHandler) Update(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid test suite ID"})
		return
	}

	var req UpdateTestSuiteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	if err := h.validator.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updateReq := &services.UpdateTestSuiteRequest{
		Name:           req.Name,
		Method:         req.Method,
		URL:            req.URL,
		Headers:        req.Headers,
		ExpectedStatus: req.ExpectedStatus,
		ExpectedBody:   req.ExpectedBody,
	}

	testSuite, err := h.testSuiteService.Update(c.Request.Context(), id, updateReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update test suite"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":              testSuite.ID,
		"company_id":      testSuite.CompanyID,
		"name":            testSuite.Name,
		"method":          testSuite.Method,
		"url":             testSuite.URL,
		"headers":         testSuite.Headers,
		"expected_status": testSuite.ExpectedStatus,
		"expected_body":   testSuite.ExpectedBody,
		"created_at":      testSuite.CreatedAt,
		"updated_at":      testSuite.UpdatedAt,
	})
}

// Delete godoc
// @Summary Deletar suíte de teste
// @Description Remove uma suíte de teste do sistema
// @Tags test-suites
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "ID da suíte de teste"
// @Success 200 {object} map[string]interface{} "Suíte de teste deletada com sucesso"
// @Failure 400 {object} map[string]interface{} "ID inválido"
// @Failure 500 {object} map[string]interface{} "Erro interno do servidor"
// @Router /test-suites/{id} [delete]
func (h *TestSuiteHandler) Delete(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid test suite ID"})
		return
	}

	err = h.testSuiteService.Delete(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete test suite"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Test suite deleted successfully"})
}

// List godoc
// @Summary Listar suítes de teste
// @Description Retorna uma lista paginada de suítes de teste
// @Tags test-suites
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param company_id query string false "ID da empresa para filtrar"
// @Param limit query int false "Limite de resultados" default(10)
// @Param offset query int false "Offset para paginação" default(0)
// @Success 200 {array} entities.TestSuite "Lista de suítes de teste"
// @Failure 500 {object} map[string]interface{} "Erro interno do servidor"
// @Router /test-suites [get]
func (h *TestSuiteHandler) List(c *gin.Context) {
	limitStr := c.DefaultQuery("limit", "10")
	offsetStr := c.DefaultQuery("offset", "0")
	companyIDStr := c.Query("company_id")

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		limit = 10
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil || offset < 0 {
		offset = 0
	}

	req := &services.ListTestSuitesRequest{
		Limit:  limit,
		Offset: offset,
	}

	// Se company_id foi fornecido, adicionar ao filtro
	if companyIDStr != "" {
		companyID, err := uuid.Parse(companyIDStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid company ID"})
			return
		}
		req.CompanyID = companyID
	}

	response, err := h.testSuiteService.List(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list test suites"})
		return
	}

	c.JSON(http.StatusOK, response.TestSuites)
}

// GetByCompanyID godoc
// @Summary Obter suítes de teste por empresa
// @Description Retorna todas as suítes de teste de uma empresa específica
// @Tags test-suites
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param company_id path string true "ID da empresa"
// @Success 200 {array} entities.TestSuite "Suítes de teste da empresa"
// @Failure 400 {object} map[string]interface{} "ID inválido"
// @Failure 500 {object} map[string]interface{} "Erro interno do servidor"
// @Router /companies/{company_id}/test-suites [get]
func (h *TestSuiteHandler) GetByCompanyID(c *gin.Context) {
	companyIDParam := c.Param("company_id")
	companyID, err := uuid.Parse(companyIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid company ID"})
		return
	}

	testSuites, err := h.testSuiteService.GetByCompanyID(c.Request.Context(), companyID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get test suites"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"test_suites": testSuites,
		"count":       len(testSuites),
	})
}
