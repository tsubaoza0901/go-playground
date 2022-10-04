package rest

import "go-playground/m/v1/src/controllers/rest/handler"

// AppHandlers ...
type AppHandlers struct {
	handler.UserHandler
}

// NewAppHandlers ...
func NewAppHandlers() *AppHandlers {
	return &AppHandlers{}
}

// // User ...
// type User interface {
// 	// GetUsers(c echo.Context) error
// 	GetUsers(injector.UserDependency) func(echo.Context) error
// }
