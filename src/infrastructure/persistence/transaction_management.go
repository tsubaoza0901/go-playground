package persistence

import (
	"context"
	"go-playground/m/v1/infrastructure/rdb"
)

// TransactionManagement ...
type TransactionManagement struct {
	rdb.ManageDBConn
}

// NewTransactionManagement ...
func NewTransactionManagement(mdc rdb.ManageDBConn) TransactionManagement {
	return TransactionManagement{mdc}
}

// Transaction ...
func (r TransactionManagement) Transaction(ctx context.Context, fc func(context.Context) error) (err error) {
	return r.StartTransaction(ctx, fc)
}
