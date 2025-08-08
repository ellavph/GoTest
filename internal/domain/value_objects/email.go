package value_objects

import (
	"fmt"
	"regexp"
	"strings"
)

// Email representa um value object para email
type Email struct {
	value string
}

// NewEmail cria um novo email com validação
func NewEmail(email string) (*Email, error) {
	if err := validateEmail(email); err != nil {
		return nil, err
	}
	
	return &Email{
		value: strings.ToLower(strings.TrimSpace(email)),
	}, nil
}

// String retorna o valor do email
func (e Email) String() string {
	return e.value
}

// Value retorna o valor do email (para compatibilidade)
func (e Email) Value() string {
	return e.value
}

// Domain retorna o domínio do email
func (e Email) Domain() string {
	parts := strings.Split(e.value, "@")
	if len(parts) != 2 {
		return ""
	}
	return parts[1]
}

// LocalPart retorna a parte local do email (antes do @)
func (e Email) LocalPart() string {
	parts := strings.Split(e.value, "@")
	if len(parts) != 2 {
		return ""
	}
	return parts[0]
}

// Equals verifica se dois emails são iguais
func (e Email) Equals(other Email) bool {
	return e.value == other.value
}

// validateEmail valida o formato do email
func validateEmail(email string) error {
	if email == "" {
		return fmt.Errorf("email cannot be empty")
	}
	
	email = strings.TrimSpace(email)
	
	// Regex básica para validação de email
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	
	if !emailRegex.MatchString(email) {
		return fmt.Errorf("invalid email format")
	}
	
	if len(email) > 254 {
		return fmt.Errorf("email too long (max 254 characters)")
	}
	
	return nil
}