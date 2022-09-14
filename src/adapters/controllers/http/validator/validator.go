package validator

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

// CustomValidator CustomValidator
type CustomValidator struct {
	validator *validator.Validate
}

// NewValidator NewValidator
func NewValidator() echo.Validator {
	return &CustomValidator{validator: validator.New()}
}

// Validate validate
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}
