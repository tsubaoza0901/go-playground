//go:generate go run github.com/99designs/gqlgen generate

package graph

import "go-playground/m/v1/src/adapters/controllers/graphql/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	// todos []*model.Todo
	users []*model.User
}
