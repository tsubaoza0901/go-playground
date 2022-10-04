package presenters

import (
	"go-playground/m/v1/src/usecases/data/output"
	"net/http"

	"github.com/labstack/echo/v4"
)

// UserPresenter ...
type UserPresenter struct {
	c echo.Context
}

// NewUserPresenter ...
func NewUserPresenter(c echo.Context) *UserPresenter {
	return &UserPresenter{c}
}

// SetEchoContext ...
func (u *UserPresenter) SetEchoContext(c echo.Context) {
	u.c = c
}

// OutputUsers ...
func (u *UserPresenter) OutputUsers(users []*output.User) error {
	return u.c.JSON(http.StatusOK, users)
}

// OutputError ...
func (u *UserPresenter) OutputError(err error) error {
	return u.c.JSON(http.StatusInternalServerError, err)
}
