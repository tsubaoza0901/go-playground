package interactor

import (
	"context"
	"errors"
	"go-playground/m/v1/src/domain/model/balance"
	"go-playground/m/v1/src/domain/model/deal"
	"go-playground/m/v1/src/domain/model/user"
	"go-playground/m/v1/src/domain/repository"
	"go-playground/m/v1/src/usecase/data/input"
	"go-playground/m/v1/src/usecase/data/output"
)

// UserManagementUsecase ...
type UserManagementUsecase struct {
	repository.IUserManagementRepository
	repository.IBalanceRepository
	repository.IDealHistoryRepository
}

// NewUserManagementUsecase ...
func NewUserManagementUsecase(
	umr repository.IUserManagementRepository,
	br repository.IBalanceRepository,
	thr repository.IDealHistoryRepository,
) UserManagementUsecase {
	return UserManagementUsecase{umr, br, thr}
}

// CreateUser ...
func (u UserManagementUsecase) CreateUser(ctx context.Context, inputUserCreate input.UserCreate, inputTopUpAmount uint) error {
	generalUser, err := user.InitGeneral(inputUserCreate.FirstName, inputUserCreate.LastName, inputUserCreate.Age, inputUserCreate.EmailAddress)
	if err != nil {
		return err
	}

	if err := u.verifyIfNoUserHasSameEmailAddress(ctx, generalUser); err != nil {
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
	dealHistory := deal.NewHistory("", uint(inputTopUpAmount))
	if err := u.RegisterDealHistory(ctx, deal.NewRegisterHistoryDTO(*dealHistory, uint(generalUser.ID()))); err != nil {
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

func (u UserManagementUsecase) verifyIfNoUserHasSameEmailAddress(ctx context.Context, generalUser *user.General) error {
	count, err := u.CountTheNumberOfUsersByEmail(ctx, generalUser.EmailAddress())
	if err != nil {
		return err
	}
	if !user.IsSameUsersCountZero(count) {
		return errors.New("すでに同一ユーザーが存在します。")
	}
	return nil
}
