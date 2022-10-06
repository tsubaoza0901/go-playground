package presenters

import (
	"go-playground/m/v1/usecases/data/output"
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

// OutputUsers ...
func (u *UserPresenter) OutputUsers(users []*output.User) error {
	return u.c.JSON(http.StatusOK, users)
}

// OutputUser ...
func (u *UserPresenter) OutputUser(user *output.User) error {
	return u.c.JSON(http.StatusOK, user)
}

// OutputUserWithItem ...
func (u *UserPresenter) OutputUserWithItem(user *output.UserWithItem) error {
	return u.c.JSON(http.StatusOK, user)
}

// OutputError ...
func (u *UserPresenter) OutputError(err error) error {
	return u.c.JSON(http.StatusInternalServerError, err)
}
