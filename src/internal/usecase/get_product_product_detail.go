package usecase

import (
	"context"
	"product_item_api/src/internal/entity"
)

type GetProductDetailOutputInput string

type GetProductDetailOutputDTO entity.ProductDetail

type GetProductDetailUseCase struct {
	Repository entity.ProductDetailRepository
}

func NewGetProductDetailUseCase(r entity.ProductDetailRepository) *GetProductDetailUseCase {
	return &GetProductDetailUseCase{
		Repository: r,
	}
}

func (uc *GetProductDetailUseCase) Execute(ctx context.Context, id string) (GetProductDetailOutputDTO, error) {
	detail, err := uc.Repository.GetProductDetail(ctx, id)
	if err != nil {
		return GetProductDetailOutputDTO{}, err
	}
	if detail == nil {
		return GetProductDetailOutputDTO{}, nil
	}
	return GetProductDetailOutputDTO(*detail), nil
}
