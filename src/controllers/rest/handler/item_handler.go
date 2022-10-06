package handler

import (
	"go-playground/m/v1/controllers/rest/handler/request"
	"go-playground/m/v1/injector"
	"go-playground/m/v1/usecases/data/input"
	"go-playground/m/v1/usecases/ports"

	"github.com/labstack/echo/v4"
)

// ItemHandler ...
type ItemHandler struct {
	itemInportPort ports.ItemInportPort
}

// CreateItem ...
func (u *ItemHandler) CreateItem(ud *injector.ItemDependency) func(echo.Context) error {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		req := new(request.Item)
		if err := c.Bind(req); err != nil {
			return err
		}

		item := &input.Item{
			Name: req.Name,
		}
		return u.newInputPort(c, ud).AddItem(ctx, item)
	}
}

// GetItems ...
func (u *ItemHandler) GetItems(ud *injector.ItemDependency) func(echo.Context) error {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		return u.newInputPort(c, ud).FetchItems(ctx)
	}
}

func (u *ItemHandler) newInputPort(c echo.Context, dependency *injector.ItemDependency) ports.ItemInportPort {
	return dependency.ItemInputPortFactory(dependency.ItemOutputPortFactory(c))
}
