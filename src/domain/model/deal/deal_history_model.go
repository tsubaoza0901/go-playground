package deal

import (
	"go-playground/m/v1/domain/model/balance"
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
func (d *Entity) CreatedAt() time.Time {
	return time.Time(d.createdAt)
}

// ItemName Getter
func (d *Entity) ItemName() string {
	return string(d.itemName)
}

// Amount Getter
func (d *Entity) Amount() uint {
	return uint(d.amount)
}

// History ...
type History struct {
	Entity
}

// InitPaymentHistory 初期化関数。原則、新規レコード登録用。
func InitPaymentHistory(itemName string, amount balance.PaymentAmount) *History {
	return newHistory(ItemName(itemName), Amount(amount))
}

// InitTopUpHistory 初期化関数（一部、初期値の設定あり）。原則、新規レコード登録用。
func InitTopUpHistory(amount balance.TopUpAmount) *History {
	return newHistory("チャージ", Amount(amount))
}

// NewHistory 新規インスタンス生成関数（初期値の設定なし）
func NewHistory(createdAt CreatedAt, itemName ItemName, amount Amount) *History {
	history := newHistory(itemName, amount)
	history.createdAt = createdAt
	return history
}

// Histories ...
type Histories []History
