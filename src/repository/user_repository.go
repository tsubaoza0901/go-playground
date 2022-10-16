package repository

import (
	"context"
	"go-playground/m/v1/repository/persistence"
	"go-playground/m/v1/usecase/dto"
)

// User ...
type User struct {
	userDataAccess persistence.UserDataAccess
}

// NewUser ...
func NewUser(uda persistence.UserDataAccess) User {
	return User{uda}
}

// RegisterUser ...
func (r User) RegisterUser(ctx context.Context, dto dto.RegisterUser) (*dto.FetchUserResult, error) {
	return r.userDataAccess.RegisterUser(ctx, dto)
}

// UpdateUser ...
func (r User) UpdateUser(ctx context.Context, id uint, dto dto.UpdateUser) (*dto.FetchUserResult, error) {
	return r.userDataAccess.UpdateUser(ctx, id, dto)
}

// FetchUserByID ...
func (r User) FetchUserByID(ctx context.Context, id uint) (*dto.FetchUserResult, error) {
	return r.userDataAccess.FetchUserByID(ctx, id)
}

// FetchUserByEmail ...
func (r User) FetchUserByEmail(ctx context.Context, email string) (*dto.FetchUserResult, error) {
	return r.userDataAccess.FetchUserByEmail(ctx, email)
}

// FetchUserList ...
func (r User) FetchUserList(ctx context.Context) (*dto.FetchUserListResult, error) {
	return r.userDataAccess.FetchUserList(ctx)
}
