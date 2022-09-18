package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"math/rand"

	"go-playground/m/v1/src/adapters/controllers/graphql/graph/generated"
	"go-playground/m/v1/src/adapters/controllers/graphql/graph/model"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	user := &model.User{
		ID:     		fmt.Sprintf("T%d", rand.Int()),
		Name:   		input.Name,
		Age:   			input.Age,
		EmailAddress:   input.EmailAddress,
		Grade: 			&model.Grade{ID: input.GradeID, Name: "grade " + input.GradeID},
	}
	r.users = append(r.users, user)
	return user, nil
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	return r.users, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
