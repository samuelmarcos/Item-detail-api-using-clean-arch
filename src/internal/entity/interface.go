package entity

import (
	"context"
)

type ProductDetailRepository interface {
	GetProductDetail(ctx context.Context, id string) (*ProductDetail, error)
	GetAllProductDetails(ctx context.Context) ([]*ProductDetail, error)
	SaveProductDetail(ctx context.Context, product *ProductDetail) error
	FindByProductID(ctx context.Context, productID string) (*ProductDetail, error)
	UpdateProductDetail(ctx context.Context, product *ProductDetail) error
}
