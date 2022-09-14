package repository

import (
	"context"
	dbModel "go-playground/m/v1/src/adapters/gateways/persistance/rdb/model"
	"go-playground/m/v1/src/domain/model/balance"

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

// CreateBalance ...
func (r BalanceRepository) CreateBalance(ctx context.Context, userID uint, bl balance.CreateBalanceDTO) error {
	balanceDBModel := dbModel.InitBalance(userID, bl.RemainingAmount)
	if err := r.dbConn.Create(&balanceDBModel).Error; err != nil {
		return err
	}
	return nil
}

// UpdateBalance ...
func (r BalanceRepository) UpdateBalance(ctx context.Context, userID uint, updateBalanceDTO balance.UpdateBalanceDTO) error {
	if err := r.dbConn.Model(&dbModel.Balance{}).Where("user_id = ?", userID).Update("amount", updateBalanceDTO.RemainingAmount).Error; err != nil {
		return err
	}
	return nil
}

// FetchBalanceByUserID ...
func (r BalanceRepository) FetchBalanceByUserID(ctx context.Context, userID uint) (*balance.FetchAmountDTO, error) {
	return r.fetchBy(ctx, userID)
}

func (r BalanceRepository) fetchBy(ctx context.Context, userID uint) (*balance.FetchAmountDTO, error) {
	balanceDBModel := new(dbModel.Balance)
	if err := r.dbConn.Where("user_id = ?", userID).First(&balanceDBModel).Error; err != nil {
		return nil, err
	}
	return dbModel.MakeBalanceFetchAmountDTO(*balanceDBModel), nil
}
