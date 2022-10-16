package repository

import (
	"context"
	"errors"
	dbModel "go-playground/m/v1/infrastructure/rdb/model"
	"go-playground/m/v1/repository/rdb"
	"go-playground/m/v1/usecase/dto"
)

// BalanceRepository ...
type BalanceRepository struct {
	rdb.IManageDBConn
}

// NewBalanceRepository ...
func NewBalanceRepository(mdc rdb.IManageDBConn) BalanceRepository {
	return BalanceRepository{mdc}
}

// RegisterBalance ...
func (r BalanceRepository) RegisterBalance(ctx context.Context, dto dto.RegisterBalance) error {
	balanceDBModel := dbModel.ConvertToBalance(dto.UserID, dto.RemainingAmount)
	if err := r.GetConnection(ctx).Create(&balanceDBModel).Error; err != nil {
		return err
	}
	return nil
}

// UpdateBalance ...
func (r BalanceRepository) UpdateBalance(ctx context.Context, dto dto.UpdateBalance) error {
	result := r.GetConnection(ctx).Model(&dbModel.Balance{}).Where("user_id = ?", dto.UserID).Update("amount", dto.RemainingAmount)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("レコードが更新されませんでした。")
	}

	return nil
}

// FetchBalanceByUserID ...
func (r BalanceRepository) FetchBalanceByUserID(ctx context.Context, userID uint) (*dto.FetchBlanceResult, error) {
	return r.fetchBy(ctx, userID)
}

func (r BalanceRepository) fetchBy(ctx context.Context, userID uint) (*dto.FetchBlanceResult, error) {
	balanceDBModel := new(dbModel.Balance)
	if err := r.GetConnection(ctx).Where("user_id = ?", userID).Limit(1).Find(&balanceDBModel).Error; err != nil {
		return nil, err
	}
	return dbModel.MakeFetchBlanceResultDTO(*balanceDBModel), nil
}
