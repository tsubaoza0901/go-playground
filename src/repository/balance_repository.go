package repository

import (
	"context"
	"go-playground/m/v1/repository/persistence"
	"go-playground/m/v1/usecase/dto"
)

// Balance ...
type Balance struct {
	balanceDataAccess persistence.BalanceDataAccess
}

// NewBalance ...
func NewBalance(bda persistence.BalanceDataAccess) Balance {
	return Balance{bda}
}

// RegisterBalance ...
func (r Balance) RegisterBalance(ctx context.Context, dto dto.RegisterBalance) error {
	return r.balanceDataAccess.RegisterBalance(ctx, dto)
}

// UpdateBalance ...
func (r Balance) UpdateBalance(ctx context.Context, dto dto.UpdateBalance) error {
	return r.balanceDataAccess.UpdateBalance(ctx, dto)
}

// FetchBalanceByUserID ...
func (r Balance) FetchBalanceByUserID(ctx context.Context, userID uint) (*dto.FetchBlanceResult, error) {
	return r.balanceDataAccess.FetchBalanceByUserID(ctx, userID)
}
