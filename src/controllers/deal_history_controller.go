package controllers

import (
	"context"
	"go-playground/m/v1/controllers/request"
	"go-playground/m/v1/usecase"
)

// DealHistory ...
type DealHistory struct {
	dealUsecase usecase.IDealUsecase
}

// NewDealHistory ...
func NewDealHistory(du usecase.IDealUsecase) *DealHistory {
	return &DealHistory{du}
}

// GetDealHistoryList ...
func (c *DealHistory) GetDealHistoryList(ctx context.Context, req *request.RetrieveDealHistories) {
	c.dealUsecase.RetrieveDealHistoriesByUserID(ctx, req.UserID)
}
