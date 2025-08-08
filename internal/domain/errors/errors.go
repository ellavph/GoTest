package errors

import (
	"fmt"
	"net/http"
)

// ErrorType representa o tipo de erro
type ErrorType string

const (
	ErrorTypeValidation   ErrorType = "VALIDATION_ERROR"
	ErrorTypeNotFound     ErrorType = "NOT_FOUND"
	ErrorTypeUnauthorized ErrorType = "UNAUTHORIZED"
	ErrorTypeForbidden    ErrorType = "FORBIDDEN"
	ErrorTypeConflict     ErrorType = "CONFLICT"
	ErrorTypeInternal     ErrorType = "INTERNAL_ERROR"
	ErrorTypeBadRequest   ErrorType = "BAD_REQUEST"
)

// DomainError representa um erro de domínio
type DomainError struct {
	Type    ErrorType              `json:"type"`
	Message string                 `json:"message"`
	Code    string                 `json:"code,omitempty"`
	Details map[string]interface{} `json:"details,omitempty"`
}

func (e *DomainError) Error() string {
	return e.Message
}

// HTTPStatus retorna o status HTTP apropriado para o erro
func (e *DomainError) HTTPStatus() int {
	switch e.Type {
	case ErrorTypeValidation, ErrorTypeBadRequest:
		return http.StatusBadRequest
	case ErrorTypeNotFound:
		return http.StatusNotFound
	case ErrorTypeUnauthorized:
		return http.StatusUnauthorized
	case ErrorTypeForbidden:
		return http.StatusForbidden
	case ErrorTypeConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}

// Funções helper para criar erros específicos
func NewValidationError(message string, details map[string]interface{}) *DomainError {
	return &DomainError{
		Type:    ErrorTypeValidation,
		Message: message,
		Details: details,
	}
}

func NewNotFoundError(resource string) *DomainError {
	return &DomainError{
		Type:    ErrorTypeNotFound,
		Message: fmt.Sprintf("%s not found", resource),
	}
}

func NewUnauthorizedError(message string) *DomainError {
	return &DomainError{
		Type:    ErrorTypeUnauthorized,
		Message: message,
	}
}

func NewConflictError(message string) *DomainError {
	return &DomainError{
		Type:    ErrorTypeConflict,
		Message: message,
	}
}

func NewInternalError(message string) *DomainError {
	return &DomainError{
		Type:    ErrorTypeInternal,
		Message: message,
	}
}
