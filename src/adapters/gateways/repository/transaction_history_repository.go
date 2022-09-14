package repository

import (
	"context"
	dbModel "go-playground/m/v1/src/adapters/gateways/persistance/rdb/model"
	"go-playground/m/v1/src/domain/model/transaction"

	"gorm.io/gorm"
)

// TransactionHistoryRepository ...
type TransactionHistoryRepository struct {
	dbConn *gorm.DB
}

// NewTransactionHistoryRepository ...
func NewTransactionHistoryRepository(conn *gorm.DB) TransactionHistoryRepository {
	return TransactionHistoryRepository{
		dbConn: conn,
	}
}

// RegisterTransactionHistory ...
func (r TransactionHistoryRepository) RegisterTransactionHistory(ctx context.Context, p transaction.RegisterHistoryDTO) error {
	transactionHistory := dbModel.NewCreateTransactionHistory(p.History, p.UserID)
	if err := r.dbConn.Create(&transactionHistory).Error; err != nil {
		return err
	}
	return nil
}

// FetchTransactionHistoriesByUserID ...
func (r TransactionHistoryRepository) FetchTransactionHistoriesByUserID(ctx context.Context, userID uint) (*transaction.FetchHistoriesDTO, error) {
	transactionHistories := new(dbModel.TransactionHistories)
	if err := r.dbConn.Where("user_id = ?", userID).Find(&transactionHistories).Error; err != nil {
		return nil, err
	}
	return dbModel.MakeFetchHistoriesDTO(*transactionHistories), nil
}
