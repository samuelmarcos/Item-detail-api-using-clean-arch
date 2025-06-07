#!/bin/bash

# Compila o projeto Go usando main.go como entrada

echo "Compilando o projeto..."
go build -o bin/app ./src/cmd/main.go

if [ $? -eq 0 ]; then
  echo "Compilação bem-sucedida! Binário gerado em bin/app"
else
  echo "Erro na compilação."
  exit 1
fi 

./bin/app