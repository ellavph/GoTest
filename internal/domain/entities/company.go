package entities

import (
	"time"

	"github.com/google/uuid"
)

// Company representa a entidade de empresa no domínio
type Company struct {
	ID        uuid.UUID `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	Email     string    `json:"email" db:"email"`
	Phone     string    `json:"phone" db:"phone"`
	Address   string    `json:"address" db:"address"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// NewCompany cria uma nova instância de empresa
func NewCompany(name, email, phone, address string) *Company {
	return &Company{
		ID:        uuid.New(),
		Name:      name,
		Email:     email,
		Phone:     phone,
		Address:   address,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

// UpdateInfo atualiza as informações da empresa
func (c *Company) UpdateInfo(name, email, phone, address string) {
	if name != "" {
		c.Name = name
	}
	if email != "" {
		c.Email = email
	}
	if phone != "" {
		c.Phone = phone
	}
	if address != "" {
		c.Address = address
	}
	c.UpdatedAt = time.Now()
}