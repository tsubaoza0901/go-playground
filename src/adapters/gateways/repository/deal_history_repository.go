package repository

import (
	"context"
	"go-playground/m/v1/src/adapters/gateways/persistance/rdb"
	dbModel "go-playground/m/v1/src/adapters/gateways/persistance/rdb/model"
	"go-playground/m/v1/src/usecase/repository/dto"
)

// DealHistoryRepository ...
type DealHistoryRepository struct {
	rdb.IManageDBConn
}

// NewDealHistoryRepository ...
func NewDealHistoryRepository(mdc rdb.IManageDBConn) DealHistoryRepository {
	return DealHistoryRepository{mdc}
}

// RegisterDealHistory ...
func (r DealHistoryRepository) RegisterDealHistory(ctx context.Context, dto dto.RegisterDealHistory) error {
	dealHistory := dbModel.ConvertToDealHistory(dto.UserID, dto.ItemName, dto.Amount)
	if err := r.GetConnection(ctx).Create(&dealHistory).Error; err != nil {
		return err
	}
	return nil
}

// FetchDealHistoriesByUserID ...
func (r DealHistoryRepository) FetchDealHistoriesByUserID(ctx context.Context, userID uint) (*dto.FetchDealHistoryListResult, error) {
	dealHistories := new(dbModel.DealHistories)
	if err := r.GetConnection(ctx).Where("user_id = ?", userID).Find(&dealHistories).Error; err != nil {
		return nil, err
	}
	return dbModel.MakeFetchHistoryListResultDTO(*dealHistories), nil
}
