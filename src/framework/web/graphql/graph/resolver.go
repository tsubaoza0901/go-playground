//go:generate go get github.com/99designs/gqlgen@v0.17.19
//go:generate go run github.com/99designs/gqlgen generate

package graph

import (
	"go-playground/m/v1/framework/web/graphql/graph/controllers"
	"go-playground/m/v1/presenters"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// Resolver ...
type Resolver struct {
	// usecase.IUserManagementUsecase
	userController *controllers.User
	*presenters.User
}

// NewResolver ...
func NewResolver(uc *controllers.User, up *presenters.User) *Resolver {
	return &Resolver{uc, up}
}
