package controllers

import (
	"context"
	"go-playground/m/v1/controllers/request"
	"go-playground/m/v1/usecases/data/input"
	"go-playground/m/v1/usecases/ports"
)

// User ...
type User struct {
	userInportPort ports.UserInportPort
}

// NewUser ...
func NewUser(userInportPort ports.UserInportPort) *User {
	return &User{userInportPort}
}

// CreateUser ...
func (c *User) CreateUser(ctx context.Context, req *request.User) {
	user := &input.User{
		Name: req.Name,
		Age:  req.Age,
	}

	c.userInportPort.AddUser(ctx, user)
}

// GetUserByID ...
func (c *User) GetUserByID(ctx context.Context, id uint) {
	c.userInportPort.FetchUserByID(ctx, id)
}

// GetUsers ...
func (c *User) GetUsers(ctx context.Context) {
	c.userInportPort.FetchUsers(ctx)
}
