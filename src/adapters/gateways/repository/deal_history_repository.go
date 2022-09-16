package repository

import (
	"context"
	dbModel "go-playground/m/v1/src/adapters/gateways/persistance/rdb/model"
	"go-playground/m/v1/src/usecase/repository/dto"

	"gorm.io/gorm"
)

// DealHistoryRepository ...
type DealHistoryRepository struct {
	dbConn *gorm.DB
}

// NewDealHistoryRepository ...
func NewDealHistoryRepository(conn *gorm.DB) DealHistoryRepository {
	return DealHistoryRepository{
		dbConn: conn,
	}
}

// RegisterDealHistory ...
func (r DealHistoryRepository) RegisterDealHistory(ctx context.Context, dto dto.RegisterDealHistory) error {
	tx, ok := getTxFromContext(ctx)
	if !ok {
		tx = r.dbConn
	}
	dealHistory := dbModel.ConvertToDealHistory(dto.UserID, dto.ItemName, dto.Amount)
	if err := tx.Create(&dealHistory).Error; err != nil {
		return err
	}
	return nil
}

// FetchDealHistoriesByUserID ...
func (r DealHistoryRepository) FetchDealHistoriesByUserID(ctx context.Context, userID uint) (*dto.FetchDealHistoryListResult, error) {
	dealHistories := new(dbModel.DealHistories)
	if err := r.dbConn.Where("user_id = ?", userID).Find(&dealHistories).Error; err != nil {
		return nil, err
	}
	return dbModel.MakeFetchHistoryListResultDTO(*dealHistories), nil
}
