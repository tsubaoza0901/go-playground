package repository

import (
	"context"
	"go-playground/m/v1/repository/persistence"
)

// TransactionManagement ...
type TransactionManagement struct {
	transactionManagementDataAccess persistence.TransactionManagementDataAccess
}

// NewTransactionManagement ...
func NewTransactionManagement(tmda persistence.TransactionManagementDataAccess) TransactionManagement {
	return TransactionManagement{tmda}
}

// Transaction ...
func (r TransactionManagement) Transaction(ctx context.Context, fc func(context.Context) error) error {
	return r.transactionManagementDataAccess.Transaction(ctx, fc)
}
