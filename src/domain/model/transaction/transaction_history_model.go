package transaction

import (
	"log"
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

// CreatedAt Getter
func (p *Entity) CreatedAt() CreatedAt {
	return p.createdAt
}

// SetCreatedAt Setter
func (p *Entity) SetCreatedAt(createdAt time.Time) {
	p.createdAt = CreatedAt(createdAt)
}

// ItemName Getter
func (p *Entity) ItemName() ItemName {
	return p.itemName
}

func (p *Entity) setItemName(itemName string) {
	p.itemName = ItemName(itemName)
}

// Amount Getter
func (p *Entity) Amount() Amount {
	return p.amount
}

func (p *Entity) setAmount(amount uint) {
	p.amount = Amount(amount)
}

// History ...
type History struct {
	Entity
}

// NewHistory ...
func NewHistory(itemName string, amount uint) *History {
	history := new(History)
	log.Println(itemName, amount)

	if itemName == "" && amount >= 0 {
		history.setItemName("チャージ")
	} else {
		history.setItemName(itemName)
	}
	history.setAmount(amount)
	return history
}

// Histories ...
type Histories []History
