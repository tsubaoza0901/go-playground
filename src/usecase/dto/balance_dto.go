package dto

import (
	"go-playground/m/v1/domain/model/balance"
	"go-playground/m/v1/domain/model/user"
)

// RegisterBalance DTO
type RegisterBalance struct {
	UserID          uint
	RemainingAmount uint
}

// NewRegisterBalance ...
func NewRegisterBalance(userID user.ID, remainingAmount uint) RegisterBalance {
	return RegisterBalance{
		UserID:          uint(userID),
		RemainingAmount: remainingAmount,
	}
}

// UpdateBalance DTO
type UpdateBalance struct {
	UserID          uint
	RemainingAmount uint
}

// NewUpdateBalance ...
func NewUpdateBalance(userID user.ID, remainingAmount uint) UpdateBalance {
	return UpdateBalance{
		UserID:          uint(userID),
		RemainingAmount: remainingAmount,
	}
}

// FetchBlanceResult DTO
type FetchBlanceResult struct {
	UserID          uint
	RemainingAmount uint
}

// NewFetchBlanceResult ...
func NewFetchBlanceResult(userID uint, remainingAmount uint) *FetchBlanceResult {
	return &FetchBlanceResult{
		UserID:          userID,
		RemainingAmount: remainingAmount,
	}
}

// ToBalanceModel ...
func (b FetchBlanceResult) ToBalanceModel() balance.Entity {
	balanceEntity := balance.NewEntity(user.ID(b.UserID), balance.RemainingAmount(b.RemainingAmount))
	return *balanceEntity
}
