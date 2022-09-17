package http

import "github.com/labstack/echo/v4"

// IUserHandler ...
type IUserHandler interface {
	CreateNewUser(c echo.Context) error
	GetUser(c echo.Context) error
	GetUserList(c echo.Context) error
}

// IGradeHandler ...
type IGradeHandler interface {
	GetGradeList(c echo.Context) error
}

// IDealHistoryHandler ...
type IDealHistoryHandler interface {
	GetDealHistoryList(c echo.Context) error
}

// IBalanceControlHandler ...
type IBalanceControlHandler interface {
	Pay(c echo.Context) error
	TopUp(c echo.Context) error
	GetRemainingBalance(c echo.Context) error
}
