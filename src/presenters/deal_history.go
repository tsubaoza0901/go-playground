package presenters

import (
	"go-playground/m/v1/presenters/response"
	"go-playground/m/v1/usecase/data/output"
)

// DealHistory ...
type DealHistory struct {
	response.AppResponse
}

// NewDealHistory ...
func NewDealHistory() *DealHistory {
	return &DealHistory{}
}

// DealHistoryList ...
func (p *DealHistory) DealHistoryList(dealHistoryList []*output.DealHistory) {
	data := make([]*response.DealHistory, len(dealHistoryList))
	for i, v := range dealHistoryList {
		data[i] = &response.DealHistory{
			Date:     v.CreatedAt.ToFormatedString(),
			ItemName: v.ItemName,
			Amount:   v.Amount,
		}
	}
	p.AppResponse = response.AppResponse{
		Data: data,
	}
}

// AppError ...
func (p *DealHistory) AppError(errorCode int) {
	p.AppResponse = response.AppResponse{
		Error: (*response.ErrorCode)(&errorCode),
	}
}
