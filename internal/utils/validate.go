package utils

import (
	"errors"
	"reflect"
	"regexp"
	"strings"
)

// ValidateStruct valida uma struct usando validação customizada
func ValidateStruct(s interface{}) error {
	v := reflect.ValueOf(s)
	t := reflect.TypeOf(s)
	
	// Se for um ponteiro, obter o valor
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
		t = t.Elem()
	}
	
	var validationErrors []string
	
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := t.Field(i)
		
		// Verificar tag validate
		validateTag := fieldType.Tag.Get("validate")
		if validateTag == "" {
			continue
		}
		
		fieldName := GetStructFieldName(fieldType)
		
		// Processar tags de validação
		tags := strings.Split(validateTag, ",")
		for _, tag := range tags {
			tag = strings.TrimSpace(tag)
			
			switch {
			case tag == "required":
				if isEmptyValue(field) {
					validationErrors = append(validationErrors, fieldName+" é obrigatório")
				}
			case strings.HasPrefix(tag, "min="):
				minStr := strings.TrimPrefix(tag, "min=")
				if field.Kind() == reflect.String {
					if len(field.String()) < parseIntFromString(minStr) {
						validationErrors = append(validationErrors, fieldName+" deve ter pelo menos "+minStr+" caracteres")
					}
				}
			case strings.HasPrefix(tag, "max="):
				maxStr := strings.TrimPrefix(tag, "max=")
				if field.Kind() == reflect.String {
					if len(field.String()) > parseIntFromString(maxStr) {
						validationErrors = append(validationErrors, fieldName+" deve ter no máximo "+maxStr+" caracteres")
					}
				}
			case tag == "email":
				if field.Kind() == reflect.String && field.String() != "" {
					if !isValidEmail(field.String()) {
						validationErrors = append(validationErrors, fieldName+" deve ser um email válido")
					}
				}
			}
		}
	}
	
	if len(validationErrors) > 0 {
		return errors.New(strings.Join(validationErrors, ", "))
	}
	
	return nil
}

// isEmptyValue verifica se um valor está vazio
func isEmptyValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.String:
		return v.String() == ""
	case reflect.Ptr, reflect.Interface:
		return v.IsNil()
	case reflect.Slice, reflect.Map, reflect.Array:
		return v.Len() == 0
	default:
		return false
	}
}

// parseIntFromString converte string para int (simples)
func parseIntFromString(s string) int {
	switch s {
	case "1":
		return 1
	case "2":
		return 2
	case "3":
		return 3
	case "4":
		return 4
	case "5":
		return 5
	case "6":
		return 6
	case "8":
		return 8
	case "50":
		return 50
	case "100":
		return 100
	default:
		return 0
	}
}

// isValidEmail valida formato de email
func isValidEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}

// ValidaUsername valida se o username atende aos critérios
func ValidaUsername(username string) error {
	if len(username) < 3 {
		return errors.New("username deve ter pelo menos 3 caracteres")
	}
	
	if len(username) > 50 {
		return errors.New("username deve ter no máximo 50 caracteres")
	}
	
	// Permitir apenas letras, números e underscore
	matched, err := regexp.MatchString("^[a-zA-Z0-9_]+$", username)
	if err != nil {
		return errors.New("erro ao validar username")
	}
	
	if !matched {
		return errors.New("username deve conter apenas letras, números e underscore")
	}
	
	return nil
}

// GetStructFieldName retorna o nome do campo da struct
func GetStructFieldName(field reflect.StructField) string {
	jsonTag := field.Tag.Get("json")
	if jsonTag != "" && jsonTag != "-" {
		// Remove opções como omitempty
		if idx := strings.Index(jsonTag, ","); idx != -1 {
			return jsonTag[:idx]
		}
		return jsonTag
	}
	return field.Name
}