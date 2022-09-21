package repository

import (
	"context"
	"errors"
	dbModel "go-playground/m/v1/adapters/gateways/persistance/rdb/model"
	"go-playground/m/v1/adapters/gateways/repository/rdb"
	"go-playground/m/v1/usecase/repository/dto"
)

// UserRepository ...
type UserRepository struct {
	rdb.IManageDBConn
}

// NewUserRepository ...
func NewUserRepository(mdc rdb.IManageDBConn) UserRepository {
	return UserRepository{mdc}
}

// RegisterUser ...
func (r UserRepository) RegisterUser(ctx context.Context, dto dto.RegisterUser) (*dto.FetchUserResult, error) {
	userDBModel := dbModel.ConvertToUser(dto)
	if err := r.GetConnection(ctx).Create(&userDBModel).Error; err != nil {
		return nil, err
	}
	return dbModel.MakeFetchUserResultDTO(userDBModel), nil
}

// UpdateUser ...
func (r UserRepository) UpdateUser(ctx context.Context, id uint, dto dto.UpdateUser) (*dto.FetchUserResult, error) {
	userDBModel := dbModel.ConvertToUpdateUser(dto)
	result := r.GetConnection(ctx).Where("id = ?", id).Updates(&userDBModel)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("レコードが更新されませんでした。")
	}
	return dbModel.MakeFetchUserResultDTO(userDBModel), nil
}

// FetchUserByID ...
func (r UserRepository) FetchUserByID(ctx context.Context, id uint) (*dto.FetchUserResult, error) {
	return r.fetchByID(ctx, id)
}

func (r UserRepository) fetchByID(ctx context.Context, id uint) (*dto.FetchUserResult, error) {
	userDBModel := new(dbModel.User)
	if err := r.GetConnection(ctx).Preload("Grade").Where("id = ?", id).Limit(1).Find(&userDBModel).Error; err != nil {
		return nil, err
	}
	return dbModel.MakeFetchUserResultDTO(*userDBModel), nil
}

// FetchUserByEmail ...
func (r UserRepository) FetchUserByEmail(ctx context.Context, email string) (*dto.FetchUserResult, error) {
	userDBModel := new(dbModel.User)
	if err := r.GetConnection(ctx).Where("email_address = ?", email).Limit(1).Find(&userDBModel).Error; err != nil {
		return nil, err
	}
	return dbModel.MakeFetchUserResultDTO(*userDBModel), nil
}

// FetchUserList ...
func (r UserRepository) FetchUserList(ctx context.Context) (*dto.FetchUserListResult, error) {
	return r.fetchUserList(ctx)
}

func (r UserRepository) fetchUserList(ctx context.Context) (*dto.FetchUserListResult, error) {
	usersDBModel := new(dbModel.Users)
	if err := r.GetConnection(ctx).Preload("Grade").Find(&usersDBModel).Error; err != nil {
		return nil, err
	}
	return dbModel.MakeFetchUserListResultDTO(*usersDBModel), nil
}
