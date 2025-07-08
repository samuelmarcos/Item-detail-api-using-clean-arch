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

type CreateProductDetailUseCase interface {
	Execute(ctx context.Context, input usecase.CreateProductDetailInputDTO) (usecase.CreateProductDetailOutputDTO, error)
}

func NewCreateProductDetailController(usecase CreateProductDetailUseCase, logger logger.Logger) *CreateProductDetailController {
	return &CreateProductDetailController{
		usecase: usecase,
		logger:  logger,
	}
}

// CreateProductDetail godoc
// @Summary      Create a new product
// @Description  Create a new product with all its details
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        product  body      entity.ProductDetail  true  "Product Details"
// @Success      201     {object}  entity.ProductDetail
// @Failure      400     {object}  ErrorResponse
// @Failure      500     {object}  ErrorResponse
// @Router       /products [post]
func (c *CreateProductDetailController) CreateProductDetail(ctx *gin.Context) {
	var product usecase.CreateProductDetailInputDTO
	if err := ctx.ShouldBindJSON(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid product data: " + err.Error()})
		return
	}

	output, err := c.usecase.Execute(ctx, product)
	if err != nil {
		switch err {
		case entity.ErrInvalidProductID:
			ctx.JSON(http.StatusBadRequest, ErrorResponse{Error: "Product ID is required"})
		case entity.ErrInvalidProductName:
			ctx.JSON(http.StatusBadRequest, ErrorResponse{Error: "Product name is required"})
		case entity.ErrInvalidProductPrice:
			ctx.JSON(http.StatusBadRequest, ErrorResponse{Error: "Product price must be greater than zero"})
		case entity.ErrInvalidProductStock:
			ctx.JSON(http.StatusBadRequest, ErrorResponse{Error: "Product stock cannot be negative"})
		case entity.ErrInvalidSellerName:
			ctx.JSON(http.StatusBadRequest, ErrorResponse{Error: "Seller name is required"})
		case entity.ErrInvalidProductDescription, entity.ErrInvalidProductBrand,
			entity.ErrInvalidProductModel, entity.ErrInvalidProductColor,
			entity.ErrInvalidProductCategory, entity.ErrInvalidProductImages,
			entity.ErrInvalidSellerType, entity.ErrInvalidSellerReputation,
			entity.ErrInvalidSellerSales, entity.ErrInvalidWarranty,
			entity.ErrInvalidPaymentOptions, entity.ErrInvalidSpecs,
			entity.ErrInvalidRelatedProducts, entity.ErrInvalidRating,
			entity.ErrInvalidReviewCount, entity.ErrInvalidPurchaseOptions,
			entity.ErrInvalidHighlights:
			ctx.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		default:
			ctx.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Failed to create product: " + err.Error()})
		}
		return
	}

	inf := fmt.Sprintf("Product detail created with success: %v", output)
	c.logger.Info(inf)

	ctx.JSON(http.StatusCreated, output) // Changed to return the product itself
}
