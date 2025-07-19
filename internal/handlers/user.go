package handlers

import (
	"TestGO/internal/dto"
	"TestGO/internal/models"
	"TestGO/internal/repositories"
	"TestGO/internal/services"
	"TestGO/internal/utils"
	"context"
	"encoding/json"
	"log"
	"net/http"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.RespondWithError(w, http.StatusMethodNotAllowed, "Método não permitido")
		return
	}

	var req dto.RegisterRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Dados inválidos")
		return
	}

	// Verificar se username já existe (opcional)
	_, err = repositories.GetUserByUsername(context.Background(), req.Username)
	if err == nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Username já cadastrado")
		return
	}

	// Criar modelo User
	hashedPassword, err := services.HashPassword(req.Password)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Erro ao processar senha")
		return
	}

	user := models.User{
		Username: req.Username,
		Email:    req.Email,
		Password: hashedPassword,
	}

	err = repositories.CreateUser(context.Background(), &user)
	if err != nil {
		log.Printf("Erro ao salvar usuário: %v", err) // <--- aqui o log
		utils.RespondWithError(w, http.StatusInternalServerError, "Erro ao salvar usuário")
		return
	}

	resp := dto.RegisterResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}
	utils.RespondWithSuccess(w, http.StatusCreated, "Usuário criado com sucesso", resp)
}
