package main

import (
	"go-playground/m/v1/src/http/router"
	"log"

	"go-playground/m/v1/src/http/middleware"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	middleware.InitMiddleware(e)

	router.InitRouting(e)

	log.Println(e.Start(":8444"))
}
