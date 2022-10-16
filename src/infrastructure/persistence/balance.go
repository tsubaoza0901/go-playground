package persistence

import (
	"context"
	"errors"
	"go-playground/m/v1/infrastructure/rdb"
	dbModel "go-playground/m/v1/infrastructure/rdb/model"
	"go-playground/m/v1/usecase/dto"
)

// Balance ...
type Balance struct {
	rdb.ManageDBConn
}

// NewBalance ...
func NewBalance(mdc rdb.ManageDBConn) Balance {
	return Balance{mdc}
}

// RegisterBalance ...
func (r Balance) RegisterBalance(ctx context.Context, dto dto.RegisterBalance) error {
	balanceDBModel := dbModel.ConvertToBalance(dto.UserID, dto.RemainingAmount)
	if err := r.GetConnection(ctx).Create(&balanceDBModel).Error; err != nil {
		return err
	}
	return nil
}

// UpdateBalance ...
func (r Balance) UpdateBalance(ctx context.Context, dto dto.UpdateBalance) error {
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
func (r Balance) FetchBalanceByUserID(ctx context.Context, userID uint) (*dto.FetchBlanceResult, error) {
	return r.fetchBy(ctx, userID)
}

func (r Balance) fetchBy(ctx context.Context, userID uint) (*dto.FetchBlanceResult, error) {
	balanceDBModel := new(dbModel.Balance)
	if err := r.GetConnection(ctx).Where("user_id = ?", userID).Limit(1).Find(&balanceDBModel).Error; err != nil {
		return nil, err
	}
	return dbModel.MakeFetchBlanceResultDTO(*balanceDBModel), nil
}
