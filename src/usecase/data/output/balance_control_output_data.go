package output

import (
	"fmt"
	"go-playground/m/v1/src/domain/model/balance"
)

// Amount ...
type Amount uint

// ToJPYString ..
func (a Amount) ToJPYString() string {
	return fmt.Sprintf("%v円", a)
}

// Balance ...
type Balance struct {
	Amount Amount
}

// MakeBalance ...
func MakeBalance(balance balance.RemainingAmount) Balance {
	return Balance{
		Amount: Amount(balance),
	}
}
