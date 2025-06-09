package main

import (
	_ "desafio_mercado_livre/src/docs"
	"desafio_mercado_livre/src/internal/infra/database"
	"desafio_mercado_livre/src/internal/infra/database/repository"
	"desafio_mercado_livre/src/internal/infra/logger"
	gin "desafio_mercado_livre/src/internal/infra/router"
	"desafio_mercado_livre/src/internal/infra/web/controller"
	"desafio_mercado_livre/src/internal/infra/web/webserver"
	"desafio_mercado_livre/src/internal/usecase"
)

// @title           Product Detail API
// @version         1.0
// @description     API for managing product details with clean architecture
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

func main() {
	// Initialize database
	db, err := database.NewDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	logger := logger.NewLogger()

	productRepo := repository.NewProductItemRepository(db.DB)
	productUseCase := usecase.NewProductDetailUseCase(productRepo)
	productController := controller.NewProductDetailController(logger, productUseCase)
	router := gin.NewRouter()
	server := webserver.NewWebServer(logger, router, ":8080")
	server.AddHandler("GET", "/api/v1/products", productController.GetAllProductDetails)
	server.AddHandler("GET", "/api/v1/products/:id", productController.GetProductDetail)
	server.AddHandler("POST", "/api/v1/products", productController.CreateProductDetail)
	server.Start()
}
