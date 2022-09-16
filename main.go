package main

import (
	"log"
	"time"

	"go-playground/m/v1/src/adapters/controllers/http/middleware"
	"go-playground/m/v1/src/adapters/controllers/http/router"
	"go-playground/m/v1/src/dependency"
	"go-playground/m/v1/src/infrastructure/driver"

	"github.com/labstack/echo/v4"
)

const location = "Asia/Tokyo"

// SetTimeZone ...
func SetTimeZone() {
	loc, err := time.LoadLocation(location)
	if err != nil {
		loc = time.FixedZone(location, 9*60*60)
	}
	time.Local = loc
}

func main() {
	SetTimeZone()

	conn := driver.InitDBConn()
	sqlDB, err := conn.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer sqlDB.Close()

	e := echo.New()
	middleware.InitMiddleware(e)

	di := dependency.NewInjection(conn)

	router.InitRouting(e, di.InitHTTPHandler())

	if err := e.Start(":8444"); err != nil {
		log.Fatal(err)
	}
}
