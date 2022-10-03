package gateways

import (
	"context"
	"go-playground/m/v1/src/entities"
	"go-playground/m/v1/src/usecases/ports"
)

// UserGateway ...
type UserGateway struct {
}

// NewUserRepository ...
func NewUserRepository() ports.UserRepository {
	return &UserGateway{}
}

// GetUsers ...
func (g *UserGateway) GetUsers(ctx context.Context) ([]*entities.User, error) {
	users := []*entities.User{
		{
			Name: "Yamada Taro",
			Age:  20,
		},
		{
			Name: "Hanagaki Takemichi",
			Age:  10,
		},
	}
	return users, nil
}
