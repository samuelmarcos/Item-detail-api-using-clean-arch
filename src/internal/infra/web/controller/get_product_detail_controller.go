package controller

import (
	"context"
	"fmt"
	"net/http"
	"product_item_api/src/internal/entity"
	"product_item_api/src/internal/infra/logger"
	"product_item_api/src/internal/usecase"

	"github.com/gin-gonic/gin"
)

type GetProductDetailUseCase interface {
	Execute(ctx context.Context, id string) (usecase.GetProductDetailOutputDTO, error)
}

func NewGetProductDetailController(usecase GetProductDetailUseCase, logger logger.Logger) *GetProductDetailController {
	return &GetProductDetailController{
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
func (c *GetProductDetailController) GetProductDetail(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, ErrorResponse{Error: "Product ID is required"})
		return
	}

	product, err := c.usecase.Execute(ctx, id)
	if err != nil {
		if err == entity.ErrInvalidProductID {
			c.logger.Info(err.Error())
			ctx.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
			return
		}
		ctx.JSON(http.StatusNotFound, ErrorResponse{Error: "Product not found"})
		return
	}

	inf := fmt.Sprintf("Product detail retrieved with success: %v", product)
	c.logger.Info(inf)

	ctx.JSON(http.StatusOK, product)
}
