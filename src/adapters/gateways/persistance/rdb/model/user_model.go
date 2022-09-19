package model

import (
	"go-playground/m/v1/usecase/repository/dto"

	"gorm.io/gorm"
)

// User ...
type User struct {
	gorm.Model
	FirstName    string
	LastName     string
	Age          uint
	EmailAddress string
	GradeID      uint
	Grade        Grade
}

// TableName ...
func (User) TableName() string {
	return "users"
}

// ConvertToUser ...
func ConvertToUser(u dto.RegisterUser) User {
	return User{
		FirstName:    u.FirstName,
		LastName:     u.LastName,
		Age:          u.Age,
		EmailAddress: u.EmailAddress,
		GradeID:      u.GradeID,
	}
}

// ConvertToUpdateUser ...
func ConvertToUpdateUser(u dto.UpdateUser) User {
	return User{
		LastName:     u.LastName,
		EmailAddress: u.EmailAddress,
		GradeID:      u.GradeID,
	}
}

// MakeFetchUserResultDTO ...
func MakeFetchUserResultDTO(u User) *dto.FetchUserResult {
	gradeResultDTO := MakeFetchGradeResultDTO(u.Grade)
	fetchUserResult := dto.NewFetchUserResult(
		u.ID,
		u.FirstName,
		u.LastName,
		u.Age,
		u.EmailAddress,
		*gradeResultDTO,
	)
	return fetchUserResult
}

// Users ...
type Users []User

// MakeFetchUserListResultDTO ...
func MakeFetchUserListResultDTO(us Users) *dto.FetchUserListResult {
	fetchUsers := make(dto.FetchUserListResult, len(us))
	for i, u := range us {
		fetchUsers[i] = MakeFetchUserResultDTO(u)
	}
	return &fetchUsers
}
