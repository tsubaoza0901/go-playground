package output

import (
	"go-playground/m/v1/src/domain/model/deal"
	"time"
)

// CreatedAt ...
type CreatedAt time.Time

// ToFormatedString ...
func (c CreatedAt) ToFormatedString() string {
	const layout = "2006/01/02"
	return time.Time(c).Format(layout)
}

// DealHistory ...
type DealHistory struct {
	CreatedAt CreatedAt
	ItemName  string
	Amount    uint
}

// MakeDealHistory ...
func MakeDealHistory(th deal.History) DealHistory {
	return DealHistory{
		CreatedAt: CreatedAt(th.CreatedAt()),
		ItemName:  string(th.ItemName()),
		Amount:    uint(th.Amount()),
	}
}

// DealHistories ...
type DealHistories []DealHistory

// MakeDealHistories ...
func MakeDealHistories(ths deal.Histories) DealHistories {
	dealHistories := make(DealHistories, len(ths))
	for i, th := range ths {
		dealHistories[i] = MakeDealHistory(th)
	}
	return dealHistories
}
