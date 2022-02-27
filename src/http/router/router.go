package router

import (
	"go-playground/m/v1/src/http/handler"

	"github.com/labstack/echo/v4"
)

// InitRouting ...
func InitRouting(e *echo.Echo) {
	e.POST("/users", handler.CreateUser)
}
