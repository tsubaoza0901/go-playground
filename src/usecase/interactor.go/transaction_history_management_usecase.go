package interactor

import (
	"context"
	"go-playground/m/v1/src/domain/repository"
	"go-playground/m/v1/src/usecase/data/output"
)

// TransactionUsecase ...
type TransactionUsecase struct {
	repository.ITransactionHistoryRepository
}

// NewTransactionUsecase ...
func NewTransactionUsecase(pr repository.ITransactionHistoryRepository) TransactionUsecase {
	return TransactionUsecase{pr}
}

// RetrieveTransactionHistoriesByUserID ...
func (u TransactionUsecase) RetrieveTransactionHistoriesByUserID(ctx context.Context, userID uint) (output.TransactionHistories, error) {
	fetchHistoriesDTO, err := u.FetchTransactionHistoriesByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return output.MakeTransactionHistories(fetchHistoriesDTO.Histories), nil
}
