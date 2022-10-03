package handler

import (
	"go-playground/m/v1/src/usecases/ports"

	"github.com/labstack/echo/v4"
)

// User ...
type User interface {
	GetUsers(c echo.Context) error
}

// OutputFactory ...
type OutputFactory func(echo.Context) ports.UserOutputPort

// InputFactory ...
type InputFactory func(ports.UserOutputPort, ports.UserRepository) ports.UserInportPort

// RepositoryFactory ...
type RepositoryFactory func() ports.UserRepository

// UserHandler ...
type UserHandler struct {
	outputFactory     OutputFactory
	inputFactory      InputFactory
	repositoryFactory RepositoryFactory
}

// NewUserHandler ...
func NewUserHandler(outputFactory OutputFactory, inputFactory InputFactory, repositoryFactory RepositoryFactory) User {
	return &UserHandler{
		outputFactory:     outputFactory,
		inputFactory:      inputFactory,
		repositoryFactory: repositoryFactory,
	}
}

// GetUsers ...
func (u *UserHandler) GetUsers(c echo.Context) error {
	ctx := c.Request().Context()
	return u.newInputPort(c).GetUsers(ctx)
}

func (u *UserHandler) newInputPort(c echo.Context) ports.UserInportPort {
	outputPort := u.outputFactory(c)
	repository := u.repositoryFactory()
	return u.inputFactory(outputPort, repository)
}
