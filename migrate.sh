#!/bin/bash

# Carregar variáveis do .env
if [ -f .env ]; then
    export $(cat .env | xargs)
fi

# Aplicar migrations usando goose CLI
goose -dir migrations postgres "$DATABASE_URL" up