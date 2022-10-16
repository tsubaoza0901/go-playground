package persistence

import (
	"context"
	"go-playground/m/v1/usecase/dto"
)

// TransactionManagementDataAccess ...
type TransactionManagementDataAccess interface {
	Transaction(ctx context.Context, fc func(ctx context.Context) error) error
}

// BalanceDataAccess ...
type BalanceDataAccess interface {
	RegisterBalance(ctx context.Context, createBalanceDTO dto.RegisterBalance) error
	UpdateBalance(ctx context.Context, updateBalanceDTO dto.UpdateBalance) error
	FetchBalanceByUserID(ctx context.Context, userID uint) (*dto.FetchBlanceResult, error)
}

// DealHistoryDataAccess ...
type DealHistoryDataAccess interface {
	RegisterDealHistory(context.Context, dto.RegisterDealHistory) error
	FetchDealHistoriesByUserID(ctx context.Context, userID uint) (*dto.FetchDealHistoryListResult, error)
}

// GradeDataAccess ...
type GradeDataAccess interface {
	FetchGradeList(context.Context) (*dto.FetchGradeListResult, error)
}

// UserDataAccess ...
type UserDataAccess interface {
	RegisterUser(context.Context, dto.RegisterUser) (*dto.FetchUserResult, error)
	UpdateUser(ctx context.Context, id uint, dtoUpdateUser dto.UpdateUser) (*dto.FetchUserResult, error)
	FetchUserByID(ctx context.Context, id uint) (*dto.FetchUserResult, error)
	FetchUserByEmail(ctx context.Context, email string) (*dto.FetchUserResult, error)
	FetchUserList(ctx context.Context) (*dto.FetchUserListResult, error)
}
