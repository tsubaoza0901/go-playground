package ports

import (
	"context"
	"go-playground/m/v1/entities"
	"go-playground/m/v1/usecases/data/input"
	"go-playground/m/v1/usecases/data/output"
)

// UserInportPort ...
type UserInportPort interface {
	AddUser(ctx context.Context, user *input.User)
	FetchUserByID(ctx context.Context, id uint)
	FetchUsers(ctx context.Context)
}

// UserOutputPort ...
type UserOutputPort interface {
	User(*output.User)
	UserWithItem(*output.UserWithItem)
	UserList([]*output.User)
	Error(error)
}

// UserRepository ...
type UserRepository interface {
	RegisterUser(ctx context.Context, user *entities.User) (*entities.User, error)
	RetrieveUserWithItem(ctx context.Context, id uint) (*entities.User, error)
	RetrieveUsers(ctx context.Context) ([]*entities.User, error)
}
