package rest

import (
	"go-playground/m/v1/controllers"
	"go-playground/m/v1/controllers/request"
	"go-playground/m/v1/presenters"
	"net/http"

	"github.com/labstack/echo/v4"
)

// UserHandler ...
type UserHandler struct {
	userController *controllers.User
	*presenters.User
}

// NewUserHandler ...
func NewUserHandler(uc *controllers.User, up *presenters.User) *UserHandler {
	return &UserHandler{uc, up}
}

// CreateUser ...
func (h *UserHandler) CreateUser(c echo.Context) error {
	ctx := c.Request().Context()

	req := new(request.User)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	h.userController.CreateUser(ctx, req)
	return c.JSON(h.AppResponse.Status, h.AppResponse.Body)
}

// GetUserByID ...
func (h *UserHandler) GetUserByID(c echo.Context) error {
	ctx := c.Request().Context()

	req := new(request.GetUserByID)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	h.userController.GetUserByID(ctx, req.ID)
	return c.JSON(h.AppResponse.Status, h.AppResponse.Body)
}

// GetUsers ...
func (h *UserHandler) GetUsers(c echo.Context) error {
	ctx := c.Request().Context()

	h.userController.GetUsers(ctx)
	return c.JSON(h.AppResponse.Status, h.AppResponse.Body)
}
