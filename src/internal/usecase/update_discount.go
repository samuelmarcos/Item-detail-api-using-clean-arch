package usecase

import (
	"context"
	"product_item_api/src/internal/entity"
)

// Remover CreateProductDetailInputDTO e usar ProductDetail sem o campo ID
type UpdateDiscountInputDTO struct {
	ProductID       string  `json:"product_id" validate:"required"`
	DiscountPercent float64 `json:"discount_percent" validate:"required,gte=0,lte=100"`
}

type UpdateDiscountUseCase struct {
	Repository entity.ProductDetailRepository
}

type UpdateProductDetailOutputDTO entity.ProductDetail

func NewUpdateDiscount(r entity.ProductDetailRepository) *UpdateDiscountUseCase {
	return &UpdateDiscountUseCase{
		Repository: r,
	}
}

func (uc *UpdateDiscountUseCase) Execute(ctx context.Context, input UpdateDiscountInputDTO) (UpdateProductDetailOutputDTO, error) {
	prod, err := uc.Repository.FindByProductID(ctx, input.ProductID)
	if err != nil {
		return UpdateProductDetailOutputDTO{}, err
	}

	prod.SetDiscount(input.DiscountPercent)

	err = uc.Repository.UpdateProductDetail(ctx, prod)
	if err != nil {
		return UpdateProductDetailOutputDTO{}, err
	}
	return UpdateProductDetailOutputDTO(*prod), nil

}
