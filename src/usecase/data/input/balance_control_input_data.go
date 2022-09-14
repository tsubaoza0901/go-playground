package input

// Payment ...
type Payment struct {
	ItemName string
	Amount   uint
}

// NewPayment ...
func NewPayment() Payment {
	return Payment{}
}

// PuttingMoney ...
type PuttingMoney struct {
	Amount uint
}

// NewPuttingMoney ...
func NewPuttingMoney() PuttingMoney {
	return PuttingMoney{}
}
