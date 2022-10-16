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

// CreateNewUser ...
func (h *UserHandler) CreateNewUser(c echo.Context) error {
	ctx := c.Request().Context()

	req := new(request.UserCreate)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	h.userController.CreateNewUser(ctx, req)
	if h.Error != nil {
		return c.JSON(http.StatusBadGateway, h.Error.Text())
	}
	return c.JSON(http.StatusOK, "OK")
}

// UpdateUser ...
func (h *UserHandler) UpdateUser(c echo.Context) error {
	ctx := c.Request().Context()

	req := new(request.UserUpdate)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	h.userController.UpdateUser(ctx, req)
	if h.Error != nil {
		return c.JSON(http.StatusBadGateway, h.Error.Text())
	}
	return c.JSON(http.StatusOK, "OK")
}

// GetUser ...
func (h *UserHandler) GetUser(c echo.Context) error {
	ctx := c.Request().Context()

	req := new(request.UserGetByID)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	h.userController.GetUser(ctx, req)
	if h.Error != nil {
		return c.JSON(http.StatusBadGateway, h.Error.Text())
	}
	return c.JSON(http.StatusOK, h.AppResponse.Data)
}

// GetUserList ...
func (h *UserHandler) GetUserList(c echo.Context) error {
	ctx := c.Request().Context()

	h.userController.GetUserList(ctx)
	if h.Error != nil {
		return c.JSON(http.StatusBadGateway, h.Error.Text())
	}
	return c.JSON(http.StatusOK, h.AppResponse.Data)
}
