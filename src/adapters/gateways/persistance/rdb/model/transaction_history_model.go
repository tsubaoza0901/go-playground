package model

import (
	"go-playground/m/v1/src/domain/model/transaction"

	"gorm.io/gorm"
)

// TransactionHistory ...
type TransactionHistory struct {
	gorm.Model
	UserID   uint
	ItemName string
	Amount   uint
}

// TableName ...
func (TransactionHistory) TableName() string {
	return "transaction_histories"
}

func (th TransactionHistory) makeTransactionHistory() transaction.History {
	transactionHistory := transaction.NewHistory(th.ItemName, th.Amount)
	transactionHistory.SetCreatedAt(th.CreatedAt)
	return *transactionHistory
}

// NewCreateTransactionHistory ...
func NewCreateTransactionHistory(p transaction.History, userID uint) TransactionHistory {
	return TransactionHistory{
		UserID:   userID,
		ItemName: string(p.ItemName()),
		Amount:   uint(p.Amount()),
	}
}

// MakeFetchHistoryDTO ...
func MakeFetchHistoryDTO(th TransactionHistory) *transaction.FetchHistoryDTO {
	return transaction.NewFetchHistoryDTO(th.makeTransactionHistory())
}

// TransactionHistories ...
type TransactionHistories []TransactionHistory

// MakeFetchHistoriesDTO ...
func MakeFetchHistoriesDTO(ths TransactionHistories) *transaction.FetchHistoriesDTO {
	transactionHistoryEntities := make(transaction.Histories, len(ths))
	for i, th := range ths {
		transactionHistoryEntities[i] = th.makeTransactionHistory()
	}
	return transaction.NewFetchHistoriesDTO(transactionHistoryEntities)
}
