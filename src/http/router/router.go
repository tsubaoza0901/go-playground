package router

import (
	"go-playground/m/v1/src/http/handler"

	"github.com/labstack/echo/v4"
)

// InitRouting ...
func InitRouting(e *echo.Echo, jwtConfig echo.MiddlewareFunc) {
	e.POST("/login", handler.Login)

	api := e.Group("/api")

	api.Use(jwtConfig)
	api.GET("/private", handler.Private)
	api.GET("/logout", handler.Logout)
}
