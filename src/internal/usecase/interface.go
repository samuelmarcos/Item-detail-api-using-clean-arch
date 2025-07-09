package usecase

import "context"

type CreateProductDetail interface {
	Execute(ctx context.Context, input CreateProductDetailInputDTO) (CreateProductDetailOutputDTO, error)
}

type GetProductDetail interface {
	Execute(ctx context.Context, id string) (GetProductDetailOutputDTO, error)
}

type ListProductDetail interface {
	Execute(ctx context.Context) (ListProductDetailOutputDTO, error)
}

type UpdateDiscount interface {
	Execute(ctx context.Context, input UpdateDiscountInputDTO) (UpdateProductDetailOutputDTO, error)
}
