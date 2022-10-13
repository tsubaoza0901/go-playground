package handler

import (
	"go-playground/m/v1/controllers"
	"go-playground/m/v1/controllers/request"
	"go-playground/m/v1/presenters"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Item ...
type Item struct {
	itemController *controllers.Item
	*presenters.Item
}

// NewItem ...
func NewItem(ic *controllers.Item, ip *presenters.Item) *Item {
	return &Item{ic, ip}
}

// CreateItem ...
func (h *Item) CreateItem(c echo.Context) error {
	ctx := c.Request().Context()

	req := new(request.Item)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := h.itemController.CreateItem(ctx, req); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(h.AppResponse.Status, h.AppResponse.Body)
}

// GetItems ...
func (h *Item) GetItems(c echo.Context) error {
	ctx := c.Request().Context()

	if err := h.itemController.GetItems(ctx); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(h.AppResponse.Status, h.AppResponse.Body)
}
