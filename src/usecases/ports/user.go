package ports

import (
	"context"
	"go-playground/m/v1/entities"
	"go-playground/m/v1/usecases/data/input"
	"go-playground/m/v1/usecases/data/output"
)

// UserInportPort ...
type UserInportPort interface {
	AddUser(ctx context.Context, user *input.User) error
	FetchUserByID(ctx context.Context, id uint) error
	FetchUsers(ctx context.Context) error
}

// UserOutputPort ...
type UserOutputPort interface {
	OutputUsers([]*output.User) error
	OutputUser(*output.User) error
	OutputUserWithItem(*output.UserWithItem) error
	OutputError(error) error
}

// UserRepository ...
type UserRepository interface {
	RegisterUser(ctx context.Context, user *entities.User) (*entities.User, error)
	RetrieveUserWithItem(ctx context.Context, id uint) (*entities.User, error)
	RetrieveUsers(ctx context.Context) ([]*entities.User, error)
}
