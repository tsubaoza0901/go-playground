package main

import (
	"go-playground/m/v1/controllers"
	"go-playground/m/v1/controllers/rest/middleware"
	"go-playground/m/v1/injector"
	"log"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	middleware.InitMiddleware(e)

	dbConn := "*gorm.DB"
	d := injector.NewAppDependency(dbConn)

	appControllers := controllers.NewAppControllers()
	controllers.InitRouting(e, appControllers, d)

	log.Println(e.Start(":8444"))
}
