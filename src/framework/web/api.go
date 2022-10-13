package web

import (
	"go-playground/m/v1/framework/web/rest/handler"
	"log"

	"github.com/labstack/echo/v4"
)

// API ...
type API struct {
	*echoFW
	*handler.User
	*handler.Item
}

// NewAPI ...
func NewAPI(e *echo.Echo, u *handler.User, i *handler.Item) *API {
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
