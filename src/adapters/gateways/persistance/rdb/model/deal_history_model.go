package model

import (
	"go-playground/m/v1/src/domain/model/deal"
	"go-playground/m/v1/src/usecase/repository/dto"

	"gorm.io/gorm"
)

// DealHistory ...
type DealHistory struct {
	gorm.Model
	UserID   uint
	ItemName string
	Amount   uint
}

// TableName ...
func (DealHistory) TableName() string {
	return "deal_histories"
}

// ConvertToDealHistory ...
func ConvertToDealHistory(userID uint, itemName string, amount uint) DealHistory {
	return DealHistory{
		UserID:   userID,
		ItemName: itemName,
		Amount:   amount,
	}
}

// MakeFetchHistoryResultDTO ...
func MakeFetchHistoryResultDTO(dh DealHistory) *dto.FetchDealHistoryResult {
	return dto.NewFetchDealHistoryResult(
		deal.CreatedAt(dh.CreatedAt),
		deal.ItemName(dh.ItemName),
		deal.Amount(dh.Amount),
	)
}

// DealHistories ...
type DealHistories []DealHistory

// MakeFetchHistoryListResultDTO ...
func MakeFetchHistoryListResultDTO(dhs DealHistories) *dto.FetchDealHistoryListResult {
	fetchDealHistories := make(dto.FetchDealHistoryListResult, len(dhs))
	for i, dh := range dhs {
		fetchDealHistories[i] = MakeFetchHistoryResultDTO(dh)
	}
	return &fetchDealHistories
}
