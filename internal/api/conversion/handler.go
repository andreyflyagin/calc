package conversion

import (
	"calc/internal/service"
	"context"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/shopspring/decimal"
)

type handler struct {
	conversionService service.ConversionService
}

// NewHandler is a constructor for handler
func NewHandler(us service.ConversionService) func(ctx *fiber.Ctx) error {

	handler := &handler{
		conversionService: us,
	}
	return handler.getConversion
}

// Convert godoc
// @Summary get conversion
// @Description get conversion
// @Tags convert
// @Accept */*
// @Produce json
// @Param   from      query     string     false  "symbol base"       minlength(2)  maxlength(50)
// @Param   to        query     string     false  "symbol quote"       minlength(2)  maxlength(50)
// @Param   amount    query     number     false  "amount"
// @Success 200 {object} model.Conversion
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/convert [get]
func (h *handler) getConversion(ctx *fiber.Ctx) error {
	customContext, cancel := context.WithCancel(ctx.Context())
	defer cancel()

	from := ctx.Query("from")
	if from == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": "Please specify a valid from!",
		})
	}

	to := ctx.Query("to")
	if from == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": "Please specify a valid to!",
		})
	}

	amount, err := decimal.NewFromString(ctx.Query("amount"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": "Please specify a valid amount!",
		})
	}

	value, err := h.conversionService.GetConversion(customContext, from, to, amount)
	if errors.Is(err, service.ErrorConversionNotFound) {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fromConversionToRepo(value))
}
