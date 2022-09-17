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

// BalanceControlUsecase ...
type BalanceControlUsecase struct {
	repository.IBalanceRepository
	repository.IDealHistoryRepository
	repository.ITransactionManagementRepository
}

// NewBalanceControlUsecase ...
func NewBalanceControlUsecase(br repository.IBalanceRepository, th repository.IDealHistoryRepository, tmr repository.ITransactionManagementRepository) BalanceControlUsecase {
	return BalanceControlUsecase{br, th, tmr}
}

// PutMoney 電子マネーのチャージ
func (u BalanceControlUsecase) PutMoney(ctx context.Context, userID uint, inputPuttingMoney input.PuttingMoney) error {
	// 残高取得
	currentBalance, err := u.fetchBalanceByUserID(ctx, userID)
	if err != nil {
		return err
	}
	if err := currentBalance.Exist(true); err != nil {
		return err
	}

	topUpAmount := balance.TopUpAmount(inputPuttingMoney.Amount)

	// チャージ結果計算
	calculatedBalance, err := currentBalance.AddUp(topUpAmount)
	if err != nil {
		return err
	}

	// 同一トランザクション内処理開始
	if err := u.Transaction(ctx, func(ctx context.Context) (err error) {

		// 残高更新
		if err = u.UpdateBalance(ctx, dto.NewUpdateBalance(calculatedBalance.UserID(), calculatedBalance.RemainingAmount())); err != nil {
			return err
		}

		// 取引履歴登録
		dealHistory := deal.NewTopUpHistory(topUpAmount)
		if err = u.RegisterDealHistory(ctx, dto.NewRegisterDealHistory(user.ID(userID), dealHistory.ItemName(), dealHistory.Amount())); err != nil {
			return err
		}
		return
	}); err != nil {
		return err
	}

	return nil
}

// PayMoney ...
func (u BalanceControlUsecase) PayMoney(ctx context.Context, userID uint, inputPayment input.Payment) error {
	// 残高取得
	currentBalance, err := u.fetchBalanceByUserID(ctx, userID)
	if err != nil {
		return err
	}
	if err := currentBalance.Exist(true); err != nil {
		return err
	}

	paymentAmount := balance.PaymentAmount(inputPayment.Amount)

	// 支払結果計算
	calculatedBalance, err := currentBalance.Subtract(paymentAmount)
	if err != nil {
		return err
	}

	// 同一トランザクション内処理開始
	if err := u.Transaction(ctx, func(ctx context.Context) (err error) {

		// 残高更新
		if err = u.UpdateBalance(ctx, dto.NewUpdateBalance(calculatedBalance.UserID(), calculatedBalance.RemainingAmount())); err != nil {
			return err
		}

		// 取引履歴登録
		dealHistory := deal.NewPaymentHistory(inputPayment.ItemName, paymentAmount)
		if err = u.RegisterDealHistory(ctx, dto.NewRegisterDealHistory(user.ID(userID), dealHistory.ItemName(), dealHistory.Amount())); err != nil {
			return err
		}
		return
	}); err != nil {
		return err
	}

	return nil
}

// RetrieveRemainingBalanceByUserID 電子マネーの残高確認
func (u BalanceControlUsecase) RetrieveRemainingBalanceByUserID(ctx context.Context, userID uint) (output.Balance, error) {
	tragetBalance, err := u.fetchBalanceByUserID(ctx, userID)
	if err != nil {
		return output.Balance{}, err
	}
	if err := tragetBalance.Exist(true); err != nil {
		return output.Balance{}, err
	}
	return output.MakeBalance(tragetBalance.RemainingAmount()), nil
}

func (u BalanceControlUsecase) fetchBalanceByUserID(ctx context.Context, userID uint) (balance.Entity, error) {
	fetchResult, err := u.FetchBalanceByUserID(ctx, userID)
	if err != nil {
		return balance.Entity{}, err
	}
	return fetchResult.ToBalanceModel(), nil
}
