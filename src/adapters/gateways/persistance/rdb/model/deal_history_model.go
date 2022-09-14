package model

import (
	"go-playground/m/v1/src/domain/model/deal"

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

func (th DealHistory) makeDealHistory() deal.History {
	dealHistory := deal.NewHistory(th.ItemName, th.Amount)
	dealHistory.SetCreatedAt(th.CreatedAt)
	return *dealHistory
}

// NewCreateDealHistory ...
func NewCreateDealHistory(p deal.History, userID uint) DealHistory {
	return DealHistory{
		UserID:   userID,
		ItemName: string(p.ItemName()),
		Amount:   uint(p.Amount()),
	}
}

// MakeFetchHistoryDTO ...
func MakeFetchHistoryDTO(th DealHistory) *deal.FetchHistoryDTO {
	return deal.NewFetchHistoryDTO(th.makeDealHistory())
}

// DealHistories ...
type DealHistories []DealHistory

// MakeFetchHistoriesDTO ...
func MakeFetchHistoriesDTO(ths DealHistories) *deal.FetchHistoriesDTO {
	dealHistoryEntities := make(deal.Histories, len(ths))
	for i, th := range ths {
		dealHistoryEntities[i] = th.makeDealHistory()
	}
	return deal.NewFetchHistoriesDTO(dealHistoryEntities)
}
