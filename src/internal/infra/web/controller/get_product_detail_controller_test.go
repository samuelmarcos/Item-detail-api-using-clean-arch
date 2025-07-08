package controller

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"product_item_api/src/internal/entity"
	"product_item_api/src/internal/usecase"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockGetProductDetailUseCase struct {
	mock.Mock
}

func (m *mockGetProductDetailUseCase) Execute(ctx context.Context, id string) (usecase.GetProductDetailOutputDTO, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(usecase.GetProductDetailOutputDTO), args.Error(1)
}

type mockLogger struct {
	mock.Mock
}

func (m *mockLogger) Info(msg string, fields ...interface{}) {
	m.Called(msg, fields)
}

func (m *mockLogger) Debug(msg string, fields ...interface{}) {
	m.Called(msg, fields)
}

func (m *mockLogger) Error(msg string, fields ...interface{}) {
	m.Called(msg, fields)
}

func TestGetProductDetailController_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockUseCase := new(mockGetProductDetailUseCase)
	logger := new(mockLogger)
	controller := NewGetProductDetailController(mockUseCase, logger)

	output := usecase.GetProductDetailOutputDTO{}
	mockUseCase.On("Execute", mock.Anything, "123").Return(output, nil)
	logger.On("Info", mock.Anything, mock.Anything).Return()

	r := gin.Default()
	r.GET("/products/:id", controller.GetProductDetail)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/products/123", nil)

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	mockUseCase.AssertExpectations(t)
	logger.AssertExpectations(t)
}

func TestGetProductDetailController_BadRequest(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockUseCase := new(mockGetProductDetailUseCase)
	logger := new(mockLogger)
	controller := NewGetProductDetailController(mockUseCase, logger)

	r := gin.Default()
	r.GET("/products/:id", controller.GetProductDetail)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/products/", nil)

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestGetProductDetailController_NotFound(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockUseCase := new(mockGetProductDetailUseCase)
	logger := new(mockLogger)
	controller := NewGetProductDetailController(mockUseCase, logger)

	mockUseCase.On("Execute", mock.Anything, "notfound").Return(usecase.GetProductDetailOutputDTO{}, errors.New("not found"))

	r := gin.Default()
	r.GET("/products/:id", controller.GetProductDetail)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/products/notfound", nil)

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	mockUseCase.AssertExpectations(t)
}

func TestGetProductDetailController_ValidationError(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockUseCase := new(mockGetProductDetailUseCase)
	logger := new(mockLogger)
	controller := NewGetProductDetailController(mockUseCase, logger)

	mockUseCase.On("Execute", mock.Anything, "badid").Return(usecase.GetProductDetailOutputDTO{}, entity.ErrInvalidProductID)

	r := gin.Default()
	r.GET("/products/:id", controller.GetProductDetail)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/products/badid", nil)

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	mockUseCase.AssertExpectations(t)
}
