package repository

import (
	"context"
	dbModel "go-playground/m/v1/src/adapters/gateways/persistance/rdb/model"
	"go-playground/m/v1/src/usecase/repository/dto"

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
func (r UserRepository) RegisterUser(ctx context.Context, dto dto.RegisterUser) (*dto.FetchUserResult, error) {
	tx, ok := getTxFromContext(ctx)
	if !ok {
		tx = r.dbConn
	}
	userDBModel := dbModel.ConvertToUser(dto)
	if err := tx.Create(&userDBModel).Error; err != nil {
		return nil, err
	}
	return dbModel.MakeFetchUserResultDTO(userDBModel), nil
}

// FetchUserByID ...
func (r UserRepository) FetchUserByID(ctx context.Context, id uint) (*dto.FetchUserResult, error) {
	return r.fetchByID(id)
}

func (r UserRepository) fetchByID(id uint) (*dto.FetchUserResult, error) {
	userDBModel := new(dbModel.User)
	if err := r.dbConn.Preload("Grade").Where("id = ?", id).First(&userDBModel).Error; err != nil {
		return nil, err
	}
	return dbModel.MakeFetchUserResultDTO(*userDBModel), nil
}

// FetchUserList ...
func (r UserRepository) FetchUserList(ctx context.Context) (*dto.FetchUserListResult, error) {
	return r.fetchUserList()
}

func (r UserRepository) fetchUserList() (*dto.FetchUserListResult, error) {
	usersDBModel := new(dbModel.Users)
	if err := r.dbConn.Preload("Grade").Find(&usersDBModel).Error; err != nil {
		return nil, err
	}
	return dbModel.MakeFetchUserListResultDTO(*usersDBModel), nil
}

// CountTheNumberOfUsersByEmail ...
func (r UserRepository) CountTheNumberOfUsersByEmail(ctx context.Context, email string) (uint, error) {
	var count int64
	if err := r.dbConn.Model(&dbModel.User{}).Where("email_address = ?", email).Count(&count).Error; err != nil {
		return 0, err
	}
	return uint(count), nil
}
