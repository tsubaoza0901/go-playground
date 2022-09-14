package interactor

import (
	"context"
	"go-playground/m/v1/src/domain/model/balance"
	"go-playground/m/v1/src/domain/model/transaction"
	"go-playground/m/v1/src/domain/repository"
	"go-playground/m/v1/src/usecase/data/input"
	"go-playground/m/v1/src/usecase/data/output"
)

// BalanceControlUsecase ...
type BalanceControlUsecase struct {
	repository.IBalanceRepository
	repository.ITransactionHistoryRepository
}

// NewBalanceControlUsecase ...
func NewBalanceControlUsecase(br repository.IBalanceRepository, th repository.ITransactionHistoryRepository) BalanceControlUsecase {
	return BalanceControlUsecase{br, th}
}

// PutMoney 電子マネーのチャージ
func (u BalanceControlUsecase) PutMoney(ctx context.Context, userID uint, inputPuttingMoney input.PuttingMoney) error {
	// 残高取得
	balanceFetchAmountDTO, err := u.FetchBalanceByUserID(ctx, userID)
	if err != nil {
		return err
	}
	remainingAmount := balanceFetchAmountDTO.RemainingAmount

	// 金額計算
	calculatedRemainingAmount, err := remainingAmount.AddUp(balance.TopUpAmount(inputPuttingMoney.Amount))
	if err != nil {
		return err
	}

	// チャージ実行
	if err := u.UpdateBalance(ctx, userID, balance.NewUpdateBalanceDTO(*calculatedRemainingAmount)); err != nil {
		return err
	}

	// 取引履歴登録
	transactionHistory := transaction.NewHistory("", uint(inputPuttingMoney.Amount))
	if err := u.RegisterTransactionHistory(ctx, transaction.NewRegisterHistoryDTO(*transactionHistory, userID)); err != nil {
		return err
	}
	return nil
}

// PayMoney ...
func (u BalanceControlUsecase) PayMoney(ctx context.Context, userID uint, inputPayment input.Payment) error {
	// 残高取得
	balanceFetchAmountDTO, err := u.FetchBalanceByUserID(ctx, userID)
	if err != nil {
		return err
	}
	remainingAmount := balanceFetchAmountDTO.RemainingAmount

	// 金額計算
	calculatedRemainingAmount, err := remainingAmount.Subtract(balance.PaymentAmount(inputPayment.Amount))
	if err != nil {
		return err
	}

	// 支払実行
	if err := u.UpdateBalance(ctx, userID, balance.NewUpdateBalanceDTO(*calculatedRemainingAmount)); err != nil {
		return err
	}

	// 履歴登録
	transactionHistory := transaction.NewHistory(inputPayment.ItemName, uint(inputPayment.Amount))
	if err := u.RegisterTransactionHistory(ctx, transaction.NewRegisterHistoryDTO(*transactionHistory, userID)); err != nil {
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
