package rest

import (
	"go-playground/m/v1/controllers"
	"go-playground/m/v1/controllers/request"
	"go-playground/m/v1/presenters"
	"net/http"

	"github.com/labstack/echo/v4"
)

// ItemHandler ...
type ItemHandler struct {
	itemController *controllers.Item
	*presenters.Item
}

// NewItemHandler ...
func NewItemHandler(ic *controllers.Item, ip *presenters.Item) *ItemHandler {
	return &ItemHandler{ic, ip}
}

// CreateItem ...
func (h *ItemHandler) CreateItem(c echo.Context) error {
	ctx := c.Request().Context()

	req := new(request.Item)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	h.itemController.CreateItem(ctx, req)
	return c.JSON(h.AppResponse.Status, h.AppResponse.Body)
}

// GetItems ...
func (h *ItemHandler) GetItems(c echo.Context) error {
	ctx := c.Request().Context()

	h.itemController.GetItems(ctx)
	return c.JSON(h.AppResponse.Status, h.AppResponse.Body)
}
