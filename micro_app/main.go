package main

import (
	"micro_apps/micro_app/config"
	"micro_apps/micro_app/router"
)


func main() {
	// create a new echo instance
	config.Load()
	e := router.New()
	e.Logger.Fatal(e.Start(":8000"))
}

