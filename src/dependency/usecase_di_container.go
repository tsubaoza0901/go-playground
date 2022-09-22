package dependency

import (
	"go-playground/m/v1/usecase/interactor"
)

// 簡易DIコンテナ（Usecase用）

// InitUserManagementUsecase ...
func (i Injection) InitUserManagementUsecase() interactor.UserManagementUsecase {
	return interactor.NewUserManagementUsecase(
		i.InitUserRepository(),
		i.InitBalanceRepository(),
		i.InitDealRepository(),
		i.InitTransactionManagementRepository(),
	)
}

// InitGradeUsecase ...
func (i Injection) InitGradeUsecase() interactor.GradeUsecase {
	return interactor.NewGradeUsecase(
		i.InitGradeRepository(),
	)
}

// InitDealUsecase ...
func (i Injection) InitDealUsecase() interactor.DealHistoryUsecase {
	return interactor.NewDealHistoryUsecase(
		i.InitDealRepository(),
	)
}

// InitBalanceControlUsecase ...
func (i Injection) InitBalanceControlUsecase() interactor.BalanceControlUsecase {
	return interactor.NewBalanceControlUsecase(
		i.InitBalanceRepository(),
		i.InitDealRepository(),
		i.InitTransactionManagementRepository(),
	)
}
