package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

type AdminHandler struct {
}

func NewAdminHandler(e *echo.Echo) {
	handler := &AdminHandler{}

	e.GET("/", handler.MainAdmin)

}

func (h *AdminHandler) MainAdmin(c echo.Context) error {
	return c.String(http.StatusOK, " You are on the Admin Page !!!")
}
