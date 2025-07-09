package controller

import "product_item_api/src/internal/infra/logger"

type ErrorResponse struct {
	Error string `json:"error" example:"Error message"`
}

type CreateProductDetailController struct {
	usecase CreateProductDetailUseCase
	logger  logger.Logger
}

type ListDetailController struct {
	usecase ListProductDetailUseCase
	logger  logger.Logger
}

type GetProductDetailController struct {
	usecase GetProductDetailUseCase
	logger  logger.Logger
}

type UpdateDiscountController struct {
	usecase UpdateDiscountUseCase
	logger  logger.Logger
}
