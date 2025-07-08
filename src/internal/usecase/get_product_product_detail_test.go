package usecase

import (
	"context"
	"errors"
	"testing"

	"product_item_api/src/internal/entity"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetProductDetailUseCase_Success(t *testing.T) {
	repo := new(MockProductRepository)
	uc := NewGetProductDetailUseCase(repo)
	entityDetail := &entity.ProductDetail{ProductID: "1"}
	repo.On("GetProductDetail", mock.Anything, "1").Return(entityDetail, nil)

	output, err := uc.Execute(context.Background(), "1")
	assert.NoError(t, err)
	assert.Equal(t, "1", output.ProductID)
	repo.AssertExpectations(t)
}

func TestGetProductDetailUseCase_Error(t *testing.T) {
	repo := new(MockProductRepository)
	uc := NewGetProductDetailUseCase(repo)
	repo.On("GetProductDetail", mock.Anything, "2").Return(nil, errors.New("db error"))

	output, err := uc.Execute(context.Background(), "2")
	assert.Error(t, err)
	assert.Equal(t, GetProductDetailOutputDTO{}, output)
	repo.AssertExpectations(t)
}

func TestGetProductDetailUseCase_NotFound(t *testing.T) {
	repo := new(MockProductRepository)
	uc := NewGetProductDetailUseCase(repo)
	repo.On("GetProductDetail", mock.Anything, "3").Return(nil, nil)

	output, err := uc.Execute(context.Background(), "3")
	assert.NoError(t, err)
	assert.Equal(t, GetProductDetailOutputDTO{}, output)
	repo.AssertExpectations(t)
}
