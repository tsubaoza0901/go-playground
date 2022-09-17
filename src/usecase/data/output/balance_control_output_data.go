package output

import (
	"fmt"
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
func MakeBalance(remainingAmount uint) Balance {
	return Balance{
		Amount: Amount(remainingAmount),
	}
}
