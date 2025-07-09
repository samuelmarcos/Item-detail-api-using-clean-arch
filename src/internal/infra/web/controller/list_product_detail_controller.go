package controller

import (
	"fmt"
	"net/http"
	"product_item_api/src/internal/infra/logger"
	"product_item_api/src/internal/usecase"

	"github.com/gin-gonic/gin"
)

func NewListProductDetailController(usecase usecase.ListProductDetail, logger logger.Logger) *ListDetailController {
	return &ListDetailController{
		usecase: usecase,
		logger:  logger,
	}
}

// GetProductDetail godoc
// @Summary      Get a product by ID
// @Description  Get detailed information about a specific product
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Product ID"
// @Success      200  {object}  entity.ProductDetail
// @Failure      404  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /products/{id} [get]

// GetAllProductDetails godoc
// @Summary      Get all products
// @Description  Get a list of all products with their details
// @Tags         products
// @Accept       json
// @Produce      json
// @Success      200  {array}   entity.ProductDetail
// @Failure      500  {object}  ErrorResponse
// @Router       /products [get]
func (c *ListDetailController) GetAllProductDetails(ctx *gin.Context) {
	products, err := c.usecase.Execute(ctx)
	if err != nil {
		c.logger.Info(err.Error())
		ctx.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Failed to fetch products"})
		return
	}

	inf := fmt.Sprintf("All product details retrieved with success: %v", products)
	c.logger.Info(inf)

	ctx.JSON(http.StatusOK, products)
}
