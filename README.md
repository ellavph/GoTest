TestGO API

API desenvolvida em Go para gerenciamento e controle de execuções de testes automatizados. Esta aplicação fornece endpoints REST, gerenciamento de usuários e autenticação via JWT. O banco de dados utilizado é PostgreSQL. Todo o projeto é organizado em módulos internos para facilitar a manutenção e escalabilidade.

Tecnologias Usadas

    Golang 1.24+

    PostgreSQL (NeonDB)

    pgx v5

    JWT (github.com/golang-jwt/jwt/v4)

    bcrypt para segurança das senhas

    golang-migrate para migrations

    godotenv para variáveis de ambiente


Como Configurar

    Crie um arquivo .env na raiz do projeto com a seguinte variável:

    DATABASE_URL=postgresql://SEU_USUARIO:SUA_SENHA@SEU_HOST/SEU_BANCO?sslmode=require

    Em seguida, execute os seguintes comandos no terminal:

        Instale as dependências do projeto:
        Bash

go mod tidy

Rode as migrations para configurar o banco de dados:
Bash

go run cmd/migrate/migrate.go

Inicie a API:
Bash

        go run cmd/app/main.go

    A API estará rodando em http://localhost:8080.

Endpoints Disponíveis

    POST /login → Realiza o login do usuário e retorna um token JWT.

    POST /register → Cadastra um novo usuário no sistema.

A API responde em JSON com um padrão estruturado. Recomenda-se testar os endpoints utilizando ferramentas como Postman ou Insomnia.
