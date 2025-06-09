package controller

import (
	"context"
	"desafio_mercado_livre/src/internal/entity"
	"desafio_mercado_livre/src/internal/infra/logger"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductDetailController struct {
	usecase ProductDetailUseCase
	logger  logger.Logger
}

type ErrorResponse struct {
	Error string `json:"error" example:"Error message"`
}

type ProductDetailUseCase interface {
	GetProductDetail(ctx context.Context, id string) (*entity.ProductDetail, error)
	GetAllProductDetails(ctx context.Context) ([]*entity.ProductDetail, error)
	CreateProductDetail(ctx context.Context, product *entity.ProductDetail) error
}

func NewProductDetailController(logger logger.Logger, usecase ProductDetailUseCase) *ProductDetailController {
	return &ProductDetailController{
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
func (c *ProductDetailController) GetProductDetail(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, ErrorResponse{Error: "Product ID is required"})
		return
	}

	product, err := c.usecase.GetProductDetail(ctx, id)
	if err != nil {
		if err == entity.ErrInvalidProductID {
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

// GetAllProductDetails godoc
// @Summary      Get all products
// @Description  Get a list of all products with their details
// @Tags         products
// @Accept       json
// @Produce      json
// @Success      200  {array}   entity.ProductDetail
// @Failure      500  {object}  ErrorResponse
// @Router       /products [get]
func (c *ProductDetailController) GetAllProductDetails(ctx *gin.Context) {
	products, err := c.usecase.GetAllProductDetails(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Failed to fetch products"})
		return
	}

	inf := fmt.Sprintf("All product details retrieved with success: %v", products)
	c.logger.Info(inf)

	ctx.JSON(http.StatusOK, products)
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
func (c *ProductDetailController) CreateProductDetail(ctx *gin.Context) {
	var product entity.ProductDetail
	if err := ctx.ShouldBindJSON(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid product data: " + err.Error()})
		return
	}

	if err := c.usecase.CreateProductDetail(ctx, &product); err != nil {
		switch err {
		case entity.ErrInvalidProductID, entity.ErrInvalidProductName, entity.ErrInvalidProductDescription,
			entity.ErrInvalidProductBrand, entity.ErrInvalidProductModel, entity.ErrInvalidProductColor,
			entity.ErrInvalidProductCategory, entity.ErrInvalidProductImages, entity.ErrInvalidProductPrice,
			entity.ErrInvalidProductStock, entity.ErrInvalidSellerName, entity.ErrInvalidSellerType,
			entity.ErrInvalidSellerReputation, entity.ErrInvalidSellerSales, entity.ErrInvalidWarranty,
			entity.ErrInvalidPaymentOptions, entity.ErrInvalidSpecs, entity.ErrInvalidRelatedProducts,
			entity.ErrInvalidRating, entity.ErrInvalidReviewCount, entity.ErrInvalidPurchaseOptions,
			entity.ErrInvalidHighlights:
			ctx.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		default:
			ctx.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Failed to create product: " + err.Error()})
		}
		return
	}

	inf := fmt.Sprintf("Product detail created with success: %v", product)
	c.logger.Info(inf)

	ctx.JSON(http.StatusCreated, product)
}
