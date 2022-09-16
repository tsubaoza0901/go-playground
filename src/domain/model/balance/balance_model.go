package balance

import (
	"errors"
)

type (
	// TopUpAmount チャージ額
	TopUpAmount uint

	// PaymentAmount 支払額
	PaymentAmount uint
)

// MinimumTopUpAmount 最低チャージ金額
const MinimumTopUpAmount uint = 500

// IsMinimumAmountOrMore ...
func (a TopUpAmount) isMinimumAmountOrMore() bool {
	return uint(a) >= MinimumTopUpAmount
}

// RemainingAmount 残高（エンティティ）
type RemainingAmount uint

// AddUp 残高足し算
func (a RemainingAmount) AddUp(topUpAmount TopUpAmount) (*RemainingAmount, error) {
	if !topUpAmount.isMinimumAmountOrMore() {
		return nil, errors.New("チャージ額は500円以上必要です。")
	}
	calcuratedResult := uint(a) + uint(topUpAmount)
	a = RemainingAmount(calcuratedResult)
	return &a, nil
}

// Subtract 残高引き算
func (a RemainingAmount) Subtract(paymentAmount PaymentAmount) (*RemainingAmount, error) {
	if !a.hasEnoughRemainingAmount(paymentAmount) {
		return nil, errors.New("残高不足です。")
	}
	calcuratedResult := uint(a) - uint(paymentAmount)
	a = RemainingAmount(calcuratedResult)
	return &a, nil
}

func (a RemainingAmount) hasEnoughRemainingAmount(amount PaymentAmount) bool {
	return uint(a) >= uint(amount)
}

// -------------------------------

// // Entity 残高エンティティ
// type Entity struct {
// 	remainingAmount RemainingAmount
// }

// func newEntity() *Entity {
// 	return new(Entity)
// }

// // SetRemainingAmount Setter
// func (b *Entity) SetRemainingAmount(remainingAmount uint) {
// 	b.remainingAmount = RemainingAmount(remainingAmount)
// }

// // RemainingAmount Getter
// func (b *Entity) RemainingAmount() RemainingAmount {
// 	return b.remainingAmount
// }

// // Balance ...
// type Balance struct {
// 	Entity
// }

// // NewBalance ...
// func NewBalance() *Balance {
// 	entity := newEntity()
// 	return &Balance{*entity}
// }

// // AddUp 残高足し算
// func (b *Balance) AddUp(topUpAmount TopUpAmount) (*Balance, error) {
// 	if !topUpAmount.isMinimumAmountOrMore() {
// 		return nil, errors.New("please input minimum amount or more")
// 	}
// 	calcuratedRemainingAmount := uint(b.remainingAmount) + uint(topUpAmount)
// 	b.remainingAmount = RemainingAmount(calcuratedRemainingAmount)
// 	return b, nil
// }

// // Subtract 残高引き算
// func (b *Balance) Subtract(paymentAmount PaymentAmount) (*Balance, error) {
// 	if !b.hasEnoughAmount(paymentAmount) {
// 		return nil, errors.New("残高不足")
// 	}
// 	calcuratedRemainingAmount := uint(b.remainingAmount) - uint(paymentAmount)
// 	b.remainingAmount = RemainingAmount(calcuratedRemainingAmount)
// 	return b, nil
// }

// func (b *Balance) hasEnoughAmount(paymentAmount PaymentAmount) bool {
// 	return uint(b.remainingAmount) >= uint(paymentAmount)
// }
