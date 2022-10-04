package main

import (
	"log"

	"go-playground/m/v1/src/controllers"
	"go-playground/m/v1/src/controllers/rest/middleware"
	"go-playground/m/v1/src/injector"

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
