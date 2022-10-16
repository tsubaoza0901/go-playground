package interactor

import (
	"context"
	"go-playground/m/v1/domain/model/balance"
	"go-playground/m/v1/domain/model/deal"
	"go-playground/m/v1/domain/model/grade"
	"go-playground/m/v1/domain/model/user"
	"go-playground/m/v1/usecase/data/input"
	"go-playground/m/v1/usecase/data/output"
	"go-playground/m/v1/usecase/dto"
	"go-playground/m/v1/usecase/interactor/port"
	"go-playground/m/v1/usecase/rule"
)

// UserManagementUsecase ...
type UserManagementUsecase struct {
	port.IUserManagementRepository
	port.IBalanceRepository
	port.IDealHistoryRepository
	port.ITransactionManagementRepository
	userOutputPort port.UserOutput
}

// NewUserManagementUsecase ...
func NewUserManagementUsecase(
	umr port.IUserManagementRepository,
	br port.IBalanceRepository,
	thr port.IDealHistoryRepository,
	tmr port.ITransactionManagementRepository,
	uop port.UserOutput,
) *UserManagementUsecase {
	return &UserManagementUsecase{umr, br, thr, tmr, uop}
}

// CreateUser ...
func (u *UserManagementUsecase) CreateUser(ctx context.Context, inputUserCreate input.UserCreate, inputTopUpAmount uint) {
	// ユーザー重複確認
	if err := u.verifyThatNoUserHasSameEmail(ctx, inputUserCreate.EmailAddress); err != nil {
		u.userOutputPort.AppError(rule.InternalServerError)
		return
	}

	initialUser, err := user.InitGeneral(inputUserCreate.FirstName, inputUserCreate.LastName, inputUserCreate.Age, inputUserCreate.EmailAddress)
	if err != nil {
		u.userOutputPort.AppError(rule.InternalServerError)
		return
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
		u.userOutputPort.AppError(rule.InternalServerError)
		return
	}
	// u.userOutputPort.User()
}

// EditUser ...
func (u *UserManagementUsecase) EditUser(ctx context.Context, inputUserUpdate input.UserUpdate) {
	userID := user.ID(inputUserUpdate.ID)

	// ユーザー存在確認
	if err := u.verifyThatUserExist(ctx, userID); err != nil {
		u.userOutputPort.AppError(rule.InternalServerError)
		return
	}

	// ユーザー重複確認
	if err := u.verifyThatNoUserHasSameEmail(ctx, inputUserUpdate.EmailAddress); err != nil {
		u.userOutputPort.AppError(rule.InternalServerError)
		return
	}

	updateUser, err := user.UpdateGeneral(userID, inputUserUpdate.LastName, inputUserUpdate.EmailAddress, grade.ID(inputUserUpdate.GradeID))
	if err != nil {
		u.userOutputPort.AppError(rule.InternalServerError)
		return
	}

	// ユーザー情報更新
	_, err = u.editUser(ctx, updateUser)
	if err != nil {
		u.userOutputPort.AppError(rule.InternalServerError)
		return
	}
	// u.userOutputPort.User()
}

// RetrieveUserByCondition ...
func (u *UserManagementUsecase) RetrieveUserByCondition(ctx context.Context, id uint) {
	targetUser, err := u.fetchUserbyID(ctx, id)
	if err != nil {
		u.userOutputPort.AppError(rule.InternalServerError)
		return
	}
	if err := targetUser.Exist(true); err != nil {
		u.userOutputPort.AppError(rule.InternalServerError)
		return
	}
	out := &output.User{
		ID:           uint(targetUser.ID()),
		FirstName:    targetUser.FirstName(),
		LastName:     targetUser.LastName(),
		Age:          targetUser.Age(),
		EmailAddress: targetUser.EmailAddress(),
		GradeName:    targetUser.GradeName(),
	}
	u.userOutputPort.User(out)
}

// RetrieveUsers ...
func (u *UserManagementUsecase) RetrieveUsers(ctx context.Context) {
	targetUserList, err := u.fetchUserList(ctx)
	if err != nil {
		u.userOutputPort.AppError(rule.InternalServerError)
		return
	}
	outputList := make([]*output.User, len(targetUserList))
	for i, v := range targetUserList {
		outputList[i] = &output.User{
			ID:           uint(v.ID()),
			FirstName:    v.FirstName(),
			LastName:     v.LastName(),
			Age:          v.Age(),
			EmailAddress: v.EmailAddress(),
			GradeName:    v.GradeName(),
		}
	}
	u.userOutputPort.UserList(outputList)
}

func (u *UserManagementUsecase) verifyThatNoUserHasSameEmail(ctx context.Context, email string) error {
	targetUser, err := u.fetchUserbyEmail(ctx, email)
	if err != nil {
		return err
	}
	if err := targetUser.Exist(false); err != nil {
		return err
	}
	return nil
}

func (u *UserManagementUsecase) verifyThatUserExist(ctx context.Context, id user.ID) error {
	targetUser, err := u.fetchUserbyID(ctx, uint(id))
	if err != nil {
		return err
	}
	if err := targetUser.Exist(true); err != nil {
		return err
	}
	return nil
}

func (u *UserManagementUsecase) registerUser(ctx context.Context, generalUser *user.General) (*user.General, error) {
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

func (u *UserManagementUsecase) editUser(ctx context.Context, generalUser *user.General) (*user.General, error) {
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

func (u *UserManagementUsecase) fetchUserbyID(ctx context.Context, id uint) (*user.General, error) {
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

func (u *UserManagementUsecase) fetchUserbyEmail(ctx context.Context, email string) (*user.General, error) {
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

func (u *UserManagementUsecase) fetchUserList(ctx context.Context) (user.Generals, error) {
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
