package utils

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status    bool        `json:"status"`
	Descricao string      `json:"descricao"`
	Data      interface{} `json:"data"`
}

func RespondWithError(w http.ResponseWriter, code int, descricao string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	resp := Response{
		Status:    false,
		Descricao: descricao,
		Data:      struct{}{},
	}
	json.NewEncoder(w).Encode(resp)
}

func RespondWithSuccess(w http.ResponseWriter, code int, descricao string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	resp := Response{
		Status:    true,
		Descricao: descricao,
		Data:      data,
	}
	json.NewEncoder(w).Encode(resp)
}

// SendErrorResponse envia uma resposta de erro
func SendErrorResponse(w http.ResponseWriter, message string, statusCode int) {
	RespondWithError(w, statusCode, message)
}

// SendSuccessResponse envia uma resposta de sucesso
func SendSuccessResponse(w http.ResponseWriter, message string, data interface{}) {
	RespondWithSuccess(w, http.StatusOK, message, data)
}
