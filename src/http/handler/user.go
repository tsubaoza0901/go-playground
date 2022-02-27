package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// User user info
type User struct {
	ID        string `json:"id_str" validate:"required,len=5"`
	Name      string `json:"name"  validate:"contains=yama"`
	Age       uint   `json:"age" validate:"min=1,max=100"`
	Email     string `json:"email" validate:"email"`
	Array     []uint `json:"array" validate:"len=3"`
	URL       string `json:"url" validate:"url_encoded"`
	Base64    string `json:"base_64" validate:"base64"`
	Base64URL string `json:"base_64_url" validate:"base64url"`
	Data      string `json:"data" validate:"datauri"`
}

func CreateUser(c echo.Context) error {
	_ = c.Request().Context()

	req := &User{}
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusOK, err.Error())
	}

	// validate req
	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusOK, err.Error())
	}
	return c.JSON(http.StatusOK, "Hello, World!")
}
