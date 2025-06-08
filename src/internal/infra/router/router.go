package gin

import (
	"desafio_mercado_livre/src/internal/infra/web/controller"

	"github.com/gin-gonic/gin"
)

func NewRouter(productController *controller.ProductDetailController) *gin.Engine {
	router := gin.Default()

	api := router.Group("/api/v1")
	{
		items := api.Group("/products")
		{
			items.GET("", productController.GetAllProductDetails)
			items.GET("/:id", productController.GetProductDetail)
			items.POST("", productController.CreateProductDetail)
		}
	}

	return router
}
