package request

import (
	"go-playground/m/v1/usecase/data/input"
)

// UserCreate ...
type UserCreate struct {
	FirstName    string `json:"firstName" validate:"required"`
	LastName     string `json:"lastName" validate:"required"`
	Age          uint   `json:"age" validate:"required"`
	EmailAddress string `json:"email" validate:"required"`
	TopUpAmount  uint   `json:"topUpAmount" validate:"required"`
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

// UserUpdate ...
type UserUpdate struct {
	ID           uint   `param:"id" validate:"required"`
	LastName     string `json:"lastName"`
	EmailAddress string `json:"email"`
	GradeID      uint   `json:"gradeId"`
}

// ConvertToUserModel ...
func (u UserUpdate) ConvertToUserModel() input.UserUpdate {
	user := input.NewUserUpdate()
	user.ID = u.ID
	user.LastName = u.LastName
	user.EmailAddress = u.EmailAddress
	user.GradeID = u.GradeID
	return user
}

// UserGetByID ...
type UserGetByID struct {
	ID uint `param:"id" validate:"required"`
}
