package web

import (
	"go-playground/m/v1/framework/web/rest"

	"github.com/labstack/echo/v4"
)

// AppHandler ...
type AppHandler struct {
	*rest.BalanceControlHandler
	*rest.DealHistoryHandler
	*rest.GradeHandler
	*rest.UserHandler
}

func (h *AppHandler) initRouter(e *echo.Echo) {
	{
		// User Handler
		e.POST("/user", h.CreateNewUser)
		e.PUT("/user/:id", h.UpdateUser)
		e.GET("/user/:id", h.GetUser)
		e.GET("/users", h.GetUserList)
	}

	{
		// Grade Handler
		e.GET("/grades", h.GetGradeList)
	}

	{
		// DealHistory Handler
		e.GET("/dealHistories/:userId", h.GetDealHistoryList)
	}

	{
		// BalanceControl Handler
		e.PUT("/pay/:userId", h.Pay)
		e.PUT("/topup/:userId", h.TopUp)
		e.GET("/remainingBalance/:userId", h.GetRemainingBalance)
	}

	// {
	// 	// GraphQL Handler
	// 	a.GET("/graphql-playground", func(c echo.Context) error { // GUIからのGraphQL実行用（http://localhost:8444/graphql-playground）
	// 		playground.Handler("GraphQL playground", "/query").ServeHTTP(c.Response(), c.Request())
	// 		return nil
	// 	})
	// 	a.POST("/query", func(c echo.Context) error {
	// 		appCtr.Server.ServeHTTP(c.Response(), c.Request())
	// 		return nil
	// 	})
	// }
}
