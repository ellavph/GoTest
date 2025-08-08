package entities

import (
	"time"

	"github.com/google/uuid"
)

// User representa a entidade de usuário no domínio
type User struct {
	ID        uuid.UUID  `json:"id" db:"id"`
	Username  string     `json:"username" db:"username"`
	Email     string     `json:"email" db:"email"`
	Password  string     `json:"-" db:"password"` // Não expor no JSON
	Name      string     `json:"name" db:"name"`
	CompanyID *uuid.UUID `json:"company_id" db:"company_id"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt time.Time  `json:"updated_at" db:"updated_at"`
}

// NewUser cria uma nova instância de usuário
func NewUser(username, email, hashedPassword string) *User {
	return &User{
		ID:        uuid.New(),
		Username:  username,
		Email:     email,
		Password:  hashedPassword,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

// NewUserWithName cria uma nova instância de usuário com nome
func NewUserWithName(username, email, hashedPassword, name string) *User {
	return &User{
		ID:        uuid.New(),
		Username:  username,
		Email:     email,
		Password:  hashedPassword,
		Name:      name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

// UpdatePassword atualiza a senha do usuário
func (u *User) UpdatePassword(hashedPassword string) {
	u.Password = hashedPassword
	u.UpdatedAt = time.Now()
}

// UpdateProfile atualiza o perfil do usuário
func (u *User) UpdateProfile(username, email, name string) {
	if username != "" {
		u.Username = username
	}
	if email != "" {
		u.Email = email
	}
	if name != "" {
		u.Name = name
	}
	u.UpdatedAt = time.Now()
}