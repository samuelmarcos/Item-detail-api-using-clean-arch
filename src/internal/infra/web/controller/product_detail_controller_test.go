package controller

import (
	"bytes"
	"context"
	"desafio_mercado_livre/src/internal/entity"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type MockUseCase struct {
	mock.Mock
}

func (m *MockUseCase) GetProductDetail(ctx context.Context, id string) (*entity.ProductDetail, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.ProductDetail), args.Error(1)
}

func (m *MockUseCase) GetAllProductDetails(ctx context.Context) ([]*entity.ProductDetail, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*entity.ProductDetail), args.Error(1)
}

func (m *MockUseCase) CreateProductDetail(ctx context.Context, product *entity.ProductDetail) error {
	args := m.Called(ctx, product)
	return args.Error(0)
}

// MockLogger é um mock do Logger
type MockLogger struct {
	mock.Mock
}

func (m *MockLogger) Info(msg string, fields ...interface{}) {
	m.Called(msg, fields)
}

func (m *MockLogger) Debug(msg string, fields ...interface{}) {
	m.Called(msg, fields)
}

func (m *MockLogger) Error(msg string, fields ...interface{}) {
	m.Called(msg, fields)
}

func setupTestRouter(controller *ProductDetailController) *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.GET("/api/v1/products/:id", controller.GetProductDetail)
	router.GET("/api/v1/products", controller.GetAllProductDetails)
	router.POST("/api/v1/products", controller.CreateProductDetail)
	return router
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
	mockUseCase := new(MockUseCase)
	mockLogger := new(MockLogger)
	controller := NewProductDetailController(mockLogger, mockUseCase)
	router := setupTestRouter(controller)

	t.Run("should get product successfully", func(t *testing.T) {
		product := createTestProduct()
		mockUseCase.On("GetProductDetail", mock.Anything, "MLB123456").Return(product, nil)
		mockLogger.On("Info", mock.Anything, mock.Anything).Return()

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/v1/products/MLB123456", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		var response entity.ProductDetail
		err := json.Unmarshal(w.Body.Bytes(), &response)
		require.NoError(t, err)
		assert.Equal(t, product.ProductID, response.ProductID)
		assert.Equal(t, product.Name, response.Name)
	})

	t.Run("should return 404 when product not found", func(t *testing.T) {
		mockUseCase.On("GetProductDetail", mock.Anything, "NONEXISTENT").Return(nil, errors.New("not found"))

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/v1/products/NONEXISTENT", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)
		var response ErrorResponse
		err := json.Unmarshal(w.Body.Bytes(), &response)
		require.NoError(t, err)
		assert.Equal(t, "Product not found", response.Error)
	})
}

func TestGetAllProductDetails(t *testing.T) {
	mockUseCase := new(MockUseCase)
	mockLogger := new(MockLogger)
	controller := NewProductDetailController(mockLogger, mockUseCase)
	router := setupTestRouter(controller)

	t.Run("should get all products successfully", func(t *testing.T) {
		products := []*entity.ProductDetail{
			createTestProduct(),
			createTestProduct(),
		}
		products[1].ProductID = "MLB789012"
		products[1].Name = "Another Product"

		mockUseCase.On("GetAllProductDetails", mock.Anything).Return(products, nil)
		mockLogger.On("Info", mock.Anything, mock.Anything).Return()

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/v1/products", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		var response []*entity.ProductDetail
		err := json.Unmarshal(w.Body.Bytes(), &response)
		require.NoError(t, err)
		assert.Len(t, response, 2)
		assert.Equal(t, products[0].ProductID, response[0].ProductID)
		assert.Equal(t, products[1].ProductID, response[1].ProductID)
	})

	t.Run("should return 500 when error occurs", func(t *testing.T) {
		mockUseCase.On("GetAllProductDetails", mock.Anything).Return(nil, errors.New("database error"))

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/v1/products", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		var response ErrorResponse
		err := json.Unmarshal(w.Body.Bytes(), &response)
		require.NoError(t, err)
		assert.Equal(t, "Failed to fetch products", response.Error)
	})
}

