package usecase

import (
	"context"
	"go-playground/m/v1/src/usecase/data/input"
	"go-playground/m/v1/src/usecase/data/output"
)

// IBalanceControlUsecase ...
type IBalanceControlUsecase interface {
	RetrieveRemainingBalanceByUserID(ctx context.Context, userID uint) (output.Balance, error)
	PutMoney(ctx context.Context, userID uint, topUpAmount input.PuttingMoney) error
	PayMoney(ctx context.Context, userID uint, payment input.Payment) error
}

// IDealUsecase ...
type IDealUsecase interface {
	RetrieveDealHistoriesByUserID(ctx context.Context, userID uint) (output.DealHistories, error)
}

// IGradeUsecase ...
type IGradeUsecase interface {
	RetrieveGrades(ctx context.Context) (*output.Grades, error)
}

// IUserManagementUsecase ...
type IUserManagementUsecase interface {
	CreateUser(ctx context.Context, input input.UserCreate, topUpAmount uint) error
	RetrieveUserByCondition(ctx context.Context, id uint) (output.User, error)
	RetrieveUsers(ctx context.Context) (output.Users, error)
}
