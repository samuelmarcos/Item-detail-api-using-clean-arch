package usecase

import (
	"context"
	"product_item_api/src/internal/entity"
)

type ListProductDetailOutputDTO []entity.ProductDetail

type ListProductDetailUseCase struct {
	Repository entity.ProductDetailRepository
}

func NewListProductDetailUseCase(r entity.ProductDetailRepository) *ListProductDetailUseCase {
	return &ListProductDetailUseCase{
		Repository: r,
	}
}

func (uc *ListProductDetailUseCase) Execute(ctx context.Context) (ListProductDetailOutputDTO, error) {
	detailPtrs, err := uc.Repository.GetAllProductDetails(ctx)
	if err != nil {
		return nil, err
	}
	details := make(ListProductDetailOutputDTO, 0, len(detailPtrs))
	for _, ptr := range detailPtrs {
		if ptr != nil {
			details = append(details, *ptr)
		}
	}
	return details, nil
}
