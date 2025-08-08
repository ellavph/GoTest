// Package database_models contém modelos que espelham as tabelas do banco de dados
// Estes modelos são usados como referência para migrations e devem estar
// sincronizados com o schema atual do banco de dados
package database_models

import "time"

// User representa a estrutura da tabela 'users' no banco de dados
// Este modelo é usado como referência para migrations e não deve ser
// confundido com a entidade de domínio User
type User struct {
	ID        int64     `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	Name      string    `json:"name"`
	CompanyID *int64    `json:"company_id"`
	Company   *Company  `json:"company,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
