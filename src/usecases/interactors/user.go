package interactors

import (
	"context"
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

// FetchUsers ...
func (u *UserInteractor) FetchUsers(ctx context.Context) error {
	users, err := u.Repository.GetUsers(ctx)
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
