package controllers

import (
	"context"
	"go-playground/m/v1/controllers/request"
	"go-playground/m/v1/usecases/data/input"
	"go-playground/m/v1/usecases/ports"
)

// Item ...
type Item struct {
	itemInportPort ports.ItemInportPort
}

// NewItem ...
func NewItem(itemInportPort ports.ItemInportPort) *Item {
	return &Item{itemInportPort}
}

// CreateItem ...
func (u *Item) CreateItem(ctx context.Context, req *request.Item) {
	item := &input.Item{
		Name: req.Name,
	}
	u.itemInportPort.AddItem(ctx, item)
}

// GetItems ...
func (u *Item) GetItems(ctx context.Context) {
	u.itemInportPort.FetchItems(ctx)
}
