package handler

import (
	"go-playground/m/v1/controllers/rest/handler/request"
	"go-playground/m/v1/injector"
	"go-playground/m/v1/usecases/data/input"
	"go-playground/m/v1/usecases/ports"

	"github.com/labstack/echo/v4"
)

// UserHandler ...
type UserHandler struct {
	userInportPort ports.UserInportPort
}

// CreateUser ...
func (u *UserHandler) CreateUser(ud *injector.UserDependency) func(echo.Context) error {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		req := new(request.User)
		if err := c.Bind(req); err != nil {
			return err
		}

		user := &input.User{
			Name: req.Name,
			Age:  req.Age,
		}
		return u.newInputPort(c, ud).AddUser(ctx, user)
	}
}

// GetUsers ...
func (u *UserHandler) GetUsers(ud *injector.UserDependency) func(echo.Context) error {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		return u.newInputPort(c, ud).FetchUsers(ctx)
	}
}

func (u *UserHandler) newInputPort(c echo.Context, dependency *injector.UserDependency) ports.UserInportPort {
	return dependency.UserInputPortFactory(dependency.UserOutputPortFactory(c))
}
