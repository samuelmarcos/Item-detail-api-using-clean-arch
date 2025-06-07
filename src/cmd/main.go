package main

import (
	gin "desafio_mercado_livre/src/internal/infra/router"
	"desafio_mercado_livre/src/internal/infra/web/webserver"
)

func main() {
	router := gin.NewRouter()
	webserver := webserver.NewWebServer(router, ":8080")
	webserver.Start()
}
