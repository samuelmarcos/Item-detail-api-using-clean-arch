package usecase

import (
	"context"
	"desafio_mercado_livre/src/internal/entity"
	"time"
)

type ProductDetailUseCase struct {
	Repository entity.ProductDetailRepository
}

func NewProductDetailUseCase(r entity.ProductDetailRepository) *ProductDetailUseCase {
	return &ProductDetailUseCase{
		Repository: r,
	}
}

func (uc *ProductDetailUseCase) GetProductDetail(ctx context.Context, id string) (*entity.ProductDetail, error) {
	return uc.Repository.GetProductDetail(ctx, id)
}

func (uc *ProductDetailUseCase) GetAllProductDetails(ctx context.Context) ([]*entity.ProductDetail, error) {
	return uc.Repository.GetAllProductDetails(ctx)
}

func (uc *ProductDetailUseCase) CreateProductDetail(ctx context.Context, product *entity.ProductDetail) error {
	now := time.Now()
	product.CreatedAt = now
	product.UpdatedAt = now

	return uc.Repository.SaveProductDetail(ctx, product)
}
