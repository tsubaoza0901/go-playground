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

type (
	// ItemInputPortFactory ...
	ItemInputPortFactory func(ports.ItemOutputPort) *interactors.ItemInteractor

	// ItemOutputPortFactory ...
	ItemOutputPortFactory func(e echo.Context) *presenters.ItemPresenter

	// ItemDependency ...
	ItemDependency struct {
		ItemInputPortFactory  ItemInputPortFactory
		ItemOutputPortFactory ItemOutputPortFactory
		ItemRepository        ports.ItemRepository
	}
)

// NewItemDependency ...
func NewItemDependency(itemInputPortFactory ItemInputPortFactory, itemOutputPortFactory ItemOutputPortFactory, itemRepository ports.ItemRepository) *ItemDependency {
	return &ItemDependency{
		ItemInputPortFactory:  itemInputPortFactory,
		ItemOutputPortFactory: itemOutputPortFactory,
		ItemRepository:        itemRepository,
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

// InitItemDI ...
func (d *AppDependency) InitItemDI() *ItemDependency {
	return NewItemDependency(
		d.InitItemInteractor,
		presenters.NewItemPresenter,
		d.InitItemGateway(),
	)
}

// InitItemInteractor ...
func (d *AppDependency) InitItemInteractor(itemOutputPort ports.ItemOutputPort) *interactors.ItemInteractor {
	return interactors.NewItemInteractor(
		itemOutputPort,
		d.InitItemGateway(),
	)
}

// InitUserGateway ...
func (d *db) InitUserGateway() *gateways.UserGateway {
	return gateways.NewUserGateway(d.conn)
}

// InitItemGateway ...
func (d *db) InitItemGateway() *gateways.ItemGateway {
	return gateways.NewItemGateway(d.conn)
}
