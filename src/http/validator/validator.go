package validator

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

// CustomValidator CustomValidator構造体
type CustomValidator struct {
	validator *validator.Validate
}

// NewValidator NewValidator
func NewValidator() echo.Validator {
	return &CustomValidator{validator: validator.New()}
}

// Validate Validateメソッド
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}
