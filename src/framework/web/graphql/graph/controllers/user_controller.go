package controllers

import (
	"context"
	"go-playground/m/v1/framework/web/graphql/graph/model"
	"go-playground/m/v1/usecase"
	"go-playground/m/v1/usecase/data/input"
)

// User ...
type User struct {
	userManagementUsecase usecase.IUserManagementUsecase
}

// NewUser ...
func NewUser(umu usecase.IUserManagementUsecase) *User {
	return &User{umu}
}

// CreateNewUser ...
func (h *User) CreateNewUser(ctx context.Context, req model.NewUser) {
	in := input.UserCreate{
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		Age:          uint(req.Age),
		EmailAddress: req.Email,
	}
	h.userManagementUsecase.CreateUser(ctx, in, uint(req.TopUpAmount))
}

// GetUserList ...
func (h *User) GetUserList(ctx context.Context) {
	h.userManagementUsecase.RetrieveUsers(ctx)
}
