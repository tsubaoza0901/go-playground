package ports

import (
	"context"
	"go-playground/m/v1/entities"
	"go-playground/m/v1/usecases/data/input"
	"go-playground/m/v1/usecases/data/output"
)

// ItemInportPort ...
type ItemInportPort interface {
	AddItem(ctx context.Context, item *input.Item) error
	FetchItems(ctx context.Context) error
}

// ItemOutputPort ...
type ItemOutputPort interface {
	OutputItems([]*output.Item) error
	OutputItem(*output.Item) error
	OutputError(error) error
}

// ItemRepository ...
type ItemRepository interface {
	RegisterItem(ctx context.Context, item *entities.Item) (*entities.Item, error)
	RetrieveItems(ctx context.Context) ([]*entities.Item, error)
}
