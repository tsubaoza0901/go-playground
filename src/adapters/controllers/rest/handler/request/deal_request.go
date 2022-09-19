package request

import (
	"go-playground/m/v1/usecase/data/input"
)

// Payment ...
type Payment struct {
	UserID   uint   `param:"userId" validate:"required"`
	ItemName string `json:"itemName" validate:"required"`
	Price    uint   `json:"price" validate:"required"`
}

// ConvertToPaymentInput ...
func (p Payment) ConvertToPaymentInput() input.Payment {
	payment := input.NewPayment()
	payment.ItemName = p.ItemName
	payment.Amount = p.Price
	return payment
}

// PuttingMoney ...
type PuttingMoney struct {
	UserID uint `param:"userId" validate:"required"`
	Amount uint `json:"amount" validate:"required"`
}

// ConvertToPuttingMoneyInput ...
func (p PuttingMoney) ConvertToPuttingMoneyInput() input.PuttingMoney {
	puttingMoney := input.NewPuttingMoney()
	puttingMoney.Amount = p.Amount
	return puttingMoney
}

// RetrieveDealHistories ...
type RetrieveDealHistories struct {
	UserID uint `param:"userId" validate:"required"`
}
