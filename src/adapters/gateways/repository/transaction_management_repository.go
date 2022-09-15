package repository

import (
	"gorm.io/gorm"
)

// TransactionManagementRepository ...
type TransactionManagementRepository struct {
	dbConn *gorm.DB
}

// NewTransactionManagementRepository ...
func NewTransactionManagementRepository(conn *gorm.DB) TransactionManagementRepository {
	return TransactionManagementRepository{conn}
}

// BeginConnection ...
func (r TransactionManagementRepository) BeginConnection() *gorm.DB {
	return r.dbConn.Begin()
}

// Commit ...
func (r TransactionManagementRepository) Commit(tx *gorm.DB) {
	tx.Commit()
}

// Rollback ...
func (r TransactionManagementRepository) Rollback(tx *gorm.DB) {
	tx.Rollback()
}
