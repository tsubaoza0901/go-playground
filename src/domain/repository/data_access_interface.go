package repository

import (
	"context"
	"go-playground/m/v1/src/domain/model/balance"
	"go-playground/m/v1/src/domain/model/deal"
	"go-playground/m/v1/src/domain/model/grade"
	"go-playground/m/v1/src/domain/model/user"
)

// IBalanceRepository ...
type IBalanceRepository interface {
	FetchBalanceByUserID(ctx context.Context, userID uint) (*balance.FetchAmountDTO, error)
	CreateBalance(ctx context.Context, userID uint, bl balance.CreateBalanceDTO) error
	UpdateBalance(ctx context.Context, userID uint, bl balance.UpdateBalanceDTO) error
}

// IGradeRepository ...
type IGradeRepository interface {
	FetchAllGrades(context.Context) (*grade.FetchAllDTO, error)
}

// IDealHistoryRepository ...
type IDealHistoryRepository interface {
	RegisterDealHistory(context.Context, deal.RegisterHistoryDTO) error
	FetchDealHistoriesByUserID(ctx context.Context, userID uint) (*deal.FetchHistoriesDTO, error)
}

// IUserManagementRepository ...
type IUserManagementRepository interface {
	RegisterUser(context.Context, user.RegistrationDTO) (*user.FetchDTO, error)
	FetchUser(ctx context.Context, id uint) (*user.FetchDTO, error)
	FetchAllUsers(ctx context.Context) (*user.FetchAllDTO, error)
}
