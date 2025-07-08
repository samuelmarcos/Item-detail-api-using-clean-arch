package usecase

import (
	"context"
	"errors"
	"product_item_api/src/internal/entity"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type MockProductRepository struct {
	mock.Mock
}

func (m *MockProductRepository) SaveProductDetail(ctx context.Context, product *entity.ProductDetail) error {
	args := m.Called(ctx, product)
	return args.Error(0)
}

func (m *MockProductRepository) GetProductDetail(ctx context.Context, id string) (*entity.ProductDetail, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.ProductDetail), args.Error(1)
}

func (m *MockProductRepository) GetAllProductDetails(ctx context.Context) ([]*entity.ProductDetail, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*entity.ProductDetail), args.Error(1)
}

func createTestProduct() *entity.ProductDetail {
	now := time.Now()
	return &entity.ProductDetail{
		ProductID:       "MLB123456",
		Name:            "Test Product",
		Description:     "Test Description",
		Brand:           "Test Brand",
		Model:           "Test Model",
		Color:           "Test Color",
		Category:        "Test Category",
		Images:          []string{"http://example.com/image1.jpg"},
		Price:           100.00,
		OriginalPrice:   120.00,
		DiscountPercent: 16.67,
		Stock:           10,
		Seller: entity.SellerInfo{
			Name:       "Test Seller",
			Type:       "Official",
			Reputation: "Good",
			Sales:      100,
			Official:   true,
		},
		Warranty:       "1 year",
		PaymentOptions: []string{"Credit Card", "Boleto"},
		Specs: entity.ProductSpecs{
			Attributes: map[string]string{
				"Size":  "Medium",
				"Color": "Blue",
			},
		},
		RelatedProducts: []string{"MLB789012"},
		Rating:          4.5,
		ReviewCount:     100,
		FreeShipping:    true,
		PurchaseOptions: []string{"New"},
		Highlights:      []string{"Best Seller"},
		CreatedAt:       now,
		UpdatedAt:       now,
	}
}

func TestCreateProductDetail(t *testing.T) {
	mockRepo := new(MockProductRepository)
	useCase := NewCreateProductDetailUseCase(mockRepo)
	ctx := context.Background()

	t.Run("should create product successfully", func(t *testing.T) {
		product := createTestProduct()
		input := CreateProductDetailInputDTO{ProductDetail: *product}
		mockRepo.On("SaveProductDetail", ctx, mock.AnythingOfType("*entity.ProductDetail")).Return(nil)

		output, err := useCase.Execute(ctx, input)
		require.NoError(t, err)
		assert.NotZero(t, output.CreateProductDetailInputDTO.CreatedAt)
		assert.NotZero(t, output.CreateProductDetailInputDTO.UpdatedAt)
		assert.Equal(t, product.ProductID, output.CreateProductDetailInputDTO.ProductID)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should return error when repository fails", func(t *testing.T) {
		mockRepo := new(MockProductRepository)
		useCase := NewCreateProductDetailUseCase(mockRepo)
		product := createTestProduct()
		input := CreateProductDetailInputDTO{ProductDetail: *product}
		mockRepo.On("SaveProductDetail", ctx, mock.AnythingOfType("*entity.ProductDetail")).Return(errors.New("database error"))

		output, err := useCase.Execute(ctx, input)
		assert.Error(t, err)
		assert.Equal(t, "database error", err.Error())
		assert.Equal(t, CreateProductDetailOutputDTO{}, output)
		mockRepo.AssertExpectations(t)
	})
}
