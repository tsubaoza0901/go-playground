package interactor

import (
	"context"
	"go-playground/m/v1/src/domain/repository"
	"go-playground/m/v1/src/usecase/data/output"
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
	fetchHistoriesDTO, err := u.FetchDealHistoriesByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return output.MakeDealHistories(fetchHistoriesDTO.Histories), nil
}
