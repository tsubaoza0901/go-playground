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
		// Transaction Handler
		e.GET("/transactionHistories/:userId", h.TransactionHistoryHandler.RetrieveTransactionHistories)
	}

	{
		// Balance Control Handler
		e.PUT("/pay/:userId", h.BalanceControlHandler.Pay)
		e.PUT("/topup/:userId", h.BalanceControlHandler.TopUp)
		e.GET("/remainingBalance/:userId", h.BalanceControlHandler.RetrieveRemainingBalance)
	}
}
