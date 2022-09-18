package main

import (
	"log"
	"time"

	"go-playground/m/v1/src/adapters/controllers"
	"go-playground/m/v1/src/adapters/controllers/graphql/graph"
	"go-playground/m/v1/src/adapters/controllers/graphql/graph/generated"
	"go-playground/m/v1/src/adapters/controllers/http/middleware"
	"go-playground/m/v1/src/dependency"
	"go-playground/m/v1/src/infrastructure/driver"

	"github.com/labstack/echo/v4"

	gqlHandler "github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

func initRouter(e *echo.Echo, h controllers.AppController, srv *gqlHandler.Server) {
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

	{
		// GraphQL Handler
		e.GET("/graphql-playground", func(c echo.Context) error { // GUIからのGraphQL実行用（http://localhost:8444/graphql-playground）
			playground.Handler("GraphQL playground", "/query").ServeHTTP(c.Response(), c.Request())
			return nil
		})
		e.POST("/query", func(c echo.Context) error {
			srv.ServeHTTP(c.Response(), c.Request())
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

	srv := gqlHandler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	di := dependency.NewInjection(conn)

	initRouter(e, di.InitAppController(), srv)

	if err := e.Start(":8444"); err != nil {
		log.Fatal(err)
	}
}
