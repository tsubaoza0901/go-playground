package ports

import (
	"context"
	"go-playground/m/v1/src/entities"
	"go-playground/m/v1/src/usecases/data/output"
)

// UserInportPort ...
type UserInportPort interface {
	GetUsers(ctx context.Context) error
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
