package interactors

import (
	"context"
	"go-playground/m/v1/entities"
	"go-playground/m/v1/usecases/data/input"
	"go-playground/m/v1/usecases/data/output"
	"go-playground/m/v1/usecases/ports"
)

// Item ...
type Item struct {
	outputPort ports.ItemOutputPort
	repository ports.ItemRepository
}

// NewItem ...
func NewItem(iop ports.ItemOutputPort, ir ports.ItemRepository) *Item {
	return &Item{iop, ir}
}

// AddItem ...
func (u *Item) AddItem(ctx context.Context, in *input.Item) {
	item := &entities.Item{
		Name: in.Name,
	}
	result, err := u.repository.RegisterItem(ctx, item)
	if err != nil {
		u.outputPort.Error(err)
		return
	}

	out := &output.Item{
		Name: result.Name,
	}
	u.outputPort.Item(out)
}

// FetchItems ...
func (u *Item) FetchItems(ctx context.Context) {
	items, err := u.repository.RetrieveItems(ctx)
	if err != nil {
		u.outputPort.Error(err)
		return
	}
	outputs := make([]*output.Item, len(items))
	for i, v := range items {
		outputs[i] = &output.Item{
			Name: v.Name,
		}
	}
	u.outputPort.ItemList(outputs)
}
