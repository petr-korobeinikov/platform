package handler

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *ComplexHandler) HandleRequest(c echo.Context) error {
	var r complexHandlerResponse

	n := h.randomGenerator.Generate(c.Request().Context())

	r.Fib, r.Error = h.fibonacciCountingService.Count(c.Request().Context(), n)
	if r.Error != nil {
		return c.JSONPretty(http.StatusBadRequest, r, "  ")
	}

	return c.JSONPretty(http.StatusOK, r, "  ")
}

func NewComplexHandler(cfg ComplexHandlerCfg) *ComplexHandler {
	return &ComplexHandler{
		randomGenerator:          cfg.RandomGenerator,
		fibonacciCountingService: cfg.FibonacciCountingService,
	}
}

type (
	ComplexHandler struct {
		randomGenerator          randomGenerator
		fibonacciCountingService fibonacciCountingService
	}

	ComplexHandlerCfg struct {
		RandomGenerator          randomGenerator
		FibonacciCountingService fibonacciCountingService
	}

	fibonacciCountingService interface {
		Count(ctx context.Context, n int) (int, error)
	}

	randomGenerator interface {
		Generate(ctx context.Context) int
	}

	complexHandlerResponse struct {
		Error error `json:"error,omitempty"`
		Fib   int   `json:"fib"`
	}
)
