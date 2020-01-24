package router

import (

	"github.com/labstack/echo"

	"micro_apps/micro_app/api"
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

	//set main routes
	api.MainGroup(e)

	//set groupRoutes
	api.AdminGroup(adminGroup)
	return e
}