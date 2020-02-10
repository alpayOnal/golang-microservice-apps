package handlers

import (
	"net/http"

	"github.com/labstack/echo"

	"micro_apps/micro_app/models"
)

type HealthCheckHandler struct {
}

func NewHealthCheckHandler(e *echo.Echo) {
	handler := &AdminHandler{}

	e.GET("/", handler.MainAdmin)

}

func (h *HealthCheckHandler) HealthCheck(c echo.Context) error {

	resp := models.HealthCheckResponse{
		Message: "Everything is good!",
	}
	return c.JSON(http.StatusOK, resp)
}
