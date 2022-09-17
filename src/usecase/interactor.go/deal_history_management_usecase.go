package interactor

import (
	"context"
	"go-playground/m/v1/src/domain/model/deal"
	"go-playground/m/v1/src/usecase/data/output"
	"go-playground/m/v1/src/usecase/repository"
)

// DealHistoryUsecase ...
type DealHistoryUsecase struct {
	repository.IDealHistoryRepository
}

// NewDealHistoryUsecase ...
func NewDealHistoryUsecase(pr repository.IDealHistoryRepository) DealHistoryUsecase {
	return DealHistoryUsecase{pr}
}

// RetrieveDealHistoriesByUserID ...
func (u DealHistoryUsecase) RetrieveDealHistoriesByUserID(ctx context.Context, userID uint) (output.DealHistories, error) {
	tragetDealHistories, err := u.fetchDealHistoriesByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return output.MakeDealHistories(tragetDealHistories), nil
}

func (u DealHistoryUsecase) fetchDealHistoriesByUserID(ctx context.Context, userID uint) (deal.Histories, error) {
	fetchResult, err := u.FetchDealHistoriesByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return fetchResult.ToDealHistoriesModel(), nil
}
