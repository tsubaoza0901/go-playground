package dto

import (
	"go-playground/m/v1/src/domain/model/grade"
	"go-playground/m/v1/src/domain/model/user"
)

// RegisterUser ...
type RegisterUser struct {
	FirstName    string
	LastName     string
	Age          uint
	EmailAddress string
	GradeID      uint
}

// NewRegisterUser ...
func NewRegisterUser(general user.General) RegisterUser {
	return RegisterUser{
		FirstName:    string(general.FirstName()),
		LastName:     string(general.LastName()),
		Age:          uint(general.Age()),
		EmailAddress: string(general.EmailAddress()),
		GradeID:      uint(general.GradeID()),
	}
}

// FetchUserResult ...
type FetchUserResult struct {
	ID           uint
	FirstName    string
	LastName     string
	Age          uint
	EmailAddress string
	Grade        FetchGradeResult
}

// NewFetchUserResult ...
func NewFetchUserResult(id user.ID, firstName user.FirstName, lastName user.LastName, age user.Age, email user.EmailAddress, grade FetchGradeResult) *FetchUserResult {
	return &FetchUserResult{
		ID:           uint(id),
		FirstName:    string(firstName),
		LastName:     string(lastName),
		Age:          uint(age),
		EmailAddress: string(email),
		Grade: FetchGradeResult{
			ID:   grade.ID,
			Name: grade.Name,
		},
	}
}

// ToGeneralUserModel ...
func (u FetchUserResult) ToGeneralUserModel() (*user.General, error) {
	gradeEntity := grade.MakeEntity(
		grade.ID(u.Grade.ID),
		grade.Name(u.Grade.Name),
	)

	generalUser, err := user.MakeGeneral(
		user.ID(u.ID),
		user.FirstName(u.FirstName),
		user.LastName(u.LastName),
		user.Age(u.Age),
		user.EmailAddress(u.EmailAddress),
		*gradeEntity,
	)
	if err != nil {
		return nil, err
	}
	return generalUser, nil
}

// FetchUserListResult ...
type FetchUserListResult []*FetchUserResult

// ToGeneralUsersModel ...
func (us FetchUserListResult) ToGeneralUsersModel() (user.Generals, error) {
	var err error
	generalUsers := make(user.Generals, len(us))
	for i, u := range us {
		gus, err := u.ToGeneralUserModel()
		if err != nil {
			break
		}
		generalUsers[i] = *gus
	}
	if err != nil {
		return nil, err
	}
	return generalUsers, nil
}
