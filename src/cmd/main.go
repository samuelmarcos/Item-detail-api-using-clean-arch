package main

import (
	_ "product_item_api/src/docs"
	"product_item_api/src/internal/infra/database"
	"product_item_api/src/internal/infra/database/repository"
	"product_item_api/src/internal/infra/logger"
	gin "product_item_api/src/internal/infra/router"
	"product_item_api/src/internal/infra/web/controller"
	"product_item_api/src/internal/infra/web/webserver"
	"product_item_api/src/internal/usecase"
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

	createProductUsecase := usecase.NewCreateProductDetailUseCase(productRepo)
	listProductUsecase := usecase.NewListProductDetailUseCase(productRepo)
	getProductUsecase := usecase.NewGetProductDetailUseCase(productRepo)

	createProductDetailController := controller.NewCreateProductDetailController(createProductUsecase, logger)
	listProductDetailController := controller.NewListProductDetailController(listProductUsecase, logger)
	getProductDetailController := controller.NewGetProductDetailController(getProductUsecase, logger)

	router := gin.NewRouter()
	server := webserver.NewWebServer(logger, router, ":8080")
	server.AddHandler("GET", "/api/v1/products", listProductDetailController.GetAllProductDetails)
	server.AddHandler("GET", "/api/v1/products/:id", getProductDetailController.GetProductDetail)
	server.AddHandler("POST", "/api/v1/products", createProductDetailController.CreateProductDetail)
	server.Start()
}
