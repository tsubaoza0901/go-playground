package controllers

import (
	"go-playground/m/v1/src/injector"

	"github.com/labstack/echo/v4"
)

// InitRouting ...
func InitRouting(e *echo.Echo, handler injector.AppHandlers) {
	e.GET("/users", handler.GetUsers)
}
