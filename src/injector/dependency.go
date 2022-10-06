package injector

import (
	"go-playground/m/v1/gateways"
	"go-playground/m/v1/presenters"
	"go-playground/m/v1/usecases/interactors"
	"go-playground/m/v1/usecases/ports"

	"github.com/labstack/echo/v4"
)

type (
	// UserInputPortFactory ...
	UserInputPortFactory func(ports.UserOutputPort) *interactors.UserInteractor

	// UserOutputPortFactory ...
	UserOutputPortFactory func(e echo.Context) *presenters.UserPresenter

	// UserDependency ...
	UserDependency struct {
		UserInputPortFactory  UserInputPortFactory
		UserOutputPortFactory UserOutputPortFactory
		UserRepository        ports.UserRepository
	}
)

// NewUserDependency ...
func NewUserDependency(userInputPortFactory UserInputPortFactory, userOutputPortFactory UserOutputPortFactory, userRepository ports.UserRepository) *UserDependency {
	return &UserDependency{
		UserInputPortFactory:  userInputPortFactory,
		UserOutputPortFactory: userOutputPortFactory,
		UserRepository:        userRepository,
	}
}

type db struct {
	conn string
}

// AppDependency ...
type AppDependency struct {
	db
}

// NewAppDependency ...
func NewAppDependency(conn string) *AppDependency {
	return &AppDependency{
		db: db{conn},
	}
}

// InitUserDI ...
func (d *AppDependency) InitUserDI() *UserDependency {
	return NewUserDependency(
		d.InitUserInteractor,
		presenters.NewUserPresenter,
		d.InitUserGateway(),
	)
}

// InitUserInteractor ...
func (d *AppDependency) InitUserInteractor(userOutputPort ports.UserOutputPort) *interactors.UserInteractor {
	return interactors.NewUserInteractor(
		userOutputPort,
		d.InitUserGateway(),
	)
}

// InitUserGateway ...
func (d *db) InitUserGateway() *gateways.UserGateway {
	return gateways.NewUserGateway(d.conn)
}
