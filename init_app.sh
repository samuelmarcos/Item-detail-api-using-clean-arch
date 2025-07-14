#!/bin/bash

# Variáveis de ambiente para conexão MySQL
export MYSQL_ROOT_PASSWORD="root"
export MYSQL_DATABASE="orders"
export MYSQL_PASSWORD="root"
export MYSQL_PORT="3306"
export MYSQL_USER="root"
export MYSQL_HOST="localhost"

echo Running app ......
go run src/cmd/main.go

