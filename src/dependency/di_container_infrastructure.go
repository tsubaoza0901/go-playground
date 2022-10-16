package dependency

import (
	"go-playground/m/v1/infrastructure/persistence"
	"go-playground/m/v1/infrastructure/rdb"
)

// 簡易DIコンテナ（infrastructure用）

// InitTransactionManagement ...
func (i Injection) InitTransactionManagement() persistence.TransactionManagement {
	return persistence.NewTransactionManagement(i.InitManageDBConn())
}

// InitUser ...
func (i Injection) InitUser() persistence.User {
	return persistence.NewUser(i.InitManageDBConn())
}

// InitGrade ...
func (i Injection) InitGrade() persistence.Grade {
	return persistence.NewGrade(i.InitManageDBConn())
}

// InitDeal ...
func (i Injection) InitDeal() persistence.DealHistory {
	return persistence.NewDealHistory(i.InitManageDBConn())
}

// InitBalance ...
func (i Injection) InitBalance() persistence.Balance {
	return persistence.NewBalance(i.InitManageDBConn())
}

// InitManageDBConn ...
func (i Injection) InitManageDBConn() rdb.ManageDBConn {
	return rdb.NewManageDBConn(i.db)
}
