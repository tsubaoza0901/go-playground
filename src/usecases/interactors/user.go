package interactors

import (
	"context"
	"go-playground/m/v1/entities"
	"go-playground/m/v1/usecases/data/input"
	"go-playground/m/v1/usecases/data/output"
	"go-playground/m/v1/usecases/ports"
)

// User ...
type User struct {
	outputPort ports.UserOutputPort
	repository ports.UserRepository
}

// NewUser ...
func NewUser(uop ports.UserOutputPort, ur ports.UserRepository) *User {
	return &User{uop, ur}
}

// AddUser ...
func (u *User) AddUser(ctx context.Context, in *input.User) {
	user := &entities.User{
		Name: in.Name,
		Age:  in.Age,
	}
	result, err := u.repository.RegisterUser(ctx, user)
	if err != nil {
		u.outputPort.Error(err)
		return
	}

	out := &output.User{
		Name: result.Name,
		Age:  result.Age,
	}
	u.outputPort.User(out)
}

// FetchUserByID ...
func (u *User) FetchUserByID(ctx context.Context, id uint) {
	user, err := u.repository.RetrieveUserWithItem(ctx, id)
	if err != nil {
		u.outputPort.Error(err)
		return
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
	u.outputPort.UserWithItem(outputUser)
}

// FetchUsers ...
func (u *User) FetchUsers(ctx context.Context) {
	users, err := u.repository.RetrieveUsers(ctx)
	if err != nil {
		u.outputPort.Error(err)
		return
	}
	outputs := make([]*output.User, len(users))
	for i, v := range users {
		outputs[i] = &output.User{
			Name: v.Name,
			Age:  v.Age,
		}
	}
	u.outputPort.UserList(outputs)
}
