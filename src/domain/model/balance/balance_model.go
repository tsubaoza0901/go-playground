package balance

import (
	"errors"
	"go-playground/m/v1/src/domain/model/user"
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

type (
	// RemainingAmount 残金
	RemainingAmount uint
)

// InitialAmount ...
const initialAmount RemainingAmount = 0

// Entity 残高エンティティ
type Entity struct {
	userID          user.ID
	remainingAmount RemainingAmount
}

// UserID Getter
func (b *Entity) UserID() user.ID {
	return b.userID
}

// RemainingAmount Getter
func (b *Entity) RemainingAmount() RemainingAmount {
	return b.remainingAmount
}

// Exist ...
func (b *Entity) Exist(expect bool) error {
	if expect && b.UserID() == 0 {
		return errors.New("対象ユーザーの残高がありません。")
	}
	if !expect && b.UserID() != 0 {
		return errors.New("すでに残高登録済みです。")
	}
	return nil
}

// AddUp 残高足し算
func (b *Entity) AddUp(topUpAmount TopUpAmount) (*Entity, error) {
	if !topUpAmount.isMinimumAmountOrMore() {
		return nil, errors.New("チャージ額は500円以上必要です。")
	}
	calcuratedResult := uint(b.remainingAmount) + uint(topUpAmount)
	b.remainingAmount = RemainingAmount(calcuratedResult)
	return b, nil
}

// Subtract 残高引き算
func (b *Entity) Subtract(paymentAmount PaymentAmount) (*Entity, error) {
	if !b.hasEnoughRemainingAmount(paymentAmount) {
		return nil, errors.New("残高不足です。")
	}
	calcuratedResult := uint(b.remainingAmount) - uint(paymentAmount)
	b.remainingAmount = RemainingAmount(calcuratedResult)
	return b, nil
}

func (b *Entity) hasEnoughRemainingAmount(amount PaymentAmount) bool {
	return uint(b.remainingAmount) >= uint(amount)
}

// NewEntity ...
func NewEntity() *Entity {
	balance := new(Entity)
	balance.remainingAmount = initialAmount
	return balance
}

// MakeEntity ...
func MakeEntity(userID user.ID, remainingAmount RemainingAmount) *Entity {
	balance := new(Entity)
	balance.userID = userID
	balance.remainingAmount = remainingAmount
	return balance
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
