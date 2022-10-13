package gateways

import (
	"context"
	"errors"
	"go-playground/m/v1/entities"
	"go-playground/m/v1/infrastructure/rdb/model"
	"log"
)

// User ...
type User struct {
	dbConn string // NOTE: 実際の型は *gorm.DB など
}

// NewUser ...
func NewUser(dbConn string) *User {
	return &User{dbConn}
}

// RegisterUser ...
func (g *User) RegisterUser(ctx context.Context, user *entities.User) (*entities.User, error) {
	if g.dbConn == "" {
		return nil, errors.New("dbConnが空")
	}

	log.Printf("g.dbConn: %s", g.dbConn)

	rdbModel := model.User{
		ID:   1,
		Name: user.Name,
		Age:  user.Age,
	}
	entity := &entities.User{
		Name: rdbModel.Name,
		Age:  rdbModel.Age,
	}
	return entity, nil
}

// RetrieveUserWithItem ...
func (g *User) RetrieveUserWithItem(ctx context.Context, id uint) (*entities.User, error) {
	if g.dbConn == "" {
		return nil, errors.New("dbConnが空")
	}

	log.Printf("g.dbConn: %s", g.dbConn)

	rdbUserModel := model.User{
		ID:   id,
		Name: "Yamada Taro",
		Age:  20,
	}
	entityUser := &entities.User{
		Name: rdbUserModel.Name,
		Age:  rdbUserModel.Age,
	}

	rdbItemModels := []*model.Item{

		{
			ID:   1,
			Name: "xxxの本",
		},
		{
			ID:   2,
			Name: "古びた剣",
		},
	}
	entityItems := make([]*entities.Item, len(rdbItemModels))
	for i, v := range rdbItemModels {
		entityItems[i] = &entities.Item{
			Name: v.Name,
		}
	}
	entityUser.Items = entityItems
	return entityUser, nil
}

// RetrieveUsers ...
func (g *User) RetrieveUsers(ctx context.Context) ([]*entities.User, error) {
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
