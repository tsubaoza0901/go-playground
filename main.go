package main

import (
	"log"
	"time"

	"go-playground/m/v1/src/adapters/controllers"
	"go-playground/m/v1/src/adapters/controllers/http/middleware"
	"go-playground/m/v1/src/dependency"
	"go-playground/m/v1/src/infrastructure/driver"

	"github.com/labstack/echo/v4"

	"github.com/99designs/gqlgen/graphql/playground"
)

func initRouter(e *echo.Echo, ctr controllers.AppController) {
	{
		// User Handler
		e.POST("/user", ctr.IUserHandler.CreateNewUser)
		e.GET("/user/:id", ctr.IUserHandler.GetUser)
		e.GET("/users", ctr.IUserHandler.GetUserList)
	}

	{
		// Grade Handler
		e.GET("/grades", ctr.IGradeHandler.GetGradeList)
	}

	{
		// DealHistory Handler
		e.GET("/dealHistories/:userId", ctr.IDealHistoryHandler.GetDealHistoryList)
	}

	{
		// BalanceControl Handler
		e.PUT("/pay/:userId", ctr.IBalanceControlHandler.Pay)
		e.PUT("/topup/:userId", ctr.IBalanceControlHandler.TopUp)
		e.GET("/remainingBalance/:userId", ctr.IBalanceControlHandler.GetRemainingBalance)
	}

	{
		// GraphQL Handler
		e.GET("/graphql-playground", func(c echo.Context) error { // GUIからのGraphQL実行用（http://localhost:8444/graphql-playground）
			playground.Handler("GraphQL playground", "/query").ServeHTTP(c.Response(), c.Request())
			return nil
		})
		e.POST("/query", func(c echo.Context) error {
			ctr.Server.ServeHTTP(c.Response(), c.Request())
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
