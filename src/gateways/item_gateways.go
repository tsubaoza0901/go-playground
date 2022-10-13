package gateways

import (
	"context"
	"errors"
	"go-playground/m/v1/entities"
	"go-playground/m/v1/infrastructure/rdb/model"
	"log"
)

// Item ...
type Item struct {
	dbConn string // NOTE: 実際の型は *gorm.DB など
}

// NewItem ...
func NewItem(dbConn string) *Item {
	return &Item{dbConn}
}

// RegisterItem ...
func (g *Item) RegisterItem(ctx context.Context, item *entities.Item) (*entities.Item, error) {
	if g.dbConn == "" {
		return nil, errors.New("dbConnが空")
	}

	log.Printf("g.dbConn: %s", g.dbConn)

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
func (g *Item) RetrieveItems(ctx context.Context) ([]*entities.Item, error) {
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