func TestCreateProductDetail(t *testing.T) {
	mockUseCase := new(MockUseCase)
	mockLogger := new(MockLogger)
	controller := NewProductDetailController(mockLogger, mockUseCase)
	router := setupTestRouter(controller)

	t.Run("should create product successfully", func(t *testing.T) {
		product := createTestProduct()
		mockUseCase.On("CreateProductDetail", mock.Anything, mock.AnythingOfType("*entity.ProductDetail")).Return(nil)
		mockLogger.On("Info", mock.Anything, mock.Anything).Return()

		body, _ := json.Marshal(product)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/api/v1/products", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusCreated, w.Code)
		var response entity.ProductDetail
		err := json.Unmarshal(w.Body.Bytes(), &response)
		require.NoError(t, err)
		assert.Equal(t, product.ProductID, response.ProductID)
		assert.Equal(t, product.Name, response.Name)
	})

	t.Run("should return 400 when request body is invalid", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/api/v1/products", bytes.NewBufferString("invalid json"))
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		var response ErrorResponse
		err := json.Unmarshal(w.Body.Bytes(), &response)
		require.NoError(t, err)
		assert.Contains(t, response.Error, "Invalid product data")
	})

	t.Run("should return 400 when product ID is missing", func(t *testing.T) {
		product := createTestProduct()
		product.ProductID = ""
		mockUseCase.On("CreateProductDetail", mock.Anything, mock.AnythingOfType("*entity.ProductDetail")).Return(entity.ErrInvalidProductID)

		body, _ := json.Marshal(product)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/api/v1/products", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		var response ErrorResponse
		err := json.Unmarshal(w.Body.Bytes(), &response)
		require.NoError(t, err)
		assert.Equal(t, "Product ID is required", response.Error)
	})

	t.Run("should return 400 when product name is missing", func(t *testing.T) {
		product := createTestProduct()
		product.Name = ""
		mockUseCase.On("CreateProductDetail", mock.Anything, mock.AnythingOfType("*entity.ProductDetail")).Return(entity.ErrInvalidProductName)

		body, _ := json.Marshal(product)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/api/v1/products", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		var response ErrorResponse
		err := json.Unmarshal(w.Body.Bytes(), &response)
		require.NoError(t, err)
		assert.Equal(t, "Product name is required", response.Error)
	})

	t.Run("should return 400 when product price is invalid", func(t *testing.T) {
		product := createTestProduct()
		product.Price = 0
		mockUseCase.On("CreateProductDetail", mock.Anything, mock.AnythingOfType("*entity.ProductDetail")).Return(entity.ErrInvalidProductPrice)

		body, _ := json.Marshal(product)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/api/v1/products", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		var response ErrorResponse
		err := json.Unmarshal(w.Body.Bytes(), &response)
		require.NoError(t, err)
		assert.Equal(t, "Product price must be greater than zero", response.Error)
	})

	t.Run("should return 400 when product stock is invalid", func(t *testing.T) {
		product := createTestProduct()
		product.Stock = -1
		mockUseCase.On("CreateProductDetail", mock.Anything, mock.AnythingOfType("*entity.ProductDetail")).Return(entity.ErrInvalidProductStock)

		body, _ := json.Marshal(product)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/api/v1/products", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		var response ErrorResponse
		err := json.Unmarshal(w.Body.Bytes(), &response)
		require.NoError(t, err)
		assert.Equal(t, "Product stock cannot be negative", response.Error)
	})

	t.Run("should return 400 when seller name is missing", func(t *testing.T) {
		product := createTestProduct()
		product.Seller.Name = ""
		mockUseCase.On("CreateProductDetail", mock.Anything, mock.AnythingOfType("*entity.ProductDetail")).Return(entity.ErrInvalidSellerName)

		body, _ := json.Marshal(product)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/api/v1/products", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		var response ErrorResponse
		err := json.Unmarshal(w.Body.Bytes(), &response)
		require.NoError(t, err)
		assert.Equal(t, "Seller name is required", response.Error)
	})

	t.Run("should return 500 when unexpected error occurs", func(t *testing.T) {
		product := createTestProduct()
		mockUseCase.On("CreateProductDetail", mock.Anything, mock.AnythingOfType("*entity.ProductDetail")).Return(errors.New("unexpected error"))

		body, _ := json.Marshal(product)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/api/v1/products", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		var response ErrorResponse
		err := json.Unmarshal(w.Body.Bytes(), &response)
		require.NoError(t, err)
		assert.Contains(t, response.Error, "Failed to create product")
	})
}
