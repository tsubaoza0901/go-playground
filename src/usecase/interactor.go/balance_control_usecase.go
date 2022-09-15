package interactor

import (
	"context"
	"go-playground/m/v1/src/domain/model/balance"
	"go-playground/m/v1/src/domain/model/deal"
	"go-playground/m/v1/src/domain/repository"
	"go-playground/m/v1/src/usecase/data/input"
	"go-playground/m/v1/src/usecase/data/output"
	"log"
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

	// 残高更新
	tx := u.BeginConnection()

	if err := u.UpdateBalanceWithTx(ctx, tx, userID, balance.NewUpdateBalanceDTO(*calculatedRemainingAmount)); err != nil {
		u.Rollback(tx)
		return err
	}

	// 取引履歴登録
	dealHistory := deal.NewHistory("", uint(inputPuttingMoney.Amount))
	if err := u.RegisterDealHistoryWithTx(ctx, tx, deal.NewRegisterHistoryDTO(*dealHistory, userID)); err != nil {
		log.Printf("%+v", tx)
		u.Rollback(tx)
		return err
	}

	u.Commit(tx)
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

	// 支払結果計算
	calculatedRemainingAmount, err := remainingAmount.Subtract(balance.PaymentAmount(inputPayment.Amount))
	if err != nil {
		return err
	}

	// 残高更新
	tx := u.BeginConnection()

	if err := u.UpdateBalanceWithTx(ctx, tx, userID, balance.NewUpdateBalanceDTO(*calculatedRemainingAmount)); err != nil {
		u.Rollback(tx)
		return err
	}

	// 取引履歴登録
	dealHistory := deal.NewHistory(inputPayment.ItemName, uint(inputPayment.Amount))
	if err := u.RegisterDealHistoryWithTx(ctx, tx, deal.NewRegisterHistoryDTO(*dealHistory, userID)); err != nil {
		u.Rollback(tx)
		return err
	}

	u.Commit(tx)

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
