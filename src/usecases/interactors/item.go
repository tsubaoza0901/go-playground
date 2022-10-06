package interactors

import (
	"context"
	"go-playground/m/v1/entities"
	"go-playground/m/v1/usecases/data/input"
	"go-playground/m/v1/usecases/data/output"
	"go-playground/m/v1/usecases/ports"
)

// ItemInteractor ...
type ItemInteractor struct {
	OutputPort ports.ItemOutputPort
	Repository ports.ItemRepository
}

// NewItemInteractor ...
func NewItemInteractor(outputPort ports.ItemOutputPort, repository ports.ItemRepository) *ItemInteractor {
	return &ItemInteractor{
		OutputPort: outputPort,
		Repository: repository,
	}
}

// AddItem ...
func (u *ItemInteractor) AddItem(ctx context.Context, in *input.Item) error {
	item := &entities.Item{
		Name: in.Name,
	}
	result, err := u.Repository.RegisterItem(ctx, item)
	if err != nil {
		return err
	}

	out := &output.Item{
		Name: result.Name,
	}
	return u.OutputPort.OutputItem(out)
}

// FetchItems ...
func (u *ItemInteractor) FetchItems(ctx context.Context) error {
	items, err := u.Repository.RetrieveItems(ctx)
	if err != nil {
		return u.OutputPort.OutputError(err)
	}
	outputs := make([]*output.Item, len(items))
	for i, v := range items {
		outputs[i] = &output.Item{
			Name: v.Name,
		}
	}
	return u.OutputPort.OutputItems(outputs)
}
