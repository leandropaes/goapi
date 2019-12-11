package main

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo/middleware"
	"github.com/leandropaes/goapi/routers"
	"log"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	e := routers.App

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Logger.Fatal(e.Start(":3000"))
}
