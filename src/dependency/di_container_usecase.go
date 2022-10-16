package dependency

import (
	"go-playground/m/v1/usecase/interactor"
	"go-playground/m/v1/usecase/interactor/port"
)

// 簡易DIコンテナ（Usecase用）

// InitUserManagementUsecase ...
func (i Injection) InitUserManagementUsecase(userOutputPort port.UserOutput) *interactor.UserManagementUsecase {
	return interactor.NewUserManagementUsecase(
		i.InitUserRepository(),
		i.InitBalanceRepository(),
		i.InitDealRepository(),
		i.InitTransactionManagementRepository(),
		userOutputPort,
	)
}

// InitGradeUsecase ...
func (i Injection) InitGradeUsecase(gradeOutputPort port.GradeOutput) *interactor.GradeUsecase {
	return interactor.NewGradeUsecase(
		i.InitGradeRepository(),
		gradeOutputPort,
	)
}

// InitDealHistoryUsecase ...
func (i Injection) InitDealHistoryUsecase(dealHistoryOutputPort port.DealHistoryOutput) *interactor.DealHistoryUsecase {
	return interactor.NewDealHistoryUsecase(
		i.InitDealRepository(),
		dealHistoryOutputPort,
	)
}

// InitBalanceControlUsecase ...
func (i Injection) InitBalanceControlUsecase(balanceOutputPort port.BalanceOutput) interactor.BalanceControlUsecase {
	return interactor.NewBalanceControlUsecase(
		i.InitBalanceRepository(),
		i.InitDealRepository(),
		i.InitTransactionManagementRepository(),
		balanceOutputPort,
	)
}
