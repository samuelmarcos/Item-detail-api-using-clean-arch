package controller

import (
	"bytes"
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"product_item_api/src/internal/usecase"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockCreateProductDetailUseCase struct {
	mock.Mock
}

func (m *mockCreateProductDetailUseCase) Execute(ctx context.Context, input usecase.CreateProductDetailInputDTO) (usecase.CreateProductDetailOutputDTO, error) {
	args := m.Called(ctx, input)
	return args.Get(0).(usecase.CreateProductDetailOutputDTO), args.Error(1)
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
func TestCreateProductDetailController_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockUseCase := new(mockCreateProductDetailUseCase)
	logger := new(MockLogger)
	controller := NewCreateProductDetailController(mockUseCase, logger)

	input := usecase.CreateProductDetailInputDTO{}
	output := usecase.CreateProductDetailOutputDTO{CreateProductDetailInputDTO: input}
	mockUseCase.On("Execute", mock.Anything, mock.Anything).Return(output, nil)
	logger.On("Info", mock.Anything).Maybe().Return()
	logger.On("Info", mock.Anything, mock.Anything).Maybe().Return()

	r := gin.Default()
	r.POST("/products", controller.CreateProductDetail)

	w := httptest.NewRecorder()
	body := bytes.NewBufferString(`{
    "product_id": "MLB123456",
    "name": "Test Product",
    "description": "Test Description",
    "brand": "Test Brand",
    "model": "Test Model",
    "color": "Test Color",
    "category": "Test Category",
    "images": ["http://example.com/image1.jpg"],
    "price": 100.00,
    "original_price": 120.00,
    "discount_percent": 16.67,
    "stock": 10,
    "seller": {
        "name": "Test Seller",
        "type": "Official",
        "reputation": "Good",
        "sales": 100,
        "official": true
    },
    "warranty": "1 year",
    "payment_options": ["Credit Card", "Boleto"],
    "specs": {"attributes": {"Size": "Medium", "Color": "Blue"}},
    "related_products": ["MLB789012"],
    "rating": 4.5,
    "review_count": 100,
    "free_shipping": true,
    "purchase_options": ["New"],
    "highlights": ["Best Seller"]
}`)
	req, _ := http.NewRequest("POST", "/products", body)
	req.Header.Set("Content-Type", "application/json")

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	mockUseCase.AssertExpectations(t)
	logger.AssertExpectations(t)
}

func TestCreateProductDetailController_BadRequest(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockUseCase := new(mockCreateProductDetailUseCase)
	logger := new(MockLogger)
	controller := NewCreateProductDetailController(mockUseCase, logger)

	r := gin.Default()
	r.POST("/products", controller.CreateProductDetail)

	w := httptest.NewRecorder()
	body := bytes.NewBufferString(`invalid-json`)
	req, _ := http.NewRequest("POST", "/products", body)
	req.Header.Set("Content-Type", "application/json")

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestCreateProductDetailController_InternalError(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockUseCase := new(mockCreateProductDetailUseCase)
	logger := new(MockLogger)
	controller := NewCreateProductDetailController(mockUseCase, logger)

	t.Run("should return 500 on usecase error", func(t *testing.T) {
		mockUseCase.On("Execute", mock.Anything, mock.Anything).Return(usecase.CreateProductDetailOutputDTO{}, errors.New("internal error"))
		logger.On("Error", mock.Anything).Maybe().Return()
		logger.On("Info", mock.Anything).Maybe().Return()
		logger.On("Info", mock.Anything, mock.Anything).Maybe().Return()

		r := gin.Default()
		r.POST("/products", controller.CreateProductDetail)

		w := httptest.NewRecorder()
		body := bytes.NewBufferString(`{
			"product_id": "MLB123456",
			"name": "Test Product",
			"description": "Test Description",
			"brand": "Test Brand",
			"model": "Test Model",
			"color": "Test Color",
			"category": "Test Category",
			"images": ["http://example.com/image1.jpg"],
			"price": 100.00,
			"original_price": 120.00,
			"discount_percent": 16.67,
			"stock": 10,
			"seller": {
				"name": "Test Seller",
				"type": "Official",
				"reputation": "Good",
				"sales": 100,
				"official": true
			},
			"warranty": "1 year",
			"payment_options": ["Credit Card", "Boleto"],
			"specs": {"attributes": {"Size": "Medium", "Color": "Blue"}},
			"related_products": ["MLB789012"],
			"rating": 4.5,
			"review_count": 100,
			"free_shipping": true,
			"purchase_options": ["New"],
			"highlights": ["Best Seller"]
		}`)
		req, _ := http.NewRequest("POST", "/products", body)
		req.Header.Set("Content-Type", "application/json")

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		mockUseCase.AssertExpectations(t)
	})
}
