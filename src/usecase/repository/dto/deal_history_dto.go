package dto

import (
	"go-playground/m/v1/src/domain/model/deal"
	"go-playground/m/v1/src/domain/model/user"
	"time"
)

// RegisterDealHistory 取引履歴登録用DTO
type RegisterDealHistory struct {
	UserID   uint
	ItemName string
	Amount   uint
}

// NewRegisterDealHistory ...
func NewRegisterDealHistory(userID user.ID, itemName string, amount uint) RegisterDealHistory {
	return RegisterDealHistory{
		UserID:   uint(userID),
		ItemName: string(itemName),
		Amount:   uint(amount),
	}
}

// FetchDealHistoryResult 取引履歴確認用DTO
type FetchDealHistoryResult struct {
	CreatedAt time.Time
	ItemName  string
	Amount    uint
	// Userオブジェクト
}

// NewFetchDealHistoryResult ...
func NewFetchDealHistoryResult(createdAt time.Time, itemName string, amount uint) *FetchDealHistoryResult {
	return &FetchDealHistoryResult{
		CreatedAt: createdAt,
		ItemName:  itemName,
		Amount:    amount,
	}
}

// ToDealHistoryModel ...
func (d FetchDealHistoryResult) ToDealHistoryModel() deal.History {
	dealHistory := deal.MakeHistory(
		deal.CreatedAt(d.CreatedAt),
		deal.ItemName(d.ItemName),
		deal.Amount(d.Amount),
	)
	return *dealHistory
}

// FetchDealHistoryListResult 取引履歴確認用DTO
type FetchDealHistoryListResult []*FetchDealHistoryResult

// ToDealHistoriesModel ...
func (ds FetchDealHistoryListResult) ToDealHistoriesModel() deal.Histories {
	dealHistories := make(deal.Histories, len(ds))
	for i, d := range ds {
		dealHistories[i] = d.ToDealHistoryModel()
	}
	return dealHistories
}
