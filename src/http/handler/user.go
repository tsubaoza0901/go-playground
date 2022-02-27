package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// User user info
type User struct {
	ID   string `json:"id_str"`
	Name string `json:"name"`
	Age  uint   `json:"age"`
}

func Top(c echo.Context) error {
	return c.JSON(http.StatusOK, "Hello, World!")
}

func Logout(c echo.Context) error {
	return c.Redirect(http.StatusFound, "/")
}
