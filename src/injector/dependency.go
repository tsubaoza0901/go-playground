package injector

import (
	"go-playground/m/v1/src/gateways"
	"go-playground/m/v1/src/presenters"
	"go-playground/m/v1/src/usecases/interactors"
	"go-playground/m/v1/src/usecases/ports"

	"github.com/labstack/echo/v4"
)

// InputFactory ...
type InputFactory func(ports.UserOutputPort, ports.UserRepository) *interactors.UserInteractor

// OutputFactory ...
type OutputFactory func(e echo.Context) *presenters.UserPresenter

// UserDependency ...
type UserDependency struct {
	InputFactory  InputFactory
	OutputFactory OutputFactory
	Repository    ports.UserRepository
}

// NewUserDependency ...
func NewUserDependency(inputFactory InputFactory, outputFactory OutputFactory, repository ports.UserRepository) *UserDependency {
	return &UserDependency{
		InputFactory:  inputFactory,
		OutputFactory: outputFactory,
		Repository:    repository,
	}
}

type webFW struct {
	c echo.Context
}

type db struct {
	dbConn string
}

// AppDependency ...
type AppDependency struct {
	webFW
	db
}

// NewAppDependency ...
func NewAppDependency(dbConn string) *AppDependency {
	return &AppDependency{
		db: db{"gorm.DB"},
		// webFW: webFW{c},
	}
}

// InitUserDI ...
func (d *AppDependency) InitUserDI() *UserDependency {
	return NewUserDependency(
		interactors.NewUserInteractor,
		presenters.NewUserPresenter,
		d.InitUserGateway(),
	)
}

// InitUserGateway ...
func (d *db) InitUserGateway() *gateways.UserGateway {
	return gateways.NewUserGateway(d.dbConn)
}
