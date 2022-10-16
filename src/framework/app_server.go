package framework

import (
	"go-playground/m/v1/dependency"
	"go-playground/m/v1/framework/web"

	"gorm.io/gorm"
)

// App ...
type App struct {
	webAPI *web.API
}

// NewApp ...
func NewApp(db *gorm.DB) *App {
	di := dependency.NewInjection(db)
	return &App{
		webAPI: di.InitWebAPI(),
	}
}

// Start ...
func (a *App) Start() {
	a.webAPI.StartServer()
}
