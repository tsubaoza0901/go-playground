package main

import (
	"go-playground/m/v1/src/http/middleware"
	"go-playground/m/v1/src/http/router"
	"log"
	"time"

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

	e := echo.New()

	middleware.InitMiddleware(e)
	jwtConfig := middleware.NewJwtConfig()

	router.InitRouting(e, jwtConfig)

	log.Println(e.Start(":8444"))
}
