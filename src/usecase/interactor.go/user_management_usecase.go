package interactor

import (
	"context"
	"go-playground/m/v1/src/domain/model/balance"
	"go-playground/m/v1/src/domain/model/transaction"
	"go-playground/m/v1/src/domain/model/user"
	"go-playground/m/v1/src/domain/repository"
	"go-playground/m/v1/src/usecase/data/input"
	"go-playground/m/v1/src/usecase/data/output"
)

// UserManagementUsecase ...
type UserManagementUsecase struct {
	repository.IUserManagementRepository
	repository.IBalanceRepository
	repository.ITransactionHistoryRepository
}

// NewUserManagementUsecase ...
func NewUserManagementUsecase(
	umr repository.IUserManagementRepository,
	br repository.IBalanceRepository,
	thr repository.ITransactionHistoryRepository,
) UserManagementUsecase {
	return UserManagementUsecase{umr, br, thr}
}

// CreateUser ...
func (u UserManagementUsecase) CreateUser(ctx context.Context, inputUserCreate input.UserCreate, inputTopUpAmount uint) error {
	generalUser, err := user.InitGeneral(inputUserCreate.FirstName, inputUserCreate.LastName, inputUserCreate.Age)
	if err != nil {
		return err
	}
	userFetchDTO, err := u.RegisterUser(ctx, user.SetFieldToRegistrationDTO(*generalUser))
	if err != nil {
		return err
	}

	generalUser = &userFetchDTO.General

	// 残高登録
	remainingAmount := balance.RemainingAmount(0)
	calculatedBalance, err := remainingAmount.AddUp(balance.TopUpAmount(inputTopUpAmount))
	if err != nil {
		return err
	}
	if err := u.CreateBalance(ctx, uint(generalUser.ID()), balance.NewCreateBalanceDTO(*calculatedBalance)); err != nil {
		return err
	}

	// 取引履歴登録
	transactionHistory := transaction.NewHistory("", uint(inputTopUpAmount))
	if err := u.RegisterTransactionHistory(ctx, transaction.NewRegisterHistoryDTO(*transactionHistory, uint(generalUser.ID()))); err != nil {
		return err
	}
	return nil
}

// RetrieveUserByCondition ...
func (u UserManagementUsecase) RetrieveUserByCondition(ctx context.Context, id uint) (output.User, error) {
	userFetchDTO, err := u.FetchUser(ctx, id)
	if err != nil {
		return output.User{}, err
	}
	return output.MakeUser(userFetchDTO.General), nil
}

// RetrieveUsers ...
func (u UserManagementUsecase) RetrieveUsers(ctx context.Context) (output.Users, error) {
	userFetchAllDTO, err := u.FetchAllUsers(ctx)
	if err != nil {
		return nil, err
	}

	return output.MakeUsers(userFetchAllDTO.Generals), nil
}
