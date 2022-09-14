package response

import (
	"go-playground/m/v1/src/usecase/data/output"
)

// DealHistory ...
type DealHistory struct {
	CreatedAt string `json:"created_at"`
	ItemName  string `json:"item_name"`
	Amount    uint   `json:"amount"`
}

// NewDealHistory ...
func NewDealHistory(p output.DealHistory) DealHistory {
	return DealHistory{
		CreatedAt: p.CreatedAt.ToFormatedString(),
		ItemName:  p.ItemName,
		Amount:    p.Amount,
	}
}

// DealHistories ...
type DealHistories []DealHistory

// NewDealHistories ...
func NewDealHistories(phs []output.DealHistory) DealHistories {
	dealHistories := make([]DealHistory, len(phs))
	for i, ph := range phs {
		dealHistories[i] = NewDealHistory(ph)
	}
	return dealHistories
}
