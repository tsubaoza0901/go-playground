package deal

import (
	"time"
)

type (
	// CreatedAt 取引日時
	CreatedAt time.Time
	// ItemName 取引項目
	ItemName string
	// Amount 取引金額
	Amount uint
)

// Entity ...
type Entity struct {
	// userID user.ID

	createdAt CreatedAt
	itemName  ItemName
	amount    Amount
}

func newHistory(itemName ItemName, amount Amount) *History {
	history := new(History)

	history.itemName = itemName
	history.amount = amount

	return history
}

// CreatedAt Getter
func (d *Entity) CreatedAt() CreatedAt {
	return d.createdAt
}

// ItemName Getter
func (d *Entity) ItemName() ItemName {
	return d.itemName
}

// Amount Getter
func (d *Entity) Amount() Amount {
	return d.amount
}

// History ...
type History struct {
	Entity
}

// NewPaymentHistory ...
func NewPaymentHistory(itemName string, amount uint) *History {
	return newHistory(ItemName(itemName), Amount(amount))
}

// NewTopUpHistory ...
func NewTopUpHistory(amount uint) *History {
	return newHistory("チャージ", Amount(amount))
}

// MakeHistory ...
func MakeHistory(createdAt CreatedAt, itemName ItemName, amount Amount) *History {
	history := newHistory(itemName, amount)
	history.createdAt = createdAt
	return history
}

// Histories ...
type Histories []History
