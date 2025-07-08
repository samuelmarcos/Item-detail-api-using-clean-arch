package usecase

import (
	"context"
	"product_item_api/src/internal/entity"
	"time"
)

// Remover CreateProductDetailInputDTO e usar ProductDetail sem o campo ID
type CreateProductDetailInputDTO struct {
	entity.ProductDetail
}

type CreateProductDetailUseCase struct {
	Repository entity.ProductDetailRepository
}

type CreateProductDetailOutputDTO struct {
	CreateProductDetailInputDTO
}

func NewCreateProductDetailUseCase(r entity.ProductDetailRepository) *CreateProductDetailUseCase {
	return &CreateProductDetailUseCase{
		Repository: r,
	}
}

func (uc *CreateProductDetailUseCase) Execute(ctx context.Context, input CreateProductDetailInputDTO) (CreateProductDetailOutputDTO, error) {

	now := time.Now()
	input.CreatedAt = now
	input.UpdatedAt = now

	err := uc.Repository.SaveProductDetail(ctx, &input.ProductDetail)

	if err != nil {
		return CreateProductDetailOutputDTO{}, err
	}

	outPut := CreateProductDetailOutputDTO{
		CreateProductDetailInputDTO: input,
	}

	return outPut, nil
}
