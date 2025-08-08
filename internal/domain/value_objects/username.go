package value_objects

import (
	"fmt"
	"regexp"
	"strings"
)

// Username representa um value object para username
type Username struct {
	value string
}

// NewUsername cria um novo username com validação
func NewUsername(username string) (*Username, error) {
	if err := validateUsername(username); err != nil {
		return nil, err
	}
	
	return &Username{
		value: strings.TrimSpace(username),
	}, nil
}

// String retorna o valor do username
func (u Username) String() string {
	return u.value
}

// Value retorna o valor do username (para compatibilidade)
func (u Username) Value() string {
	return u.value
}

// Length retorna o comprimento do username
func (u Username) Length() int {
	return len(u.value)
}

// Equals verifica se dois usernames são iguais
func (u Username) Equals(other Username) bool {
	return u.value == other.value
}

// IsValid verifica se o username é válido
func (u Username) IsValid() bool {
	return validateUsername(u.value) == nil
}

// validateUsername valida o formato e regras do username
func validateUsername(username string) error {
	if username == "" {
		return fmt.Errorf("username cannot be empty")
	}
	
	username = strings.TrimSpace(username)
	
	if len(username) < 3 {
		return fmt.Errorf("username must be at least 3 characters long")
	}
	
	if len(username) > 50 {
		return fmt.Errorf("username must be at most 50 characters long")
	}
	
	// Regex para validar username: apenas letras, números, underscore e hífen
	usernameRegex := regexp.MustCompile(`^[a-zA-Z0-9_-]+$`)
	
	if !usernameRegex.MatchString(username) {
		return fmt.Errorf("username can only contain letters, numbers, underscores and hyphens")
	}
	
	// Não pode começar ou terminar com underscore ou hífen
	if strings.HasPrefix(username, "_") || strings.HasPrefix(username, "-") ||
		strings.HasSuffix(username, "_") || strings.HasSuffix(username, "-") {
		return fmt.Errorf("username cannot start or end with underscore or hyphen")
	}
	
	// Lista de usernames reservados
	reservedUsernames := []string{
		"admin", "administrator", "root", "system", "api", "www", "mail",
		"ftp", "test", "guest", "user", "null", "undefined", "support",
	}
	
	for _, reserved := range reservedUsernames {
		if strings.EqualFold(username, reserved) {
			return fmt.Errorf("username '%s' is reserved", username)
		}
	}
	
	return nil
}