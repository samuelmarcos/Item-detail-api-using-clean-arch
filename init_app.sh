#!/bin/bash

# Variáveis de ambiente para conexão MySQL
export MYSQL_ROOT_PASSWORD="root"
export MYSQL_DATABASE="orders"
export MYSQL_PASSWORD="root"
export MYSQL_PORT="3306"
export MYSQL_USER="root"

# Build da aplicação
GOOS=linux GOARCH=amd64 go build -o bin/app ./src/cmd/main.go

# Execução da aplicação
./bin/app 