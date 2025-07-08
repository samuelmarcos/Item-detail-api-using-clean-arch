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
	mockUseCase.On("Execute", mock.Anything, input).Return(output, nil)
	logger.On("Info", mock.Anything, mock.Anything).Return()

	r := gin.Default()
	r.POST("/products", controller.CreateProductDetail)

	w := httptest.NewRecorder()
	body := bytes.NewBufferString(`{}`)
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

	input := usecase.CreateProductDetailInputDTO{}
	mockUseCase.On("Execute", mock.Anything, input).Return(usecase.CreateProductDetailOutputDTO{}, errors.New("internal error"))

	r := gin.Default()
	r.POST("/products", controller.CreateProductDetail)

	w := httptest.NewRecorder()
	body := bytes.NewBufferString(`{}`)
	req, _ := http.NewRequest("POST", "/products", body)
	req.Header.Set("Content-Type", "application/json")

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	mockUseCase.AssertExpectations(t)
}
