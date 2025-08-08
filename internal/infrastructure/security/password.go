package security

import (
	"golang.org/x/crypto/bcrypt"
)

type PasswordService struct {
	cost int
}

// NewPasswordService cria uma nova instância do serviço de senha
func NewPasswordService(cost int) *PasswordService {
	if cost < bcrypt.MinCost {
		cost = bcrypt.DefaultCost
	}
	return &PasswordService{cost: cost}
}

// HashPassword gera um hash da senha
func (p *PasswordService) HashPassword(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), p.cost)
	if err != nil {
		return "", err
	}
	return string(hashedBytes), nil
}

// CheckPassword verifica se a senha corresponde ao hash
func (p *PasswordService) CheckPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
