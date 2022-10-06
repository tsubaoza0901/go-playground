package interactors

import (
	"context"
	"go-playground/m/v1/entities"
	"go-playground/m/v1/usecases/data/input"
	"go-playground/m/v1/usecases/data/output"
	"go-playground/m/v1/usecases/ports"
)

// UserInteractor ...
type UserInteractor struct {
	OutputPort ports.UserOutputPort
	Repository ports.UserRepository
}

// NewUserInteractor ...
func NewUserInteractor(outputPort ports.UserOutputPort, repository ports.UserRepository) *UserInteractor {
	return &UserInteractor{
		OutputPort: outputPort,
		Repository: repository,
	}
}

// AddUser ...
func (u *UserInteractor) AddUser(ctx context.Context, in *input.User) error {
	user := &entities.User{
		Name: in.Name,
		Age:  in.Age,
	}
	result, err := u.Repository.RegisterUser(ctx, user)
	if err != nil {
		return err
	}

	out := &output.User{
		Name: result.Name,
		Age:  result.Age,
	}
	return u.OutputPort.OutputUser(out)
}

// FetchUserByID ...
func (u *UserInteractor) FetchUserByID(ctx context.Context, id uint) error {
	user, err := u.Repository.RetrieveUserWithItem(ctx, id)
	if err != nil {
		return err
	}
	outputUser := &output.UserWithItem{
		Name: user.Name,
		Age:  user.Age,
	}
	if user.Items != nil {
		outputItems := make([]*output.Item, len(user.Items))
		for i, v := range user.Items {
			outputItems[i] = &output.Item{
				Name: v.Name,
			}
		}
		outputUser.Items = outputItems
	}
	return u.OutputPort.OutputUserWithItem(outputUser)
}

// FetchUsers ...
func (u *UserInteractor) FetchUsers(ctx context.Context) error {
	users, err := u.Repository.RetrieveUsers(ctx)
	if err != nil {
		return u.OutputPort.OutputError(err)
	}
	outputs := make([]*output.User, len(users))
	for i, v := range users {
		outputs[i] = &output.User{
			Name: v.Name,
			Age:  v.Age,
		}
	}
	return u.OutputPort.OutputUsers(outputs)
}
