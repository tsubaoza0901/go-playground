package interactor

import (
	"context"
	"go-playground/m/v1/domain/model/deal"
	"go-playground/m/v1/usecase/data/output"
	"go-playground/m/v1/usecase/interactor/port"
	"go-playground/m/v1/usecase/rule"
)

// DealHistoryUsecase ...
type DealHistoryUsecase struct {
	port.IDealHistoryRepository
	dealHistoryOutputPort port.DealHistoryOutput
}

// NewDealHistoryUsecase ...
func NewDealHistoryUsecase(dhr port.IDealHistoryRepository, dhop port.DealHistoryOutput) *DealHistoryUsecase {
	return &DealHistoryUsecase{dhr, dhop}
}

// RetrieveDealHistoriesByUserID ...
func (u *DealHistoryUsecase) RetrieveDealHistoriesByUserID(ctx context.Context, userID uint) {
	tragetDealHistories, err := u.fetchDealHistoriesByUserID(ctx, userID)
	if err != nil {
		u.dealHistoryOutputPort.AppError(rule.InternalServerError)
		return
	}
	outputList := make([]*output.DealHistory, len(tragetDealHistories))
	for i, v := range tragetDealHistories {
		outputList[i] = &output.DealHistory{
			CreatedAt: output.CreatedAt(v.CreatedAt()),
			ItemName:  v.ItemName(),
			Amount:    v.Amount(),
		}
	}
	u.dealHistoryOutputPort.DealHistoryList(outputList)
}

func (u *DealHistoryUsecase) fetchDealHistoriesByUserID(ctx context.Context, userID uint) (deal.Histories, error) {
	fetchResult, err := u.FetchDealHistoriesByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return fetchResult.ToDealHistoriesModel(), nil
}
