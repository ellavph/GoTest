package security

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type JWTService struct {
	secretKey     string
	accessExpiry  time.Duration
	refreshExpiry time.Duration
}

// NewJWTService cria uma nova instância do serviço JWT
func NewJWTService(secretKey string, accessExpiry, refreshExpiry time.Duration) *JWTService {
	return &JWTService{
		secretKey:     secretKey,
		accessExpiry:  accessExpiry,
		refreshExpiry: refreshExpiry,
	}
}

// GenerateAccessToken gera um token de acesso
func (j *JWTService) GenerateAccessToken(userID uuid.UUID, username, email string, companyID *uuid.UUID) (string, error) {
	claims := jwt.MapClaims{
		"user_id":    userID.String(),
		"username":   username,
		"email":      email,
		"company_id": companyID,
		"type":       "access",
		"exp":        time.Now().Add(j.accessExpiry).Unix(),
		"iat":        time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.secretKey))
}

// GenerateRefreshToken gera um token de refresh
func (j *JWTService) GenerateRefreshToken(userID uuid.UUID) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID.String(),
		"type":    "refresh",
		"exp":     time.Now().Add(j.refreshExpiry).Unix(),
		"iat":     time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.secretKey))
}

// GenerateTokenPair gera um par de tokens (access e refresh)
func (j *JWTService) GenerateTokenPair(userID uuid.UUID, username, email string, companyID *uuid.UUID) (*TokenPair, error) {
	accessToken, err := j.GenerateAccessToken(userID, username, email, companyID)
	if err != nil {
		return nil, fmt.Errorf("failed to generate access token: %w", err)
	}

	refreshToken, err := j.GenerateRefreshToken(userID)
	if err != nil {
		return nil, fmt.Errorf("failed to generate refresh token: %w", err)
	}

	return &TokenPair{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresIn:    int64(j.accessExpiry.Seconds()),
	}, nil
}

// ValidateToken valida um token JWT e retorna as claims
func (j *JWTService) ValidateToken(tokenString string) (*TokenClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(j.secretKey), nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}

	// Verificar se é um access token
	tokenType, ok := claims["type"].(string)
	if !ok || tokenType != "access" {
		return nil, fmt.Errorf("invalid token type")
	}

	userIDStr, ok := claims["user_id"].(string)
	if !ok {
		return nil, fmt.Errorf("invalid user_id claim")
	}

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return nil, fmt.Errorf("invalid user_id format: %w", err)
	}

	username, _ := claims["username"].(string)
	email, _ := claims["email"].(string)
	
	var companyID *uuid.UUID
	if companyIDStr, ok := claims["company_id"].(string); ok && companyIDStr != "" {
		if parsed, err := uuid.Parse(companyIDStr); err == nil {
			companyID = &parsed
		}
	}

	exp, ok := claims["exp"].(float64)
	if !ok {
		return nil, fmt.Errorf("invalid exp claim")
	}

	return &TokenClaims{
		UserID:    userID,
		Username:  username,
		Email:     email,
		CompanyID: companyID,
		ExpiresAt: int64(exp),
	}, nil
}

// RefreshToken gera um novo access token a partir de um refresh token válido
func (j *JWTService) RefreshToken(refreshTokenString string) (*TokenPair, error) {
	token, err := jwt.Parse(refreshTokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(j.secretKey), nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to parse refresh token: %w", err)
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid refresh token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}

	// Verificar se é um refresh token
	tokenType, ok := claims["type"].(string)
	if !ok || tokenType != "refresh" {
		return nil, fmt.Errorf("invalid token type")
	}

	userIDStr, ok := claims["user_id"].(string)
	if !ok {
		return nil, fmt.Errorf("invalid user_id claim")
	}

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return nil, fmt.Errorf("invalid user_id format: %w", err)
	}

	// TODO: Em produção, buscar dados atualizados do usuário no banco
	// Por enquanto, mantemos as claims básicas para funcionalidade
	// IMPORTANTE: Este é um ponto de melhoria - implementar busca no repositório
	
	accessClaims := jwt.MapClaims{
		"user_id": userID.String(),
		"type":    "access",
		"exp":     time.Now().Add(j.accessExpiry).Unix(),
		"iat":     time.Now().Unix(),
		// TODO: Adicionar username, email, company_id buscando do banco
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessTokenString, err := accessToken.SignedString([]byte(j.secretKey))
	if err != nil {
		return nil, fmt.Errorf("failed to generate access token: %w", err)
	}

	return &TokenPair{
		AccessToken:  accessTokenString,
		RefreshToken: refreshTokenString, // Mantém o mesmo refresh token
		ExpiresIn:    int64(j.accessExpiry.Seconds()),
	}, nil
}