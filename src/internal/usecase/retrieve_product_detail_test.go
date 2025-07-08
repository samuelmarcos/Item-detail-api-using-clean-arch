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

func TestGetProductDetail(t *testing.T) {
	mockRepo := new(MockProductRepository)
	useCase := NewProductDetailUseCase(mockRepo)
	ctx := context.Background()

	t.Run("should get product successfully", func(t *testing.T) {
		product := createTestProduct()
		mockRepo.On("GetProductDetail", ctx, "MLB123456").Return(product, nil)

		result, err := useCase.GetProductDetail(ctx, "MLB123456")
		require.NoError(t, err)
		assert.Equal(t, product.ProductID, result.ProductID)
		assert.Equal(t, product.Name, result.Name)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should return error when product not found", func(t *testing.T) {
		mockRepo.On("GetProductDetail", ctx, "NONEXISTENT").Return(nil, errors.New("not found"))

		result, err := useCase.GetProductDetail(ctx, "NONEXISTENT")
		assert.Error(t, err)
		assert.Nil(t, result)
		mockRepo.AssertExpectations(t)
	})
}

func TestGetAllProductDetails(t *testing.T) {
	mockRepo := new(MockProductRepository)
	useCase := NewProductDetailUseCase(mockRepo)
	ctx := context.Background()

	t.Run("should get all products successfully", func(t *testing.T) {
		products := []*entity.ProductDetail{
			createTestProduct(),
			createTestProduct(),
		}
		products[1].ProductID = "MLB789012"
		products[1].Name = "Another Product"

		mockRepo.On("GetAllProductDetails", ctx).Return(products, nil)

		result, err := useCase.GetAllProductDetails(ctx)
		require.NoError(t, err)
		assert.Len(t, result, 2)
		assert.Equal(t, products[0].ProductID, result[0].ProductID)
		assert.Equal(t, products[1].ProductID, result[1].ProductID)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should return error when database fails", func(t *testing.T) {
		mockRepo = new(MockProductRepository)
		useCase = NewProductDetailUseCase(mockRepo)

		mockRepo.On("GetAllProductDetails", ctx).Return(nil, errors.New("database error"))

		result, err := useCase.GetAllProductDetails(ctx)
		assert.Error(t, err)
		assert.Nil(t, result)
		mockRepo.AssertExpectations(t)
	})
}

func TestCreateProductDetail(t *testing.T) {
	mockRepo := new(MockProductRepository)
	useCase := NewProductDetailUseCase(mockRepo)
	ctx := context.Background()

	t.Run("should create product successfully", func(t *testing.T) {
		product := createTestProduct()
		mockRepo.On("SaveProductDetail", ctx, mock.AnythingOfType("*entity.ProductDetail")).Return(nil)

		err := useCase.CreateProductDetail(ctx, product)
		require.NoError(t, err)
		assert.NotZero(t, product.CreatedAt)
		assert.NotZero(t, product.UpdatedAt)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should return error when product ID is missing", func(t *testing.T) {
		product := createTestProduct()
		product.ProductID = ""

		err := useCase.CreateProductDetail(ctx, product)
		assert.Equal(t, entity.ErrInvalidProductID, err)
	})

	t.Run("should return error when product name is missing", func(t *testing.T) {
		product := createTestProduct()
		product.Name = ""

		err := useCase.CreateProductDetail(ctx, product)
		assert.Equal(t, entity.ErrInvalidProductName, err)
	})

	t.Run("should return error when product price is invalid", func(t *testing.T) {
		product := createTestProduct()
		product.Price = 0

		err := useCase.CreateProductDetail(ctx, product)
		assert.Equal(t, entity.ErrInvalidProductPrice, err)
	})

	t.Run("should return error when product stock is invalid", func(t *testing.T) {
		product := createTestProduct()
		product.Stock = -1

		err := useCase.CreateProductDetail(ctx, product)
		assert.Equal(t, entity.ErrInvalidProductStock, err)
	})

	t.Run("should return error when seller name is missing", func(t *testing.T) {
		product := createTestProduct()
		product.Seller.Name = ""

		err := useCase.CreateProductDetail(ctx, product)
		assert.Equal(t, entity.ErrInvalidSellerName, err)
	})

	t.Run("should return error when repository fails", func(t *testing.T) {
		mockRepo := new(MockProductRepository)
		useCase := NewProductDetailUseCase(mockRepo)
		product := createTestProduct()
		mockRepo.On("SaveProductDetail", ctx, mock.AnythingOfType("*entity.ProductDetail")).Return(errors.New("database error"))

		err := useCase.CreateProductDetail(ctx, product)
		assert.Error(t, err)
		assert.Equal(t, "database error", err.Error())
		mockRepo.AssertExpectations(t)
	})
}
