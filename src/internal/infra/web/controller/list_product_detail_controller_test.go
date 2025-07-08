package controller

import (
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

type mockListProductDetailUseCase struct {
	mock.Mock
}

func (m *mockListProductDetailUseCase) Execute(ctx context.Context) (usecase.ListProductDetailOutputDTO, error) {
	args := m.Called(ctx)
	return args.Get(0).(usecase.ListProductDetailOutputDTO), args.Error(1)
}

func TestListDetailController_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockUseCase := new(mockListProductDetailUseCase)
	logger := new(MockLogger)
	controller := NewListProductDetailController(mockUseCase, logger)

	output := usecase.ListProductDetailOutputDTO{}
	mockUseCase.On("Execute", mock.Anything).Return(output, nil)
	logger.On("Info", mock.Anything, mock.Anything).Return()

	r := gin.Default()
	r.GET("/products", controller.GetAllProductDetails)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/products", nil)

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	mockUseCase.AssertExpectations(t)
	logger.AssertExpectations(t)
}

func TestListDetailController_InternalError(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockUseCase := new(mockListProductDetailUseCase)
	logger := new(MockLogger)
	controller := NewListProductDetailController(mockUseCase, logger)

	mockUseCase.On("Execute", mock.Anything).Return(usecase.ListProductDetailOutputDTO{}, errors.New("internal error"))

	r := gin.Default()
	r.GET("/products", controller.GetAllProductDetails)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/products", nil)

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	mockUseCase.AssertExpectations(t)
}
