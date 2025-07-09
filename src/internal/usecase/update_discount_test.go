package usecase

import (
	"context"
	"errors"
	"testing"

	"product_item_api/src/internal/entity"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUpdateDiscount_Success(t *testing.T) {
	repo := new(MockProductRepository)
	ctx := context.Background()
	product := &entity.ProductDetail{ProductID: "PROD123", DiscountPercent: 0}
	repo.On("FindByProductID", ctx, "PROD123").Return(product, nil)
	repo.On("UpdateProductDetail", ctx, mock.AnythingOfType("*entity.ProductDetail")).Return(nil)

	uc := NewUpdateDiscount(repo)
	input := UpdateDiscountInputDTO{ProductID: "PROD123", DiscountPercent: 15.0}
	output, err := uc.Execute(ctx, input)
	assert.NoError(t, err)
	assert.Equal(t, input.ProductID, output.ProductID)
	assert.Equal(t, input.DiscountPercent, output.DiscountPercent)
	repo.AssertExpectations(t)
}

func TestUpdateDiscount_ProductNotFound(t *testing.T) {
	repo := new(MockProductRepository)
	ctx := context.Background()
	repo.On("FindByProductID", ctx, "PROD404").Return(nil, errors.New("not found"))

	uc := NewUpdateDiscount(repo)
	input := UpdateDiscountInputDTO{ProductID: "PROD404", DiscountPercent: 10.0}
	_, err := uc.Execute(ctx, input)
	assert.Error(t, err)
	repo.AssertExpectations(t)
}

func TestUpdateDiscount_UpdateError(t *testing.T) {
	repo := new(MockProductRepository)
	ctx := context.Background()
	product := &entity.ProductDetail{ProductID: "PROD123", DiscountPercent: 0}
	repo.On("FindByProductID", ctx, "PROD123").Return(product, nil)
	repo.On("UpdateProductDetail", ctx, mock.AnythingOfType("*entity.ProductDetail")).Return(errors.New("update failed"))

	uc := NewUpdateDiscount(repo)
	input := UpdateDiscountInputDTO{ProductID: "PROD123", DiscountPercent: 20.0}
	_, err := uc.Execute(ctx, input)
	assert.Error(t, err)
	repo.AssertExpectations(t)
}
