package main

import (
	"github.com/labstack/echo/middleware"
	"github.com/leandropaes/goapi/routers"
)

func main() {
	e := routers.App

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Logger.Fatal(e.Start(":3000"))
}
