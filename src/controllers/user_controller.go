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
func (c *User) CreateUser(ctx context.Context, req *request.User) error {
	user := &input.User{
		Name: req.Name,
		Age:  req.Age,
	}

	c.userInportPort.AddUser(ctx, user)
	return nil
}

// GetUserByID ...
func (c *User) GetUserByID(ctx context.Context, id uint) error {
	c.userInportPort.FetchUserByID(ctx, id)
	return nil
}

// GetUsers ...
func (c *User) GetUsers(ctx context.Context) error {
	c.userInportPort.FetchUsers(ctx)
	return nil
}
