package api

import (
	"github.com/labstack/echo"

	"micro_apps/micro_app/handlers"
)

func MainGroup(e *echo.Echo) {
	e.GET("/health-check", handlers.HealthCheck)

	//e.GET("/items/", handlers.GetItems)
	e.GET("/items/:id", handlers.GetItem)
	e.GET("/items", handlers.GetItems)

	e.POST("/items", handlers.AddItem)

}

func AdminGroup(g *echo.Group) {
	g.GET("/", handlers.MainAdmin)
}
