package repository

import (
	"context"
	"go-playground/m/v1/src/usecase/repository/dto"
)

type contextKey string

// TransactionKey ...
const TransactionKey contextKey = "transaction"

// IBalanceRepository ...
type IBalanceRepository interface {
	RegisterBalance(ctx context.Context, userID uint, createBalanceDTO dto.RegisterBalance) error
	UpdateBalance(ctx context.Context, userID uint, updateBalanceDTO dto.UpdateBalance) error
	FetchBalanceByUserID(ctx context.Context, userID uint) (*dto.FetchBlanceResult, error)
}

// IGradeRepository ...
type IGradeRepository interface {
	FetchGradeList(context.Context) (*dto.FetchGradeListResult, error)
}

// IDealHistoryRepository ...
type IDealHistoryRepository interface {
	RegisterDealHistory(context.Context, dto.RegisterDealHistory) error
	FetchDealHistoriesByUserID(ctx context.Context, userID uint) (*dto.FetchDealHistoryListResult, error)
}

// IUserManagementRepository ...
type IUserManagementRepository interface {
	RegisterUser(context.Context, dto.RegisterUser) (*dto.FetchUserResult, error)
	FetchUserByID(ctx context.Context, id uint) (*dto.FetchUserResult, error)
	FetchUserList(ctx context.Context) (*dto.FetchUserListResult, error)
	CountTheNumberOfUsersByEmail(ctx context.Context, email string) (count uint, err error)
}

// ITransactionManagementRepository ...
type ITransactionManagementRepository interface {
	Transaction(ctx context.Context, fc func(ctx context.Context) error) error
	// NewContextWithTx(ctx context.Context) context.Context
	// CommitByContext(ctx context.Context) error
	// RollbackByContext(ctx context.Context) error
}
