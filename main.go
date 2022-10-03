package main

import (
	"log"

	"go-playground/m/v1/src/controllers"
	"go-playground/m/v1/src/controllers/middleware"
	"go-playground/m/v1/src/injector"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	middleware.InitMiddleware(e)

	di := injector.NewDependency()

	handlers := di.InitAppHandlers()
	controllers.InitRouting(e, handlers)

	log.Println(e.Start(":8444"))
}
