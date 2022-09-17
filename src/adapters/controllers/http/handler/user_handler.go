package handler

import (
	"go-playground/m/v1/src/adapters/controllers/http/handler/request"
	"go-playground/m/v1/src/adapters/controllers/http/handler/response"
	"go-playground/m/v1/src/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

// UserHandler ...
type UserHandler struct {
	usecase.IUserManagementUsecase
}

// NewUserHandler ...
func NewUserHandler(u usecase.IUserManagementUsecase) UserHandler {
	return UserHandler{u}
}

// CreateNewUser ...
func (h UserHandler) CreateNewUser(c echo.Context) error {
	ctx := c.Request().Context()

	req := new(request.UserCreate)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := h.CreateUser(ctx, req.ConvertToUserModel(), req.Amount); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, "OK")
}

// GetUser ...
func (h UserHandler) GetUser(c echo.Context) error {
	ctx := c.Request().Context()

	req := new(request.UserGetByID)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	user, err := h.RetrieveUserByCondition(ctx, req.ID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, response.NewUser(user))
}

// GetUserList ...
func (h UserHandler) GetUserList(c echo.Context) error {
	ctx := c.Request().Context()

	users, err := h.RetrieveUsers(ctx)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, response.NewUsers(users))
}
