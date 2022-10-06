package ports

import (
	"context"
	"go-playground/m/v1/entities"
	"go-playground/m/v1/usecases/data/output"
)

// UserInportPort ...
type UserInportPort interface {
	FetchUsers(ctx context.Context) error
}

// UserOutputPort ...
type UserOutputPort interface {
	OutputUsers([]*output.User) error
	OutputError(error) error
}

// UserRepository ...
type UserRepository interface {
	GetUsers(ctx context.Context) ([]*entities.User, error)
}
