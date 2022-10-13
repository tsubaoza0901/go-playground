package injector

import (
	"go-playground/m/v1/controllers"
	"go-playground/m/v1/framework/web"
	"go-playground/m/v1/framework/web/rest/handler"
	"go-playground/m/v1/gateways"
	"go-playground/m/v1/presenters"
	"go-playground/m/v1/usecases/interactors"
	"go-playground/m/v1/usecases/ports"

	"github.com/labstack/echo/v4"
)

type echoFW struct {
	*echo.Echo
}

type db struct {
	conn string
}

// AppDependency ...
type AppDependency struct {
	echoFW
	db
}

// NewAppDependency ...
func NewAppDependency(conn string) *AppDependency {
	return &AppDependency{
		echoFW: echoFW{echo.New()},
		db:     db{conn},
	}
}

// InitWebAPI ...
func (d *AppDependency) InitWebAPI() *web.API {
	return web.NewAPI(
		d.Echo,
		d.InitUserHandler(),
		d.InitItemHandler(),
	)
}

// InitUserHandler ...
func (d *AppDependency) InitUserHandler() *handler.User {
	userPresenter := d.InitUserPresenter()
	userController := d.InitUserController(userPresenter)

	return handler.NewUser(userController, userPresenter)
}

// InitItemHandler ...
func (d *AppDependency) InitItemHandler() *handler.Item {
	itemPresenter := d.InitItemPresenter()
	itemController := d.InitItemController(itemPresenter)

	return handler.NewItem(itemController, itemPresenter)
}

// InitUserController ...
func (d *AppDependency) InitUserController(userOutputPort ports.UserOutputPort) *controllers.User {
	return controllers.NewUser(
		interactors.NewUser(
			userOutputPort,
			d.InitUserGateway(),
		),
	)
}

// InitItemController ...
func (d *AppDependency) InitItemController(itemOutputPort ports.ItemOutputPort) *controllers.Item {
	return controllers.NewItem(
		interactors.NewItem(
			itemOutputPort,
			d.InitItemGateway(),
		),
	)
}

// InitUserGateway ...
func (d *db) InitUserGateway() *gateways.User {
	return gateways.NewUser(d.conn)
}

// InitItemGateway ...
func (d *db) InitItemGateway() *gateways.Item {
	return gateways.NewItem(d.conn)
}

// InitUserPresenter ...
func (d *AppDependency) InitUserPresenter() *presenters.User {
	return presenters.NewUser()
}

// InitItemPresenter ...
func (d *AppDependency) InitItemPresenter() *presenters.Item {
	return presenters.NewItem()
}
