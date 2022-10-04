package controllers

import (
	"go-playground/m/v1/src/controllers/rest"
	"go-playground/m/v1/src/injector"

	"github.com/labstack/echo/v4"
)

// AppControllers ...
type AppControllers struct {
	rest.AppHandlers
}

// NewAppControllers ...
func NewAppControllers() *AppControllers {
	return &AppControllers{}
}

// InitRouting ...
func InitRouting(e *echo.Echo, appControllers *AppControllers, d *injector.AppDependency) {
	e.GET("/users", appControllers.GetUsers(d.InitUserDI()))
}
