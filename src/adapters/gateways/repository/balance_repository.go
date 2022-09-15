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
	db := r.dbConn
	return r.createBalance(ctx, db, userID, dto.RemainingAmount)
}

// CreateBalanceWithTx ...
func (r BalanceRepository) CreateBalanceWithTx(ctx context.Context, tx *gorm.DB, userID uint, dto balance.CreateBalanceDTO) error {
	return r.createBalance(ctx, tx, userID, dto.RemainingAmount)
}

func (r BalanceRepository) createBalance(ctx context.Context, db *gorm.DB, userID uint, amount balance.RemainingAmount) error {
	balanceDBModel := dbModel.InitBalance(userID, amount)
	if err := db.Create(&balanceDBModel).Error; err != nil {
		return err
	}
	return nil
}

// UpdateBalance ...
func (r BalanceRepository) UpdateBalance(ctx context.Context, userID uint, dto balance.UpdateBalanceDTO) error {
	db := r.dbConn
	return r.updateBalance(ctx, db, userID, dto.RemainingAmount)
}

// UpdateBalanceWithTx ...
func (r BalanceRepository) UpdateBalanceWithTx(ctx context.Context, tx *gorm.DB, userID uint, dto balance.UpdateBalanceDTO) error {
	return r.updateBalance(ctx, tx, userID, dto.RemainingAmount)
}

func (r BalanceRepository) updateBalance(ctx context.Context, db *gorm.DB, userID uint, amount balance.RemainingAmount) error {
	if err := db.Model(&dbModel.Balance{}).Where("user_id = ?", userID).Update("amount", amount).Error; err != nil {
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
