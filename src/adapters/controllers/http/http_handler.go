package http

import "github.com/labstack/echo/v4"

// IBalanceControlHandler ...
type IBalanceControlHandler interface {
	Pay(c echo.Context) error
	TopUp(c echo.Context) error
	RetrieveRemainingBalance(c echo.Context) error
}

// IDealHistoryHandler ...
type IDealHistoryHandler interface {
	RetrieveDealHistories(c echo.Context) error
}

// IGradeHandler ...
type IGradeHandler interface {
	RetrieveAllGrades(c echo.Context) error
}

// IUserHandler ...
type IUserHandler interface {
	CreateNewUser(c echo.Context) error
	SearchUser(c echo.Context) error
	RetrieveAllUsers(c echo.Context) error
}
