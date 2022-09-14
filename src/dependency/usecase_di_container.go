package dependency

import (
	"go-playground/m/v1/src/usecase/interactor.go"
)

// 簡易DIコンテナ（Usecase用）

// InitUserManagementUsecase ...
func (i Injection) InitUserManagementUsecase() interactor.UserManagementUsecase {
	return interactor.NewUserManagementUsecase(
		i.InitUserRepository(),
		i.InitBalanceRepository(),
		i.InitDealRepository(),
	)
}

// InitGradeUsecase ...
func (i Injection) InitGradeUsecase() interactor.GradeUsecase {
	return interactor.NewGradeUsecase(
		i.InitGradeRepository(),
	)
}

// InitDealUsecase ...
func (i Injection) InitDealUsecase() interactor.DealUsecase {
	return interactor.NewDealUsecase(
		i.InitDealRepository(),
	)
}

// InitBalanceControlUsecase ...
func (i Injection) InitBalanceControlUsecase() interactor.BalanceControlUsecase {
	return interactor.NewBalanceControlUsecase(
		i.InitBalanceRepository(),
		i.InitDealRepository(),
	)
}
