package interactor

import (
	"context"
	"go-playground/m/v1/src/domain/model/balance"
	"go-playground/m/v1/src/domain/model/deal"
	"go-playground/m/v1/src/usecase/data/input"
	"go-playground/m/v1/src/usecase/data/output"
	"go-playground/m/v1/src/usecase/repository"
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
	balanceFetchAmountDTO, err := u.FetchBalanceByUserID(ctx, userID)
	if err != nil {
		return err
	}
	remainingAmount := balanceFetchAmountDTO.RemainingAmount

	// チャージ結果計算
	calculatedRemainingAmount, err := remainingAmount.AddUp(balance.TopUpAmount(inputPuttingMoney.Amount))
	if err != nil {
		return err
	}

	// 同一トランザクション内処理開始
	if err := u.Transaction(ctx, func(ctx context.Context) (err error) {

		// 残高更新
		if err = u.UpdateBalance(ctx, userID, balance.NewUpdateBalanceDTO(*calculatedRemainingAmount)); err != nil {
			return err
		}

		// 取引履歴登録
		dealHistory := deal.NewHistory("", uint(inputPuttingMoney.Amount))
		if err = u.RegisterDealHistory(ctx, deal.NewRegisterHistoryDTO(*dealHistory, userID)); err != nil {
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
	var err error

	// 残高取得
	balanceFetchAmountDTO, err := u.FetchBalanceByUserID(ctx, userID)
	if err != nil {
		return err
	}
	remainingAmount := balanceFetchAmountDTO.RemainingAmount

	// 支払結果計算
	calculatedRemainingAmount, err := remainingAmount.Subtract(balance.PaymentAmount(inputPayment.Amount))
	if err != nil {
		return err
	}

	// 同一トランザクション内処理開始
	if err := u.Transaction(ctx, func(ctx context.Context) (err error) {

		// 残高更新
		if err = u.UpdateBalance(ctx, userID, balance.NewUpdateBalanceDTO(*calculatedRemainingAmount)); err != nil {
			return err
		}

		// 取引履歴登録
		dealHistory := deal.NewHistory(inputPayment.ItemName, uint(inputPayment.Amount))
		if err = u.RegisterDealHistory(ctx, deal.NewRegisterHistoryDTO(*dealHistory, userID)); err != nil {
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
	balanceFetchAmountDTO, err := u.FetchBalanceByUserID(ctx, userID)
	if err != nil {
		return output.Balance{}, err
	}
	remainingAmount := balanceFetchAmountDTO.RemainingAmount
	return output.MakeBalance(remainingAmount), nil
}
