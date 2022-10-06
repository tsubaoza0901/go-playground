package gateways

import (
	"context"
	"errors"
	"go-playground/m/v1/entities"
	"go-playground/m/v1/infrastructure/rdb/model"
	"log"
)

// ItemGateway ...
type ItemGateway struct {
	dbConn string // NOTE: 実際の型は *gorm.DB など
}

// NewItemGateway ...
func NewItemGateway(dbConn string) *ItemGateway {
	return &ItemGateway{dbConn}
}

// RegisterItem ...
func (g *ItemGateway) RegisterItem(ctx context.Context, item *entities.Item) (*entities.Item, error) {
	rdbModel := model.Item{
		ID:   1,
		Name: item.Name,
	}
	entities := &entities.Item{
		Name: rdbModel.Name,
	}
	return entities, nil
}

// RetrieveItems ...
func (g *ItemGateway) RetrieveItems(ctx context.Context) ([]*entities.Item, error) {
	if g.dbConn == "" {
		return nil, errors.New("dbConnが空")
	}

	log.Printf("g.dbConn: %s", g.dbConn)

	items := []*entities.Item{
		{
			Name: "xxxの本",
		},
		{
			Name: "古びた剣",
		},
	}
	return items, nil
}
