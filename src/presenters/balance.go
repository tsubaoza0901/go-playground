package presenters

import (
	"go-playground/m/v1/presenters/response"
	"go-playground/m/v1/usecase/data/output"
)

// Balance ...
type Balance struct {
	response.AppResponse
}

// NewBalance ...
func NewBalance() *Balance {
	return &Balance{}
}

// Balance ...
func (p *Balance) Balance(balance *output.Balance) {
	data := response.Balance{
		Amount: balance.Amount.ToJPYString(),
	}
	p.AppResponse = response.AppResponse{
		Data: data,
	}
}

// AppError ...
func (p *Balance) AppError(errorCode int) {
	p.AppResponse = response.AppResponse{
		Error: (*response.ErrorCode)(&errorCode),
	}
}
