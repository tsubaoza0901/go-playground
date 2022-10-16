package repository

import (
	"context"
	"go-playground/m/v1/repository/persistence"
	"go-playground/m/v1/usecase/dto"
)

// DealHistory ...
type DealHistory struct {
	dealHistoryDataAccess persistence.DealHistoryDataAccess
}

// NewDealHistory ...
func NewDealHistory(dhda persistence.DealHistoryDataAccess) DealHistory {
	return DealHistory{dhda}
}

// RegisterDealHistory ...
func (r DealHistory) RegisterDealHistory(ctx context.Context, dto dto.RegisterDealHistory) error {
	return r.dealHistoryDataAccess.RegisterDealHistory(ctx, dto)
}

// FetchDealHistoriesByUserID ...
func (r DealHistory) FetchDealHistoriesByUserID(ctx context.Context, userID uint) (*dto.FetchDealHistoryListResult, error) {
	return r.dealHistoryDataAccess.FetchDealHistoriesByUserID(ctx, userID)
}
