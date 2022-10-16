package interactor

import (
	"context"
	"go-playground/m/v1/domain/model/balance"
	"go-playground/m/v1/domain/model/deal"
	"go-playground/m/v1/domain/model/user"
	"go-playground/m/v1/usecase/data/input"
	"go-playground/m/v1/usecase/data/output"
	"go-playground/m/v1/usecase/dto"
	"go-playground/m/v1/usecase/interactor/port"
	"go-playground/m/v1/usecase/rule"
	"log"
)

// BalanceControlUsecase ...
type BalanceControlUsecase struct {
	port.IBalanceRepository
	port.IDealHistoryRepository
	port.ITransactionManagementRepository
	balanceOutputPort port.BalanceOutput
}

// NewBalanceControlUsecase ...
func NewBalanceControlUsecase(br port.IBalanceRepository, th port.IDealHistoryRepository, tmr port.ITransactionManagementRepository, bop port.BalanceOutput) BalanceControlUsecase {
	return BalanceControlUsecase{br, th, tmr, bop}
}

// PutMoney 電子マネーのチャージ
func (u BalanceControlUsecase) PutMoney(ctx context.Context, userID uint, inputPuttingMoney input.PuttingMoney) {
	// 残高取得
	currentBalance, err := u.fetchBalanceByUserID(ctx, userID)
	if err != nil {
		u.balanceOutputPort.AppError(rule.InternalServerError)
		return
	}
	if err := currentBalance.Exist(true); err != nil {
		u.balanceOutputPort.AppError(rule.NotFound)
		return
	}

	topUpAmount := balance.TopUpAmount(inputPuttingMoney.Amount)

	// チャージ結果計算
	calculatedBalance, err := currentBalance.AddUp(topUpAmount)
	if err != nil {
		u.balanceOutputPort.AppError(rule.InternalServerError)
		return
	}

	// 同一トランザクション内処理開始
	if err := u.Transaction(ctx, func(ctx context.Context) (err error) {

		// 残高更新
		if err = u.UpdateBalance(ctx, dto.NewUpdateBalance(calculatedBalance.UserID(), calculatedBalance.RemainingAmount())); err != nil {
			return err
		}

		// 取引履歴登録
		dealHistory := deal.InitTopUpHistory(topUpAmount)
		if err = u.RegisterDealHistory(ctx, dto.NewRegisterDealHistory(user.ID(userID), dealHistory.ItemName(), dealHistory.Amount())); err != nil {
			return err
		}
		return
	}); err != nil {
		u.balanceOutputPort.AppError(rule.InternalServerError)
		return
	}
}

// PayMoney ...
func (u BalanceControlUsecase) PayMoney(ctx context.Context, userID uint, inputPayment input.Payment) {
	// 残高取得
	currentBalance, err := u.fetchBalanceByUserID(ctx, userID)
	if err != nil {
		u.balanceOutputPort.AppError(rule.InternalServerError)
		return
	}
	if err := currentBalance.Exist(true); err != nil {
		u.balanceOutputPort.AppError(rule.InternalServerError)
		return
	}

	paymentAmount := balance.PaymentAmount(inputPayment.Amount)

	// 支払結果計算
	calculatedBalance, err := currentBalance.Subtract(paymentAmount)
	if err != nil {
		log.Println(calculatedBalance)
		u.balanceOutputPort.AppError(rule.ShortBalance)
		return
	}

	// 同一トランザクション内処理開始
	if err := u.Transaction(ctx, func(ctx context.Context) (err error) {

		// 残高更新
		if err = u.UpdateBalance(ctx, dto.NewUpdateBalance(calculatedBalance.UserID(), calculatedBalance.RemainingAmount())); err != nil {
			return err
		}

		// 取引履歴登録
		dealHistory := deal.InitPaymentHistory(inputPayment.ItemName, paymentAmount)
		if err = u.RegisterDealHistory(ctx, dto.NewRegisterDealHistory(user.ID(userID), dealHistory.ItemName(), dealHistory.Amount())); err != nil {
			return err
		}
		return
	}); err != nil {
		u.balanceOutputPort.AppError(rule.InternalServerError)
		return
	}
}

// RetrieveRemainingBalanceByUserID 電子マネーの残高確認
func (u BalanceControlUsecase) RetrieveRemainingBalanceByUserID(ctx context.Context, userID uint) {
	tragetBalance, err := u.fetchBalanceByUserID(ctx, userID)
	if err != nil {
		u.balanceOutputPort.AppError(rule.InternalServerError)
		return
	}
	if err := tragetBalance.Exist(true); err != nil {
		u.balanceOutputPort.AppError(rule.InternalServerError)
		return
	}
	out := &output.Balance{
		Amount: output.Amount(tragetBalance.RemainingAmount()),
	}
	u.balanceOutputPort.Balance(out)
}

func (u BalanceControlUsecase) fetchBalanceByUserID(ctx context.Context, userID uint) (balance.Entity, error) {
	fetchResult, err := u.FetchBalanceByUserID(ctx, userID)
	if err != nil {
		return balance.Entity{}, err
	}
	return fetchResult.ToBalanceModel(), nil
}
