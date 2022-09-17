package dependency

import (
	"go-playground/m/v1/src/adapters/controllers/http/handler"
)

// 簡易DIコンテナ（Handler用）

// InitUserHandler ...
func (i Injection) InitUserHandler() handler.UserHandler {
	return handler.NewUserHandler(
		i.InitUserManagementUsecase(),
	)
}

// InitGradeHandler ...
func (i Injection) InitGradeHandler() handler.GradeHandler {
	return handler.NewGradeHandler(
		i.InitGradeUsecase(),
	)
}

// InitDealHistoryHandler ...
func (i Injection) InitDealHistoryHandler() handler.DealHistoryHandler {
	return handler.NewDealHistoryHandler(
		i.InitDealUsecase(),
	)
}

// InitBalanceControlHandler ...
func (i Injection) InitBalanceControlHandler() handler.BalanceControlHandler {
	return handler.NewBalanceControlHandler(
		i.InitBalanceControlUsecase(),
	)
}
