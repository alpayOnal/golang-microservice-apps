package main

import (
	"micro_apps/micro_app/router"
)


func main() {
	e := router.New()
	e.Logger.Fatal(e.Start(":8000"))
}
