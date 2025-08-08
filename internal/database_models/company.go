// Package database_models contém modelos que espelham as tabelas do banco de dados
// Estes modelos são usados como referência para migrations e devem estar
// sincronizados com o schema atual do banco de dados
package database_models

// Company representa a estrutura da tabela 'companies' no banco de dados
// Este modelo é usado como referência para migrations e não deve ser
// confundido com a entidade de domínio Company
type Company struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
}
