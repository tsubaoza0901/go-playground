package gateways

import (
	"context"
	"errors"
	"go-playground/m/v1/entities"
	"go-playground/m/v1/infrastructure/rdb/model"
	"log"
)

// UserGateway ...
type UserGateway struct {
	dbConn string // NOTE: 実際の型は *gorm.DB など
}

// NewUserGateway ...
func NewUserGateway(dbConn string) *UserGateway {
	return &UserGateway{dbConn}
}

// RegisterUser ...
func (g *UserGateway) RegisterUser(ctx context.Context, user *entities.User) (*entities.User, error) {
	rdbModel := model.User{
		ID:   1,
		Name: user.Name,
		Age:  user.Age,
	}
	entities := &entities.User{
		Name: rdbModel.Name,
		Age:  rdbModel.Age,
	}
	return entities, nil
}

// RetrieveUsers ...
func (g *UserGateway) RetrieveUsers(ctx context.Context) ([]*entities.User, error) {
	if g.dbConn == "" {
		return nil, errors.New("dbConnが空")
	}

	log.Printf("g.dbConn: %s", g.dbConn)

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
