package model

import (
	"go-playground/m/v1/src/domain/model/user"

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
	Grade        *Grade
}

// TableName ...
func (User) TableName() string {
	return "users"
}

func (u User) makeGeneralUser() (user.General, error) {
	generalUser, err := user.InitGeneral(u.FirstName, u.LastName, u.Age, u.EmailAddress)
	if err != nil {
		return user.General{}, err
	}
	generalUser.SetID(u.ID)
	if u.Grade != nil {
		generalUser.SetGrade(u.Grade.makeGradeEntity())
	}
	return *generalUser, err
}

// InitUser ...
func InitUser(u user.General) *User {
	user := User{
		FirstName:    string(u.FirstName()),
		LastName:     string(u.LastName()),
		Age:          uint(u.Age()),
		EmailAddress: string(u.EmailAddress()),
		GradeID:      uint(u.GradeID()),
	}
	return &user
}

// MakeUserFetchDTO ...
func MakeUserFetchDTO(u User) (*user.FetchDTO, error) {
	generalUser, err := u.makeGeneralUser()
	if err != nil {
		return nil, err
	}
	userFetchDTO := user.NewFetchDTO(generalUser)
	return &userFetchDTO, nil
}

// Users ...
type Users []User

func (us Users) makeGeneralUsers() (user.Generals, error) {
	var err error

	generalUsers := make(user.Generals, len(us))
	for i, u := range us {
		generalUser, err := u.makeGeneralUser()
		if err != nil {
			break
		}
		generalUsers[i] = generalUser
	}
	if err != nil {
		return nil, err
	}
	return generalUsers, nil
}

// MakeUserFetchAllDTO ...
func MakeUserFetchAllDTO(us Users) (*user.FetchAllDTO, error) {
	generalUsers, err := us.makeGeneralUsers()
	if err != nil {
		return nil, err
	}
	fetchAllDTO := user.NewFetchAllDTO(generalUsers)
	return &fetchAllDTO, nil
}
