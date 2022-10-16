package controllers

import (
	"context"
	"go-playground/m/v1/controllers/request"
	"go-playground/m/v1/usecase"
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
func (h *User) CreateNewUser(ctx context.Context, req *request.UserCreate) {
	h.userManagementUsecase.CreateUser(ctx, req.ConvertToUserModel(), req.TopUpAmount)
}

// UpdateUser ...
func (h *User) UpdateUser(ctx context.Context, req *request.UserUpdate) {
	h.userManagementUsecase.EditUser(ctx, req.ConvertToUserModel())
}

// GetUser ...
func (h *User) GetUser(ctx context.Context, req *request.UserGetByID) {
	h.userManagementUsecase.RetrieveUserByCondition(ctx, req.ID)
}

// GetUserList ...
func (h *User) GetUserList(ctx context.Context) {
	h.userManagementUsecase.RetrieveUsers(ctx)
}
