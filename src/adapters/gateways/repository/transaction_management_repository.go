package repository

import (
	"context"
	"go-playground/m/v1/src/adapters/gateways/persistance/rdb"
)

// TransactionManagementRepository ...
type TransactionManagementRepository struct {
	rdb.IManageDBConn
}

// NewTransactionManagementRepository ...
func NewTransactionManagementRepository(mdc rdb.IManageDBConn) TransactionManagementRepository {
	return TransactionManagementRepository{mdc}
}

// Transaction ...
func (r TransactionManagementRepository) Transaction(ctx context.Context, fc func(context.Context) error) (err error) {
	return r.StartTransaction(ctx, fc)
}
