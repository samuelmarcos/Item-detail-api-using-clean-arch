package controller

import (
	"context"
	"fmt"
	"net/http"
	"product_item_api/src/internal/infra/logger"
	"product_item_api/src/internal/usecase"

	"github.com/gin-gonic/gin"
)

type UpdateDiscountUseCase interface {
	Execute(ctx context.Context, input usecase.UpdateDiscountInputDTO) (usecase.UpdateProductDetailOutputDTO, error)
}

func NewUpdateDiscountController(usecase UpdateDiscountUseCase, logger logger.Logger) *UpdateDiscountController {
	return &UpdateDiscountController{
		usecase: usecase,
		logger:  logger,
	}
}

// UpdateDiscount godoc
// @Summary      Update product discount
// @Description  Update the discount of a product
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        discount  body      usecase.UpdateDiscountInputDTO  true  "Discount Update Data"
// @Success      200  {object}  usecase.UpdateProductDetailOutputDTO
// @Failure      400  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /update/discount [put]
func (c *UpdateDiscountController) UpdateDiscount(ctx *gin.Context) {
	var input usecase.UpdateDiscountInputDTO
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid update discount data: " + err.Error()})
		return
	}

	output, err := c.usecase.Execute(ctx.Request.Context(), input)
	if err != nil {
		c.logger.Error(err.Error())
		ctx.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}
	inf := fmt.Sprintf("Product discount updated with success: %v", output)
	c.logger.Info(inf)
	ctx.JSON(http.StatusOK, output)
}
