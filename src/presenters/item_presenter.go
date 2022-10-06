package presenters

import (
	"go-playground/m/v1/usecases/data/output"
	"net/http"

	"github.com/labstack/echo/v4"
)

// ItemPresenter ...
type ItemPresenter struct {
	c echo.Context
}

// NewItemPresenter ...
func NewItemPresenter(c echo.Context) *ItemPresenter {
	return &ItemPresenter{c}
}

// OutputItems ...
func (u *ItemPresenter) OutputItems(items []*output.Item) error {
	return u.c.JSON(http.StatusOK, items)
}

// OutputItem ...
func (u *ItemPresenter) OutputItem(item *output.Item) error {
	return u.c.JSON(http.StatusOK, item)
}

// OutputError ...
func (u *ItemPresenter) OutputError(err error) error {
	return u.c.JSON(http.StatusInternalServerError, err)
}
