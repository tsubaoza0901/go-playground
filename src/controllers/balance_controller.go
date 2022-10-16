package controllers

import (
	"context"
	"go-playground/m/v1/controllers/request"
	"go-playground/m/v1/usecase"
)

// BalanceControl ...
type BalanceControl struct {
	balanceControlUsecase usecase.IBalanceControlUsecase
}

// NewBalanceControl ...
func NewBalanceControl(bcu usecase.IBalanceControlUsecase) *BalanceControl {
	return &BalanceControl{bcu}
}

// Pay ...
func (c *BalanceControl) Pay(ctx context.Context, req *request.Payment) {
	c.balanceControlUsecase.PayMoney(ctx, req.UserID, req.ConvertToPaymentInput())
}

// TopUp ...
func (c *BalanceControl) TopUp(ctx context.Context, req *request.PuttingMoney) {
	c.balanceControlUsecase.PutMoney(ctx, req.UserID, req.ConvertToPuttingMoneyInput())
}

// GetRemainingBalance ...
func (c *BalanceControl) GetRemainingBalance(ctx context.Context, req *request.RetrieveRemainingBalance) {
	c.balanceControlUsecase.RetrieveRemainingBalanceByUserID(ctx, req.UserID)
}
