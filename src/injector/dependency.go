package injector

import (
	"go-playground/m/v1/src/gateways"
	"go-playground/m/v1/src/presenters"
	"go-playground/m/v1/src/usecases/interactors"
	"go-playground/m/v1/src/usecases/ports"

	"github.com/labstack/echo/v4"
)

type (
	// UserInputPortFactory ...
	UserInputPortFactory func(ports.UserOutputPort) *interactors.UserInteractor

	// UserOutputPortFactory ...
	UserOutputPortFactory func(e echo.Context) *presenters.UserPresenter

	// UserDependency ...
	UserDependency struct {
		UserInputPort  UserInputPortFactory
		UserOutputPort UserOutputPortFactory
		UserRepository ports.UserRepository
	}
)

// NewUserDependency ...
func NewUserDependency(userInputPortFactory UserInputPortFactory, userOutputPortFactory UserOutputPortFactory, userRepository ports.UserRepository) *UserDependency {
	return &UserDependency{
		UserInputPort:  userInputPortFactory,
		UserOutputPort: userOutputPortFactory,
		UserRepository: userRepository,
	}
}

type db struct {
	dbConn string
}

// AppDependency ...
type AppDependency struct {
	db
}

// NewAppDependency ...
func NewAppDependency(dbConn string) *AppDependency {
	return &AppDependency{
		db: db{dbConn},
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
	return gateways.NewUserGateway(d.dbConn)
}
