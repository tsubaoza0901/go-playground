package model

import (
	"database/sql"
	"go-playground/m/v1/src/domain/model/balance"
	"time"
)

// Balance ...
type Balance struct {
	UserID    uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime `gorm:"index"`
	Amount    uint
}

// TableName ...
func (Balance) TableName() string {
	return "balance"
}

// InitBalance ...
func InitBalance(userID uint, remainingAmount balance.RemainingAmount) *Balance {
	return &Balance{
		UserID: userID,
		Amount: uint(remainingAmount),
	}
}

// MakeBalanceFetchAmountDTO ...
func MakeBalanceFetchAmountDTO(b Balance) *balance.FetchAmountDTO {
	remainingAmount := balance.RemainingAmount(b.Amount)
	fetchAmountDTO := balance.NewFetchAmountDTO(remainingAmount)
	return &fetchAmountDTO
}
