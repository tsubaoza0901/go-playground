package gateways

import (
	"context"
	"errors"
	"go-playground/m/v1/entities"
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

// GetUsers ...
func (g *UserGateway) GetUsers(ctx context.Context) ([]*entities.User, error) {

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
