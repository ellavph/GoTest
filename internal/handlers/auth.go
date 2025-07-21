package handlers

import (
	"TestGO/internal/dto"
	"TestGO/internal/repositories"
	"TestGO/internal/services"
	"TestGO/internal/utils"
	"context"
	"encoding/json"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	var req dto.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		utils.RespondWithError(w, http.StatusMethodNotAllowed, "Método não permitido")
		return
	}

	user, err := repositories.GetUserByUsername(context.Background(), req.Username)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Erro ao decodificar dados de login")
		return
	}

	if !services.CheckPassword(user.Password, req.Password) {
		utils.RespondWithError(w, http.StatusUnauthorized, "Usuário ou senha inválidos")
		return
	}

	token, err := services.GenerateJWT(user)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Erro ao gerar token")
		return
	}
	respData := map[string]string{"token": token}
	utils.RespondWithSuccess(w, http.StatusOK, "Login realizado com sucesso", respData)
}
