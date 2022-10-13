package framework

import (
	"go-playground/m/v1/framework/web"
	"go-playground/m/v1/injector"
)

// App ...
type App struct {
	webAPI *web.API
}

// NewApp ...
func NewApp() *App {
	dbConn := "*gorm.DB"
	di := injector.NewAppDependency(dbConn)

	return &App{
		webAPI: di.InitWebAPI(),
	}
}

// Start ...
func (a *App) Start() {
	a.webAPI.StartServer()
}
