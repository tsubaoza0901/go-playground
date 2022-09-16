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

	history.setItemName(itemName)
	history.setAmount(amount)

	return history
}

// CreatedAt Getter
func (p *Entity) CreatedAt() CreatedAt {
	return p.createdAt
}

func (p *Entity) setCreatedAt(createdAt CreatedAt) {
	p.createdAt = createdAt
}

// ItemName Getter
func (p *Entity) ItemName() ItemName {
	return p.itemName
}

func (p *Entity) setItemName(itemName ItemName) {
	p.itemName = itemName
}

// Amount Getter
func (p *Entity) Amount() Amount {
	return p.amount
}

func (p *Entity) setAmount(amount Amount) {
	p.amount = amount
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
	history.setCreatedAt(createdAt)
	return history
}

// Histories ...
type Histories []History
