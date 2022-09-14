package response

import (
	"go-playground/m/v1/src/usecase/data/output"
)

// TransactionHistory ...
type TransactionHistory struct {
	CreatedAt string `json:"created_at"`
	ItemName  string `json:"item_name"`
	Amount    uint   `json:"amount"`
}

// NewTransactionHistory ...
func NewTransactionHistory(p output.TransactionHistory) TransactionHistory {
	return TransactionHistory{
		CreatedAt: p.CreatedAt.ToFormatedString(),
		ItemName:  p.ItemName,
		Amount:    p.Amount,
	}
}

// TransactionHistories ...
type TransactionHistories []TransactionHistory

// NewTransactionHistories ...
func NewTransactionHistories(phs []output.TransactionHistory) TransactionHistories {
	transactionHistories := make([]TransactionHistory, len(phs))
	for i, ph := range phs {
		transactionHistories[i] = NewTransactionHistory(ph)
	}
	return transactionHistories
}
