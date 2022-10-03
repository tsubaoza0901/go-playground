package interactors

import (
	"context"
	"go-playground/m/v1/src/usecases/data/output"
	"go-playground/m/v1/src/usecases/ports"
)

// UserInteractor ...
type UserInteractor struct {
	OutputPort ports.UserOutputPort
	Repository ports.UserRepository
}

// NewUserInputPort ...
func NewUserInputPort(outputPort ports.UserOutputPort, repository ports.UserRepository) ports.UserInportPort {
	return &UserInteractor{
		OutputPort: outputPort,
		Repository: repository,
	}
}

// GetUsers ...
func (u *UserInteractor) GetUsers(ctx context.Context) error {
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
