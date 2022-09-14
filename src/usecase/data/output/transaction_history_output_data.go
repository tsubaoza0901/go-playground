package output

import (
	"go-playground/m/v1/src/domain/model/transaction"
	"time"
)

// CreatedAt ...
type CreatedAt time.Time

// ToFormatedString ...
func (c CreatedAt) ToFormatedString() string {
	const layout = "2006/01/02"
	return time.Time(c).Format(layout)
}

// TransactionHistory ...
type TransactionHistory struct {
	CreatedAt CreatedAt
	ItemName  string
	Amount    uint
}

// MakeTransactionHistory ...
func MakeTransactionHistory(th transaction.History) TransactionHistory {
	return TransactionHistory{
		CreatedAt: CreatedAt(th.CreatedAt()),
		ItemName:  string(th.ItemName()),
		Amount:    uint(th.Amount()),
	}
}

// TransactionHistories ...
type TransactionHistories []TransactionHistory

// MakeTransactionHistories ...
func MakeTransactionHistories(ths transaction.Histories) TransactionHistories {
	transactionHistories := make(TransactionHistories, len(ths))
	for i, th := range ths {
		transactionHistories[i] = MakeTransactionHistory(th)
	}
	return transactionHistories
}
