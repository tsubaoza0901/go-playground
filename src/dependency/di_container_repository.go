package dependency

import (
	"go-playground/m/v1/repository"
)

// 簡易DIコンテナ（repository用）

// InitTransactionManagementRepository ...
func (i Injection) InitTransactionManagementRepository() repository.TransactionManagement {
	return repository.NewTransactionManagement(
		i.InitTransactionManagement(),
	)
}

// InitUserRepository ...
func (i Injection) InitUserRepository() repository.User {
	return repository.NewUser(
		i.InitUser(),
	)
}

// InitGradeRepository ...
func (i Injection) InitGradeRepository() repository.Grade {
	return repository.NewGrade(
		i.InitGrade(),
	)
}

// InitDealRepository ...
func (i Injection) InitDealRepository() repository.DealHistory {
	return repository.NewDealHistory(
		i.InitDeal(),
	)
}

// InitBalanceRepository ...
func (i Injection) InitBalanceRepository() repository.Balance {
	return repository.NewBalance(
		i.InitBalance(),
	)
}
