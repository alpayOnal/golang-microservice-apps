package api

import (
	"github.com/labstack/echo"

	"micro_apps/micro_app/api/handlers"
)

func MainGroup(e *echo.Echo) {
	e.GET("/health-check", handlers.HealthCheck)
}

func AdminGroup(g *echo.Group) {
	g.GET("/", handlers.MainAdmin)
}
