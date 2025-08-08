package security

import "github.com/google/uuid"

// TokenClaims representa as claims de um token JWT
type TokenClaims struct {
	UserID    uuid.UUID  `json:"user_id"`
	Username  string     `json:"username"`
	Email     string     `json:"email"`
	CompanyID *uuid.UUID `json:"company_id"`
	ExpiresAt int64      `json:"exp"`
}

// TokenPair representa um par de tokens (access e refresh)
type TokenPair struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int64  `json:"expires_in"`
}