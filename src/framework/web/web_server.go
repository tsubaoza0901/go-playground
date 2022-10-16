package web

import (
	"log"

	"github.com/labstack/echo/v4"
)

// API ...
type API struct {
	*echoFW
	*AppHandler
}

// NewAPI ...
func NewAPI(ah *AppHandler) *API {
	return &API{
		echoFW: &echoFW{
			echo.New(),
		},
		AppHandler: ah,
	}
}

// StartServer ...
func (a *API) StartServer() {
	a.initCommonMddleware()
	a.initRouter(a.Echo)
	if err := a.Start(":8444"); err != nil {
		log.Fatal(err)
	}
}
