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
func (r BalanceRepository) CreateBalance(ctx context.Context, userID uint, dto balance.CreateBalanceDTO) error {
	tx, ok := getTxFromContext(ctx)
	if !ok {
		tx = r.dbConn
	}
	balanceDBModel := dbModel.InitBalance(userID, dto.RemainingAmount)
	if err := tx.Create(&balanceDBModel).Error; err != nil {
		return err
	}
	return nil
}

// UpdateBalance ...
func (r BalanceRepository) UpdateBalance(ctx context.Context, userID uint, dto balance.UpdateBalanceDTO) error {
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
