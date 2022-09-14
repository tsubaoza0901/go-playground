package repository

import (
	"context"
	dbModel "go-playground/m/v1/src/adapters/gateways/persistance/rdb/model"
	"go-playground/m/v1/src/domain/model/user"

	"gorm.io/gorm"
)

// UserRepository ...
type UserRepository struct {
	dbConn *gorm.DB
}

// NewUserRepository ...
func NewUserRepository(dbConn *gorm.DB) UserRepository {
	return UserRepository{dbConn}
}

// RegisterUser ...
func (r UserRepository) RegisterUser(ctx context.Context, u user.RegistrationDTO) (*user.FetchDTO, error) {
	userDBModel := dbModel.InitUser(u.General)
	if err := r.dbConn.Create(&userDBModel).Error; err != nil {
		return nil, err
	}
	userFetchDTO, err := dbModel.MakeUserFetchDTO(*userDBModel)
	if err != nil {
		return nil, err
	}
	return userFetchDTO, nil
}

// FetchUser ...
func (r UserRepository) FetchUser(ctx context.Context, id uint) (*user.FetchDTO, error) {
	return r.fetchByID(id)
}

func (r UserRepository) fetchByID(id uint) (*user.FetchDTO, error) {
	userDBModel := new(dbModel.User)
	if err := r.dbConn.Preload("Grade").Where("id = ?", id).First(&userDBModel).Error; err != nil {
		return nil, err
	}
	userFetchDTO, err := dbModel.MakeUserFetchDTO(*userDBModel)
	if err != nil {
		return nil, err
	}
	return userFetchDTO, nil
}

// FetchAllUsers ...
func (r UserRepository) FetchAllUsers(ctx context.Context) (*user.FetchAllDTO, error) {
	return r.fetchAllUsers()
}

func (r UserRepository) fetchAllUsers() (*user.FetchAllDTO, error) {
	usersDBModel := new(dbModel.Users)
	if err := r.dbConn.Preload("Grade").Find(&usersDBModel).Error; err != nil {
		return nil, err
	}
	userFetchAllDTO, err := dbModel.MakeUserFetchAllDTO(*usersDBModel)
	if err != nil {
		return nil, err
	}
	return userFetchAllDTO, nil
}

// CountTheNumberOfUsersByEmail ...
func (r UserRepository) CountTheNumberOfUsersByEmail(ctx context.Context, email user.EmailAddress) (uint, error) {
	var count int64
	if err := r.dbConn.Model(&dbModel.User{}).Where("email_address = ?", email).Count(&count).Error; err != nil {
		return 0, err
	}
	return uint(count), nil
}
