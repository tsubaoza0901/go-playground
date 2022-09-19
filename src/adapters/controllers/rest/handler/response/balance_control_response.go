package response

import "go-playground/m/v1/usecase/data/output"

// Balance ...
type Balance struct {
	Amount string `json:"amount"`
}

// NewBalance ...
func NewBalance(b output.Balance) Balance {
	return Balance{
		Amount: b.Amount.ToJPYString(),
	}
}
