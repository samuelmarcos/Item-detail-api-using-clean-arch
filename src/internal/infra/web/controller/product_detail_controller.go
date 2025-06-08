package controller

import (
	"desafio_mercado_livre/src/internal/entity"
	"desafio_mercado_livre/src/internal/infra/logger"
	"desafio_mercado_livre/src/internal/usecase"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

type ProductDetailController struct {
	usecase *usecase.ProductDetailUseCase
	logger  logger.Logger
}

func NewProductDetailController(l logger.Logger, usecase *usecase.ProductDetailUseCase) *ProductDetailController {
	return &ProductDetailController{
		usecase: usecase,
		logger:  l,
	}
}

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

	inf := fmt.Sprintf("Retrieve product %v with id %s with success", product, id)
	c.logger.Info(inf)

	ctx.JSON(http.StatusOK, product)
}

func (c *ProductDetailController) GetAllProductDetails(ctx *gin.Context) {
	products, err := c.usecase.GetAllProductDetails(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Failed to fetch products"})
		return
	}
	c.logger.Info("Retrieve products with success", products)

	ctx.JSON(http.StatusOK, products)
}

func (c *ProductDetailController) CreateProductDetail(ctx *gin.Context) {
	var product entity.ProductDetail
	if err := ctx.ShouldBindJSON(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid product data: " + err.Error()})
		return
	}

	if err := c.usecase.CreateProductDetail(ctx, &product); err != nil {
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
		default:
			ctx.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Failed to create product: " + err.Error()})
		}
		return
	}

	inf := fmt.Sprintf("Product detail created with success: %v", product)
	c.logger.Info(inf)

	ctx.JSON(http.StatusCreated, product)
}
