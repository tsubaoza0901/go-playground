package dependency

import "go-playground/m/v1/src/adapters/gateways/repository"

// 簡易DIコンテナ（Repository用）

// InitUserRepository ...
func (i Injection) InitUserRepository() repository.UserRepository {
	return repository.NewUserRepository(i.dbConn)
}

// InitGradeRepository ...
func (i Injection) InitGradeRepository() repository.GradeRepository {
	return repository.NewGradeRepository(i.dbConn)
}

// InitTransactionRepository ...
func (i Injection) InitTransactionRepository() repository.TransactionHistoryRepository {
	return repository.NewTransactionHistoryRepository(i.dbConn)
}

// InitBalanceRepository ...
func (i Injection) InitBalanceRepository() repository.BalanceRepository {
	return repository.NewBalanceRepository(i.dbConn)
}
