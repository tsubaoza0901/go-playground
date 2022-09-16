package repository

import (
	"context"
	dbModel "go-playground/m/v1/src/adapters/gateways/persistance/rdb/model"
	"go-playground/m/v1/src/domain/model/deal"

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
func (r DealHistoryRepository) RegisterDealHistory(ctx context.Context, dto deal.RegisterHistoryDTO) error {
	tx, ok := getTxFromContext(ctx)
	if !ok {
		tx = r.dbConn
	}
	dealHistory := dbModel.NewCreateDealHistory(dto.History, dto.UserID)
	if err := tx.Create(&dealHistory).Error; err != nil {
		return err
	}
	return nil
}

// FetchDealHistoriesByUserID ...
func (r DealHistoryRepository) FetchDealHistoriesByUserID(ctx context.Context, userID uint) (*deal.FetchHistoriesDTO, error) {
	dealHistories := new(dbModel.DealHistories)
	if err := r.dbConn.Where("user_id = ?", userID).Find(&dealHistories).Error; err != nil {
		return nil, err
	}
	return dbModel.MakeFetchHistoriesDTO(*dealHistories), nil
}
