package handler

import (
	"go-playground/m/v1/controllers"
	"go-playground/m/v1/controllers/request"
	"go-playground/m/v1/presenters"
	"net/http"

	"github.com/labstack/echo/v4"
)

// User ...
type User struct {
	userController *controllers.User
	*presenters.User
}

// NewUser ...
func NewUser(uc *controllers.User, up *presenters.User) *User {
	return &User{uc, up}
}

// CreateUser ...
func (h *User) CreateUser(c echo.Context) error {
	ctx := c.Request().Context()

	req := new(request.User)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := h.userController.CreateUser(ctx, req); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(h.AppResponse.Status, h.AppResponse.Body)
}

// GetUserByID ...
func (h *User) GetUserByID(c echo.Context) error {
	ctx := c.Request().Context()

	req := new(request.GetUserByID)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := h.userController.GetUserByID(ctx, req.ID); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(h.AppResponse.Status, h.AppResponse.Body)
}

// GetUsers ...
func (h *User) GetUsers(c echo.Context) error {
	ctx := c.Request().Context()

	if err := h.userController.GetUsers(ctx); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(h.AppResponse.Status, h.AppResponse.Body)
}
