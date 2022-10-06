package controllers

import (
	"go-playground/m/v1/controllers/rest"
	"go-playground/m/v1/injector"

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
	e.POST("/user", appControllers.CreateUser(d.InitUserDI()))
	e.GET("/users", appControllers.GetUsers(d.InitUserDI()))
}
