//go:generate mockgen -source=$GOFILE -package=mock -destination=$GOPATH/src/mock/$GOFILE

package repository

import (
	"context"
	"go-playground/m/v1/usecase/repository/dto"
)

// IBalanceRepository ...
type IBalanceRepository interface {
	RegisterBalance(ctx context.Context, createBalanceDTO dto.RegisterBalance) error
	UpdateBalance(ctx context.Context, updateBalanceDTO dto.UpdateBalance) error
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
	FetchUserByEmail(ctx context.Context, email string) (*dto.FetchUserResult, error)
	FetchUserList(ctx context.Context) (*dto.FetchUserListResult, error)
}

// ITransactionManagementRepository ...
type ITransactionManagementRepository interface {
	Transaction(ctx context.Context, fc func(ctx context.Context) error) error
}
