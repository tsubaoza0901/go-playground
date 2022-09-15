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
	db := r.dbConn
	return r.registerDealHistory(ctx, db, dto.History, dto.UserID)
}

// RegisterDealHistoryWithTx ...
func (r DealHistoryRepository) RegisterDealHistoryWithTx(ctx context.Context, tx *gorm.DB, dto deal.RegisterHistoryDTO) error {
	return r.registerDealHistory(ctx, tx, dto.History, dto.UserID)
}

func (r DealHistoryRepository) registerDealHistory(ctx context.Context, db *gorm.DB, history deal.History, userID uint) error {
	dealHistory := dbModel.NewCreateDealHistory(history, userID)
	if err := db.Create(&dealHistory).Error; err != nil {
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
