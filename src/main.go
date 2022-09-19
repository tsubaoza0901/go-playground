package main

import (
	"log"
	"time"

	"go-playground/m/v1/adapters/controllers"
	"go-playground/m/v1/adapters/controllers/rest/middleware"
	"go-playground/m/v1/dependency"
	"go-playground/m/v1/infrastructure/driver"

	"github.com/labstack/echo/v4"

	"github.com/99designs/gqlgen/graphql/playground"
)

func initRouter(e *echo.Echo, appCtr controllers.AppController) {
	{
		// User Handler
		e.POST("/user", appCtr.UserHandler.CreateNewUser)
		e.GET("/user/:id", appCtr.UserHandler.GetUser)
		e.GET("/users", appCtr.UserHandler.GetUserList)
	}

	{
		// Grade Handler
		e.GET("/grades", appCtr.GradeHandler.GetGradeList)
	}

	{
		// DealHistory Handler
		e.GET("/dealHistories/:userId", appCtr.DealHistoryHandler.GetDealHistoryList)
	}

	{
		// BalanceControl Handler
		e.PUT("/pay/:userId", appCtr.BalanceControlHandler.Pay)
		e.PUT("/topup/:userId", appCtr.BalanceControlHandler.TopUp)
		e.GET("/remainingBalance/:userId", appCtr.BalanceControlHandler.GetRemainingBalance)
	}

	{
		// GraphQL Handler
		e.GET("/graphql-playground", func(c echo.Context) error { // GUIからのGraphQL実行用（http://localhost:8444/graphql-playground）
			playground.Handler("GraphQL playground", "/query").ServeHTTP(c.Response(), c.Request())
			return nil
		})
		e.POST("/query", func(c echo.Context) error {
			appCtr.Server.ServeHTTP(c.Response(), c.Request())
			return nil
		})
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

	db := driver.InitDB()
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer sqlDB.Close()

	e := echo.New()
	middleware.InitMiddleware(e)

	di := dependency.NewInjection(db)

	initRouter(e, di.InitAppController())

	if err := e.Start(":8444"); err != nil {
		log.Fatal(err)
	}
}
