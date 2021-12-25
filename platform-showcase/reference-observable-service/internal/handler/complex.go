package handler

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *ComplexHandler) HandleRequest(c echo.Context) error {
	var r complexHandlerResponse

	r.Fib, r.Error = h.fibonacciCountingService.Count(c.Request().Context(), 10)
	if r.Error != nil {
		return c.JSONPretty(http.StatusBadRequest, r, "  ")
	}

	return c.JSONPretty(http.StatusOK, r, "  ")
}

func NewComplexHandler(cfg ComplexHandlerCfg) *ComplexHandler {
	return &ComplexHandler{
		fibonacciCountingService: cfg.FibonacciCountingService,
	}
}

type (
	ComplexHandler struct {
		fibonacciCountingService fibonacciCountingService
	}

	ComplexHandlerCfg struct {
		FibonacciCountingService fibonacciCountingService
	}

	fibonacciCountingService interface {
		Count(ctx context.Context, n int) (int, error)
	}

	complexHandlerResponse struct {
		Error error `json:"error,omitempty"`
		Fib   int   `json:"fib"`
	}
)
