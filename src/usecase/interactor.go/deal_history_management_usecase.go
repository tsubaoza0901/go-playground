package interactor

import (
	"context"
	"go-playground/m/v1/src/domain/model/deal"
	"go-playground/m/v1/src/usecase/data/output"
	"go-playground/m/v1/src/usecase/repository"
)

// DealUsecase ...
type DealUsecase struct {
	repository.IDealHistoryRepository
}

// NewDealUsecase ...
func NewDealUsecase(pr repository.IDealHistoryRepository) DealUsecase {
	return DealUsecase{pr}
}

// RetrieveDealHistoriesByUserID ...
func (u DealUsecase) RetrieveDealHistoriesByUserID(ctx context.Context, userID uint) (output.DealHistories, error) {
	dealHistories, err := u.fetchDealHistoriesByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return output.MakeDealHistories(dealHistories), nil
}

func (u DealUsecase) fetchDealHistoriesByUserID(ctx context.Context, userID uint) (deal.Histories, error) {
	fetchResult, err := u.FetchDealHistoriesByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return fetchResult.ToDealHistoriesModel(), nil
}
