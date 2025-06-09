package entity

import (
	"context"
)

type ProductDetailRepository interface {
	GetProductDetail(ctx context.Context, id string) (*ProductDetail, error)
	GetAllProductDetails(ctx context.Context) ([]*ProductDetail, error)
	SaveProductDetail(ctx context.Context, product *ProductDetail) error
}
