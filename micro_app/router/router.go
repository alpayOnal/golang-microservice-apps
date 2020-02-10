package router

import (
	"github.com/labstack/echo"

	"micro_apps/micro_app/api/handlers"
	"micro_apps/micro_app/api/middlewares"
)

func New() *echo.Echo {
	// create a new echo instance
	e := echo.New()

	//create groups
	adminGroup := e.Group("/admin")

	//set all middlewares
	middlewares.SetMainMiddleWares(e)
	middlewares.SetAdminMiddlewares(adminGroup)

	handlers.NewItemHandler(e)
	handlers.NewAdminHandler(e)
	handlers.NewHealthCheckHandler(e)
	return e
}
