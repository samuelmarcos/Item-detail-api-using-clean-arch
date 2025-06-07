package main

import (
	log "desafio_mercado_livre/src/internal/infra/logger"
	gin "desafio_mercado_livre/src/internal/infra/router"
	"desafio_mercado_livre/src/internal/infra/web/webserver"
)

func main() {
	log := log.NewLogger()
	router := gin.NewRouter()
	webserver := webserver.NewWebServer(log, router, ":8080")
	webserver.Start()
}
