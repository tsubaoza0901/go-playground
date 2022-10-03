package presenters

import (
	"go-playground/m/v1/src/usecases/data/output"
	"go-playground/m/v1/src/usecases/ports"
	"net/http"

	"github.com/labstack/echo/v4"
)

// UserPresenter ...
type UserPresenter struct {
	c echo.Context
}

// NewUserOutputPort ...
func NewUserOutputPort(c echo.Context) ports.UserOutputPort {
	return &UserPresenter{c}
}

// OutputUsers ...
func (u *UserPresenter) OutputUsers(users []*output.User) error {
	return u.c.JSON(http.StatusOK, users)
}

// OutputError ...
func (u *UserPresenter) OutputError(err error) error {
	return u.c.JSON(http.StatusInternalServerError, err)
}
