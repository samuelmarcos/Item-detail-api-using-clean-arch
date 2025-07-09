package controller

import (
	"product_item_api/src/internal/infra/logger"
	"product_item_api/src/internal/usecase"
)

type ErrorResponse struct {
	Error string `json:"error" example:"Error message"`
}

type CreateProductDetailController struct {
	usecase usecase.CreateProductDetail
	logger  logger.Logger
}

type ListDetailController struct {
	usecase usecase.ListProductDetail
	logger  logger.Logger
}

type GetProductDetailController struct {
	usecase usecase.GetProductDetail
	logger  logger.Logger
}

type UpdateDiscountController struct {
	usecase usecase.UpdateDiscount
	logger  logger.Logger
}
