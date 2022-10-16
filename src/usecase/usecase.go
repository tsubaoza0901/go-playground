//go:generate mockgen -source=$GOFILE -package=mock -destination=$GOPATH/src/mock/$GOFILE

package usecase

import (
	"context"
	"go-playground/m/v1/usecase/data/input"
)

// IBalanceControlUsecase ...
type IBalanceControlUsecase interface {
	RetrieveRemainingBalanceByUserID(ctx context.Context, userID uint)
	PutMoney(ctx context.Context, userID uint, topUpAmount input.PuttingMoney)
	PayMoney(ctx context.Context, userID uint, payment input.Payment)
}

// IDealUsecase ...
type IDealUsecase interface {
	RetrieveDealHistoriesByUserID(ctx context.Context, userID uint)
}

// IGradeUsecase ...
type IGradeUsecase interface {
	RetrieveGrades(ctx context.Context)
}

// IUserManagementUsecase ...
type IUserManagementUsecase interface {
	CreateUser(ctx context.Context, input input.UserCreate, topUpAmount uint)
	EditUser(ctx context.Context, input input.UserUpdate)
	RetrieveUserByCondition(ctx context.Context, id uint)
	RetrieveUsers(ctx context.Context)
}
