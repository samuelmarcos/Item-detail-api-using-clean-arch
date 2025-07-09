package controller

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"product_item_api/src/internal/usecase"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockUpdateDiscountUseCase struct {
	mock.Mock
}

func (m *mockUpdateDiscountUseCase) Execute(ctx context.Context, input usecase.UpdateDiscountInputDTO) (usecase.UpdateProductDetailOutputDTO, error) {
	args := m.Called(ctx, input)
	return args.Get(0).(usecase.UpdateProductDetailOutputDTO), args.Error(1)
}

func TestUpdateDiscountController_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	mockUseCase := new(mockUpdateDiscountUseCase)
	mockLogger := new(MockLogger)
	controller := NewUpdateDiscountController(mockUseCase, mockLogger)

	input := usecase.UpdateDiscountInputDTO{ProductID: "PROD123", DiscountPercent: 10.0}
	output := usecase.UpdateProductDetailOutputDTO{ProductID: "PROD123", DiscountPercent: 10.0}
	mockUseCase.On("Execute", mock.Anything, input).Return(output, nil)
	mockLogger.On("Info", mock.Anything, mock.Anything).Maybe().Return()
	mockLogger.On("Info", mock.Anything).Maybe().Return()
	mockLogger.On("Error", mock.Anything, mock.Anything).Maybe().Return()
	mockLogger.On("Error", mock.Anything).Maybe().Return()

	r.PUT("/update/discount", controller.UpdateDiscount)
	body, _ := json.Marshal(input)
	req, _ := http.NewRequest(http.MethodPut, "/update/discount", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var resp usecase.UpdateProductDetailOutputDTO
	json.Unmarshal(w.Body.Bytes(), &resp)
	assert.Equal(t, output.ProductID, resp.ProductID)
	assert.Equal(t, output.DiscountPercent, resp.DiscountPercent)
	mockUseCase.AssertExpectations(t)
}

func TestUpdateDiscountController_BindError(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	mockUseCase := new(mockUpdateDiscountUseCase)
	mockLogger := new(MockLogger)
	controller := NewUpdateDiscountController(mockUseCase, mockLogger)

	r.PUT("/update/discount", controller.UpdateDiscount)
	body := []byte(`{"product_id":123}`) // DiscountPercent faltando e product_id tipo errado
	req, _ := http.NewRequest(http.MethodPut, "/update/discount", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestUpdateDiscountController_UseCaseError(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	mockUseCase := new(mockUpdateDiscountUseCase)
	mockLogger := new(MockLogger)
	controller := NewUpdateDiscountController(mockUseCase, mockLogger)

	input := usecase.UpdateDiscountInputDTO{ProductID: "PROD123", DiscountPercent: 10.0}
	mockUseCase.On("Execute", mock.Anything, input).Return(usecase.UpdateProductDetailOutputDTO{}, errors.New("update failed"))

	r.PUT("/update/discount", controller.UpdateDiscount)
	body, _ := json.Marshal(input)
	req, _ := http.NewRequest(http.MethodPut, "/update/discount", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	mockUseCase.AssertExpectations(t)
}
