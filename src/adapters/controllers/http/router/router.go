package router

import (
	"go-playground/m/v1/src/dependency"

	"github.com/labstack/echo/v4"
)

// InitRouting ...
func InitRouting(e *echo.Echo, h dependency.HTTPHandler) {
	{
		// User Handler
		e.POST("/user", h.UserHandler.CreateNewUser)
		e.GET("/users", h.UserHandler.RetrieveAllUsers)
		e.GET("/user/:id", h.UserHandler.SearchUser)
	}

	{
		// Grade Handler
		e.GET("/grades", h.GradeHandler.RetrieveAllGrades)
	}

	{
		// DealHistory Handler
		e.GET("/dealHistories/:userId", h.DealHistoryHandler.RetrieveDealHistories)
	}

	{
		// BalanceControl Handler
		e.PUT("/pay/:userId", h.BalanceControlHandler.Pay)
		e.PUT("/topup/:userId", h.BalanceControlHandler.TopUp)
		e.GET("/remainingBalance/:userId", h.BalanceControlHandler.RetrieveRemainingBalance)
	}
}
