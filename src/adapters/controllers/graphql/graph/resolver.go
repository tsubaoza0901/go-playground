//go:generate go run github.com/99designs/gqlgen generate

package graph

import (
	"go-playground/m/v1/src/usecase"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// Resolver ...
type Resolver struct {
	usecase.IUserManagementUsecase
}

// NewResolver ...
func NewResolver(u usecase.IUserManagementUsecase) *Resolver {
	return &Resolver{u}
}
