package usecase

import (
	"context"
	"errors"
	"testing"

	"product_item_api/src/internal/entity"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestListProductDetailUseCase_Success(t *testing.T) {
	repo := new(MockProductRepository)
	uc := NewListProductDetailUseCase(repo)
	entities := []*entity.ProductDetail{
		{ProductID: "1"},
		{ProductID: "2"},
	}
	repo.On("GetAllProductDetails", mock.Anything).Return(entities, nil)

	output, err := uc.Execute(context.Background())
	assert.NoError(t, err)
	assert.Len(t, output, 2)
	assert.Equal(t, "1", output[0].ProductID)
	assert.Equal(t, "2", output[1].ProductID)
	repo.AssertExpectations(t)
}

func TestListProductDetailUseCase_Error(t *testing.T) {
	repo := new(MockProductRepository)
	uc := NewListProductDetailUseCase(repo)
	repo.On("GetAllProductDetails", mock.Anything).Return(nil, errors.New("db error"))

	output, err := uc.Execute(context.Background())
	assert.Error(t, err)
	assert.Nil(t, output)
	repo.AssertExpectations(t)
}
