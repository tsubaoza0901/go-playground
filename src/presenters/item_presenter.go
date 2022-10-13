package presenters

import (
	"go-playground/m/v1/presenters/response"
	"go-playground/m/v1/usecases/data/output"
	"net/http"
)

// Item ...
type Item struct {
	*response.AppResponse
}

// NewItem ...
func NewItem() *Item {
	return &Item{}
}

// ItemList ...
func (p *Item) ItemList(items []*output.Item) {
	body := make([]*response.Item, len(items))
	for i, v := range items {
		body[i] = &response.Item{
			Name: v.Name,
		}
	}

	p.AppResponse = &response.AppResponse{
		Status: http.StatusOK,
		Body:   body,
	}
}

// Item ...
func (p *Item) Item(item *output.Item) {
	body := response.Item{
		Name: item.Name,
	}

	p.AppResponse = &response.AppResponse{
		Status: http.StatusOK,
		Body:   body,
	}
}

// Error ...
func (p *Item) Error(err error) {
	p.AppResponse = &response.AppResponse{
		Status: http.StatusOK,
		Body:   err.Error(),
	}
}
