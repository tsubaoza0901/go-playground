package model

import (
	"database/sql"
	"go-playground/m/v1/src/usecase/repository/dto"
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

// ConvertToBalance ...
func ConvertToBalance(userID uint, remainingAmount uint) Balance {
	return Balance{
		UserID: userID,
		Amount: remainingAmount,
	}
}

// MakeFetchBlanceResultDTO ...
func MakeFetchBlanceResultDTO(b Balance) *dto.FetchBlanceResult {
	return dto.NewFetchBlanceResult(b.UserID, b.Amount)
}
