package main

import (
	"desafio_mercado_livre/src/internal/infra/database"
	"desafio_mercado_livre/src/internal/infra/database/repository"
	log "desafio_mercado_livre/src/internal/infra/logger"
	gin "desafio_mercado_livre/src/internal/infra/router"
	"desafio_mercado_livre/src/internal/infra/web/webserver"
)

func main() {

	db, err := database.NewDB()
	if err != nil {
		panic(err)
	}
	repository.NewProductItemRepository(db.DB)

	log := log.NewLogger()
	router := gin.NewRouter()
	webserver := webserver.NewWebServer(log, router, ":8080")
	webserver.Start()
}
