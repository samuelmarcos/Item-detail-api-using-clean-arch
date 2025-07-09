package controller

import (
	"fmt"
	"net/http"
	"product_item_api/src/internal/infra/logger"
	"product_item_api/src/internal/usecase"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func NewCreateProductDetailController(usecase usecase.CreateProductDetail, logger logger.Logger) *CreateProductDetailController {
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

	// Validação explícita
	if err := validate.Struct(product); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse{Error: "Validation failed: " + err.Error()})
		return
	}

	output, err := c.usecase.Execute(ctx, product)

	if err != nil {
		c.logger.Error(err.Error())
		ctx.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	inf := fmt.Sprintf("Product detail created with success: %v", output)
	c.logger.Info(inf)

	ctx.JSON(http.StatusCreated, output) // Changed to return the product itself
}
