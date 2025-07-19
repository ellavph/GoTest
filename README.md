TestGO API

API desenvolvida em Go para gerenciamento e controle de execuções de testes automatizados. Esta aplicação fornece endpoints REST, gerenciamento de usuários e autenticação via JWT. Banco de dados utilizado: PostgreSQL. Todo o projeto é organizado em módulos internos para facilitar a manutenção e escalabilidade.
Tecnologias Usadas

Golang 1.24+, PostgreSQL (NeonDB), pgx v5, JWT (github.com/golang-jwt/jwt/v4), bcrypt para segurança das senhas, golang-migrate para migrations e godotenv para variáveis de ambiente.
Estrutura de Pastas

TestGO/
├── cmd/
│ └── app/ (main.go - inicializa o servidor)
│ └── migrate/ (migrate.go - executa migrations)
├── configs/ (conexão com banco)
├── internal/
│ ├── dto/ (requests e responses da API)
│ ├── handlers/ (controladores HTTP)
│ ├── models/ (modelos de dados)
│ ├── repositories/ (operações SQL)
│ ├── routes/ (definição de rotas)
│ └── services/ (regras de negócio: JWT, senhas)
├── migrations/ (arquivos SQL)
├── .env (variáveis sensíveis)
├── .gitignore
├── go.mod
├── go.sum
└── README.md

Como Configurar

Crie um arquivo .env com a variável:
DATABASE_URL=postgresql://SEU_USUARIO:SUA_SENHA@SEU_HOST/SEU_BANCO?sslmode=require

Em seguida:
1️⃣ Instale dependências com go mod tidy
2️⃣ Rode as migrations usando go run cmd/migrate/migrate.go
3️⃣ Inicie a API com go run cmd/app/main.go

A API estará rodando em http://localhost:8080.
Endpoints Disponíveis

POST /login → Login com JWT
POST /register → Cadastro de usuário


A API responde em JSON com padrão estruturado. Recomendado testar via Postman ou Insomnia.