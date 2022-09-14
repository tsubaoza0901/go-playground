package dependency

import "go-playground/m/v1/src/adapters/controllers/http/handler"

// 簡易DIコンテナ（Handler用）

// HTTPHandler ...
type HTTPHandler struct {
	handler.UserHandler
	handler.GradeHandler
	handler.TransactionHistoryHandler
	handler.BalanceControlHandler
}

// InitHTTPHandler ...
func (i Injection) InitHTTPHandler() HTTPHandler {
	return HTTPHandler{
		UserHandler:               i.InitUserHandler(),
		GradeHandler:              i.InitGradeHandler(),
		TransactionHistoryHandler: i.InitTransactionHistoryHandler(),
		BalanceControlHandler:     i.InitBalanceControlHandler(),
	}
}

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

// InitTransactionHistoryHandler ...
func (i Injection) InitTransactionHistoryHandler() handler.TransactionHistoryHandler {
	return handler.NewTransactionHistoryHandler(
		i.InitTransactionUsecase(),
	)
}

// InitBalanceControlHandler ...
func (i Injection) InitBalanceControlHandler() handler.BalanceControlHandler {
	return handler.NewBalanceControlHandler(
		i.InitBalanceControlUsecase(),
	)
}
