package injector

import (
	"go-playground/m/v1/src/controllers/handler"
	"go-playground/m/v1/src/gateways"
	"go-playground/m/v1/src/presenters"
	"go-playground/m/v1/src/usecases/interactors"
)

// Dependency ...
type Dependency struct {
	// e *echo.Echo
}

// NewDependency ...
func NewDependency() *Dependency {
	return &Dependency{}
}

// AppControllers ...
type AppControllers struct {
	AppHandlers
}

// InitAppControllers ...
func (d *Dependency) InitAppControllers() AppControllers {
	return AppControllers{
		AppHandlers: d.InitAppHandlers(),
	}
}

// AppHandlers ...
type AppHandlers struct {
	handler.User
}

// InitAppHandlers ...
func (d *Dependency) InitAppHandlers() AppHandlers {
	return AppHandlers{
		User: d.InitUserHandler(),
	}
}

// InitUserHandler ...
func (d *Dependency) InitUserHandler() handler.User {
	return handler.NewUserHandler(
		d.InitUserOutputFactory(),
		d.InitUserInputFactory(),
		d.InitUserRepositoryFactory(),
	)
}

// InitUserOutputFactory ...
func (d *Dependency) InitUserOutputFactory() handler.OutputFactory {
	return presenters.NewUserOutputPort
}

// InitUserInputFactory ...
func (d *Dependency) InitUserInputFactory() handler.InputFactory {
	return interactors.NewUserInputPort
}

// InitUserRepositoryFactory ...
func (d *Dependency) InitUserRepositoryFactory() handler.RepositoryFactory {
	return gateways.NewUserRepository
}
