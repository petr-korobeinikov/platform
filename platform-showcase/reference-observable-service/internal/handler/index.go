package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *IndexHandler) HandleRequest(c echo.Context) error {
	return c.String(http.StatusOK, "Hello!")
}

func NewIndexHandler() *IndexHandler {
	return &IndexHandler{}
}

type IndexHandler struct {
}
