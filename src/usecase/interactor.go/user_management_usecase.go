package interactor

import (
	"context"
	"go-playground/m/v1/domain/model/balance"
	"go-playground/m/v1/domain/model/deal"
	"go-playground/m/v1/domain/model/grade"
	"go-playground/m/v1/domain/model/user"
	"go-playground/m/v1/usecase/data/input"
	"go-playground/m/v1/usecase/data/output"
	"go-playground/m/v1/usecase/repository"
	"go-playground/m/v1/usecase/repository/dto"
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
	initialUser, err := user.InitGeneral(inputUserCreate.FirstName, inputUserCreate.LastName, inputUserCreate.Age, inputUserCreate.EmailAddress)
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

		initialBalance := balance.InitEntity()
		calculatedBalance, err := initialBalance.AddUp(topUpAmount)
		if err != nil {
			return err
		}

		// 残高登録
		if err = u.RegisterBalance(ctx, dto.NewRegisterBalance(generalUser.ID(), calculatedBalance.RemainingAmount())); err != nil {
			return err
		}

		// 取引履歴登録
		dealHistory := deal.InitTopUpHistory(topUpAmount)
		if err = u.RegisterDealHistory(ctx, dto.NewRegisterDealHistory(generalUser.ID(), dealHistory.ItemName(), dealHistory.Amount())); err != nil {
			return err
		}
		return
	}); err != nil {
		return err
	}

	return nil
}

// EditUser ...
func (u UserManagementUsecase) EditUser(ctx context.Context, inputUserUpdate input.UserUpdate) error {
	updateUser, err := user.UpdateGeneral(inputUserUpdate.ID, inputUserUpdate.LastName, inputUserUpdate.EmailAddress, grade.ID(inputUserUpdate.GradeID))
	if err != nil {
		return err
	}

	// ユーザー存在確認
	if err := u.verifyThatUserExist(ctx, updateUser.ID()); err != nil {
		return err
	}

	// ユーザー重複確認
	if err := u.verifyThatNoUserHasSameEmail(ctx, updateUser.EmailAddress()); err != nil {
		return err
	}

	// ユーザー情報更新
	_, err = u.editUser(ctx, updateUser)
	if err != nil {
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

func (u UserManagementUsecase) verifyThatUserExist(ctx context.Context, id user.ID) error {
	targetUser, err := u.fetchUserbyID(ctx, uint(id))
	if err != nil {
		return err
	}
	if err := targetUser.Exist(true); err != nil {
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

func (u UserManagementUsecase) editUser(ctx context.Context, generalUser *user.General) (*user.General, error) {
	fetchResult, err := u.UpdateUser(ctx, uint(generalUser.ID()), dto.NewUpdateUser(*generalUser))
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
