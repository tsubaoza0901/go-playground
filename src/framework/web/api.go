package web

import (
	"go-playground/m/v1/framework/web/rest"
	"log"

	"github.com/labstack/echo/v4"
)

// API ...
type API struct {
	*echoFW
	*rest.UserHandler
	*rest.ItemHandler
}

// NewAPI ...
func NewAPI(e *echo.Echo, u *rest.UserHandler, i *rest.ItemHandler) *API {
	return &API{&echoFW{e}, u, i}
}

// StartServer ...
func (a *API) StartServer() {
	a.middleware()
	a.initRouting()
	if err := a.Start(":8444"); err != nil {
		log.Fatal(err)
	}
}

func (a *API) initRouting() {
	a.GET("/user/:id", a.GetUserByID)
	a.GET("/users", a.GetUsers)
	a.POST("/user", a.CreateUser)

	a.POST("/item", a.CreateItem)
	a.GET("/items", a.GetItems)
}
