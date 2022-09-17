package interactor

import (
	"context"
	"go-playground/m/v1/src/domain/model/balance"
	"go-playground/m/v1/src/domain/model/deal"
	"go-playground/m/v1/src/domain/model/user"
	"go-playground/m/v1/src/usecase/data/input"
	"go-playground/m/v1/src/usecase/data/output"
	"go-playground/m/v1/src/usecase/repository"
	"go-playground/m/v1/src/usecase/repository/dto"
)

// UserManagementUsecase ...
type UserManagementUsecase struct {
	repository.IUserManagementRepository
	repository.IBalanceRepository
	repository.IDealHistoryRepository
	repository.ITransactionManagementRepository
}

// NewUserManagementUsecase ...
func NewUserManagementUsecase(
	umr repository.IUserManagementRepository,
	br repository.IBalanceRepository,
	thr repository.IDealHistoryRepository,
	tmr repository.ITransactionManagementRepository,
) UserManagementUsecase {
	return UserManagementUsecase{umr, br, thr, tmr}
}

// CreateUser ...
func (u UserManagementUsecase) CreateUser(ctx context.Context, inputUserCreate input.UserCreate, inputTopUpAmount uint) error {
	initialUser, err := user.NewGeneral(inputUserCreate.FirstName, inputUserCreate.LastName, inputUserCreate.Age, inputUserCreate.EmailAddress)
	if err != nil {
		return err
	}

	// ユーザー重複確認
	if err := u.verifyThatNoUserHasSameEmail(ctx, initialUser.EmailAddress()); err != nil {
		return err
	}

	// 同一トランザクション内処理開始
	if err := u.Transaction(ctx, func(ctx context.Context) (err error) {

		// ユーザー登録
		generalUser, err := u.registerUser(ctx, initialUser)
		if err != nil {
			return err
		}

		// チャージ結果計算
		topUpAmount := balance.TopUpAmount(inputTopUpAmount)

		initialBalance := balance.NewEntity()
		calculatedBalance, err := initialBalance.AddUp(topUpAmount)
		if err != nil {
			return err
		}

		// 残高登録
		if err = u.RegisterBalance(ctx, dto.NewRegisterBalance(generalUser.ID(), calculatedBalance.RemainingAmount())); err != nil {
			return err
		}

		// 取引履歴登録
		dealHistory := deal.NewTopUpHistory(topUpAmount)
		if err = u.RegisterDealHistory(ctx, dto.NewRegisterDealHistory(generalUser.ID(), dealHistory.ItemName(), dealHistory.Amount())); err != nil {
			return err
		}
		return
	}); err != nil {
		return err
	}

	return nil
}

// RetrieveUserByCondition ...
func (u UserManagementUsecase) RetrieveUserByCondition(ctx context.Context, id uint) (output.User, error) {
	targetUser, err := u.fetchUserbyID(ctx, id)
	if err != nil {
		return output.User{}, err
	}
	if err := targetUser.Exist(true); err != nil {
		return output.User{}, err
	}
	return output.MakeUser(*targetUser), nil
}

// RetrieveUsers ...
func (u UserManagementUsecase) RetrieveUsers(ctx context.Context) (output.Users, error) {
	targetUserList, err := u.fetchUserList(ctx)
	if err != nil {
		return nil, err
	}
	return output.MakeUsers(targetUserList), nil
}

func (u UserManagementUsecase) verifyThatNoUserHasSameEmail(ctx context.Context, email string) error {
	targetUser, err := u.fetchUserbyEmail(ctx, email)
	if err != nil {
		return err
	}
	if err := targetUser.Exist(false); err != nil {
		return err
	}
	return nil
}

func (u UserManagementUsecase) registerUser(ctx context.Context, generalUser *user.General) (*user.General, error) {
	fetchResult, err := u.RegisterUser(ctx, dto.NewRegisterUser(*generalUser))
	if err != nil {
		return nil, err
	}

	generalUser, err = fetchResult.ToGeneralUserModel()
	if err != nil {
		return nil, err
	}
	return generalUser, nil
}

func (u UserManagementUsecase) fetchUserbyID(ctx context.Context, id uint) (*user.General, error) {
	fetchResult, err := u.FetchUserByID(ctx, id)
	if err != nil {
		return nil, err
	}
	generalUser, err := fetchResult.ToGeneralUserModel()
	if err != nil {
		return nil, err
	}
	return generalUser, nil
}

func (u UserManagementUsecase) fetchUserbyEmail(ctx context.Context, email string) (*user.General, error) {
	fetchResult, err := u.FetchUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	generalUser, err := fetchResult.ToGeneralUserModel()
	if err != nil {
		return nil, err
	}
	return generalUser, nil
}

func (u UserManagementUsecase) fetchUserList(ctx context.Context) (user.Generals, error) {
	fetchResult, err := u.FetchUserList(ctx)
	if err != nil {
		return nil, err
	}

	generalUserList, err := fetchResult.ToGeneralUsersModel()
	if err != nil {
		return nil, err
	}
	return generalUserList, nil
}
