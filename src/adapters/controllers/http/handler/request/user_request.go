package request

import (
	"go-playground/m/v1/usecase/data/input"
)

// UserCreate ...
type UserCreate struct {
	FirstName    string `json:"firstName" validate:"required"`
	LastName     string `json:"lastName" validate:"required"`
	Age          uint   `json:"age" validate:"required"`
	Amount       uint   `json:"amount" validate:"required"`
	EmailAddress string `json:"email" validate:"required"`
}

// ConvertToUserModel ...
func (u UserCreate) ConvertToUserModel() input.UserCreate {
	user := input.NewUserCreate()
	user.FirstName = u.FirstName
	user.LastName = u.LastName
	user.Age = u.Age
	user.EmailAddress = u.EmailAddress
	return user
}

// UserGetByID ...
type UserGetByID struct {
	ID uint `param:"id" validate:"required"`
}
