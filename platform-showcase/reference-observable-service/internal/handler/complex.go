package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *ComplexHandler) HandleRequest(c echo.Context) error {
	return c.String(http.StatusOK, "I am complex handler!")
}

func NewComplexHandler() *ComplexHandler {
	return &ComplexHandler{}
}

type ComplexHandler struct {
}
