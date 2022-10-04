package handler

import (
	"go-playground/m/v1/src/injector"
	"go-playground/m/v1/src/usecases/ports"

	"github.com/labstack/echo/v4"
)

// UserHandler ...
type UserHandler struct {
	userInportPort ports.UserInportPort
}

// GetUsers ...
func (u *UserHandler) GetUsers(ud *injector.UserDependency) func(echo.Context) error {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		return u.newInputPort(c, ud).FetchUsers(ctx)
	}
}

func (u *UserHandler) newInputPort(c echo.Context, ud *injector.UserDependency) ports.UserInportPort {
	return ud.UserInputPort(ud.UserOutputPort(c))
}
