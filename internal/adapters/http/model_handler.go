// Package http for handing http request and responding with proper html
package http

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"www.github.com/fummbly/ai-dash/internal/domain"
	"www.github.com/fummbly/ai-dash/internal/service"
)

type ModelHandler struct {
	modelService *service.ModelService
}

func NewModelHandler(svr service.ModelService) *ModelHandler {
	return &ModelHandler{
		modelService: &svr,
	}
}

func (h *ModelHandler) ListModels(c echo.Context) error {
	models, err := h.modelService.List()
	if len(models.Models) == 0 {
		return fmt.Errorf("no models were received")
	}

	if err != nil {
		fmt.Printf("failed to get models data from api: %v", err)

		return c.Render(http.StatusInternalServerError, "index", domain.Models{})
	}

	return c.Render(http.StatusOK, "response", models)
}
