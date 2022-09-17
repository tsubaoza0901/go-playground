package main

import (
	"log"
	"time"

	"go-playground/m/v1/src/adapters/controllers"
	"go-playground/m/v1/src/adapters/controllers/http/middleware"
	"go-playground/m/v1/src/dependency"
	"go-playground/m/v1/src/infrastructure/driver"

	"github.com/labstack/echo/v4"
)

func initRouter(e *echo.Echo, h controllers.AppController) {
	{
		// User Handler
		e.POST("/user", h.IUserHandler.CreateNewUser)
		e.GET("/user/:id", h.IUserHandler.GetUser)
		e.GET("/users", h.IUserHandler.GetUserList)
	}

	{
		// Grade Handler
		e.GET("/grades", h.IGradeHandler.GetGradeList)
	}

	{
		// DealHistory Handler
		e.GET("/dealHistories/:userId", h.IDealHistoryHandler.GetDealHistoryList)
	}

	{
		// BalanceControl Handler
		e.PUT("/pay/:userId", h.IBalanceControlHandler.Pay)
		e.PUT("/topup/:userId", h.IBalanceControlHandler.TopUp)
		e.GET("/remainingBalance/:userId", h.IBalanceControlHandler.GetRemainingBalance)
	}
}

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

	initRouter(e, di.InitAppController())

	if err := e.Start(":8444"); err != nil {
		log.Fatal(err)
	}
}
