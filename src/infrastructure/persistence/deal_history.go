package persistence

import (
	"context"
	"go-playground/m/v1/infrastructure/rdb"
	dbModel "go-playground/m/v1/infrastructure/rdb/model"
	"go-playground/m/v1/usecase/dto"
)

// DealHistory ...
type DealHistory struct {
	rdb.ManageDBConn
}

// NewDealHistory ...
func NewDealHistory(mdc rdb.ManageDBConn) DealHistory {
	return DealHistory{mdc}
}

// RegisterDealHistory ...
func (r DealHistory) RegisterDealHistory(ctx context.Context, dto dto.RegisterDealHistory) error {
	dealHistory := dbModel.ConvertToDealHistory(dto.UserID, dto.ItemName, dto.Amount)
	if err := r.GetConnection(ctx).Create(&dealHistory).Error; err != nil {
		return err
	}
	return nil
}

// FetchDealHistoriesByUserID ...
func (r DealHistory) FetchDealHistoriesByUserID(ctx context.Context, userID uint) (*dto.FetchDealHistoryListResult, error) {
	dealHistories := new(dbModel.DealHistories)
	if err := r.GetConnection(ctx).Where("user_id = ?", userID).Find(&dealHistories).Error; err != nil {
		return nil, err
	}
	return dbModel.MakeFetchHistoryListResultDTO(*dealHistories), nil
}
