package dependency

import "go-playground/m/v1/src/adapters/gateways/repository"

// 簡易DIコンテナ（Repository用）

// InitTransactionManagementRepository ...
func (i Injection) InitTransactionManagementRepository() repository.TransactionManagementRepository {
	return repository.NewTransactionManagementRepository(i.InitManageDBConn())
}

// InitUserRepository ...
func (i Injection) InitUserRepository() repository.UserRepository {
	return repository.NewUserRepository(i.InitManageDBConn())
}

// InitGradeRepository ...
func (i Injection) InitGradeRepository() repository.GradeRepository {
	return repository.NewGradeRepository(i.InitManageDBConn())
}

// InitDealRepository ...
func (i Injection) InitDealRepository() repository.DealHistoryRepository {
	return repository.NewDealHistoryRepository(i.InitManageDBConn())
}

// InitBalanceRepository ...
func (i Injection) InitBalanceRepository() repository.BalanceRepository {
	return repository.NewBalanceRepository(i.InitManageDBConn())
}
