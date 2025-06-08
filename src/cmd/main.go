package main

import (
	"desafio_mercado_livre/src/internal/infra/database"
	"desafio_mercado_livre/src/internal/infra/database/repository"
	log "desafio_mercado_livre/src/internal/infra/logger"
	gin "desafio_mercado_livre/src/internal/infra/router"
	"desafio_mercado_livre/src/internal/infra/web/controller"
	"desafio_mercado_livre/src/internal/infra/web/webserver"
	"desafio_mercado_livre/src/internal/usecase"
)

func main() {
	// Inicializa o banco de dados
	db, err := database.NewDB()
	if err != nil {
		panic(err)
	}

	productRepo := repository.NewProductItemRepository(db.DB)
	productUseCase := usecase.NewProductDetailUseCase(productRepo)
	logger := log.NewLogger()
	productController := controller.NewProductDetailController(logger, productUseCase)
	router := gin.NewRouter(productController)
	server := webserver.NewWebServer(logger, router, ":8080")
	server.Start()
}
