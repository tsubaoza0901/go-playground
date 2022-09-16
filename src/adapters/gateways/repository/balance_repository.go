package repository

import (
	"context"
	dbModel "go-playground/m/v1/src/adapters/gateways/persistance/rdb/model"
	"go-playground/m/v1/src/usecase/repository/dto"

	"gorm.io/gorm"
)

// BalanceRepository ...
type BalanceRepository struct {
	dbConn *gorm.DB
}

// NewBalanceRepository ...
func NewBalanceRepository(conn *gorm.DB) BalanceRepository {
	return BalanceRepository{
		dbConn: conn,
	}
}

// RegisterBalance ...
func (r BalanceRepository) RegisterBalance(ctx context.Context, userID uint, dto dto.RegisterBalance) error {
	tx, ok := getTxFromContext(ctx)
	if !ok {
		tx = r.dbConn
	}
	balanceDBModel := dbModel.ConvertToBalance(userID, dto.RemainingAmount)
	if err := tx.Create(&balanceDBModel).Error; err != nil {
		return err
	}
	return nil
}

// UpdateBalance ...
func (r BalanceRepository) UpdateBalance(ctx context.Context, userID uint, dto dto.UpdateBalance) error {
	tx, ok := getTxFromContext(ctx)
	if !ok {
		tx = r.dbConn
	}
	if err := tx.Model(&dbModel.Balance{}).Where("user_id = ?", userID).Update("amount", dto.RemainingAmount).Error; err != nil {
		return err
	}
	return nil
}

// FetchBalanceByUserID ...
func (r BalanceRepository) FetchBalanceByUserID(ctx context.Context, userID uint) (*dto.FetchBlanceResult, error) {
	return r.fetchBy(ctx, userID)
}

func (r BalanceRepository) fetchBy(ctx context.Context, userID uint) (*dto.FetchBlanceResult, error) {
	balanceDBModel := new(dbModel.Balance)
	if err := r.dbConn.Where("user_id = ?", userID).First(&balanceDBModel).Error; err != nil {
		return nil, err
	}
	return dbModel.MakeFetchBlanceResultDTO(*balanceDBModel), nil
}
