package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"go-playground/m/v1/src/adapters/controllers/graphql/graph/generated"
	"go-playground/m/v1/src/adapters/controllers/graphql/graph/model"
	usecaseInput "go-playground/m/v1/src/usecase/data/input"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (string, error) {
	user := usecaseInput.NewUserCreate()
	user.FirstName = input.FirstName
	user.LastName = input.LastName
	user.Age = uint(input.Age)
	user.EmailAddress = input.Email

	if err := r.IUserManagementUsecase.CreateUser(ctx, user, uint(input.Amount)); err != nil {
		return "", err
	}
	return "ok", nil
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	users, err := r.RetrieveUsers(ctx)
	if err != nil {
		return nil, err
	}
	us := make([]*model.User, len(users))
	for i, user := range users {
		us[i] = &model.User{
			ID:           int(user.ID),
			Name:         user.MakeJPNFullName(),
			Age:          int(user.Age),
			EmailAddress: user.EmailAddress,
			GradeName:    user.GradeName,
		}
	}
	return us, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
