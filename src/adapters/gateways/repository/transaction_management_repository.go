package repository

import (
	"context"
	"go-playground/m/v1/src/usecase"

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

// // NewContextWithTx ...
// func (r TransactionManagementRepository) NewContextWithTx(ctx context.Context) context.Context {
// 	tx := r.dbConn.Begin()
// 	return context.WithValue(ctx, transaction.Key, tx)
// }

// // CommitByContext ...
// func (r TransactionManagementRepository) CommitByContext(ctx context.Context) error {
// 	tx, ok := getTxFromContext(ctx)
// 	if !ok {
// 		return errors.New("アサーション失敗によりトランザクション取得不可")
// 	}
// 	return tx.Commit().Error
// }

// // RollbackByContext ...
// func (r TransactionManagementRepository) RollbackByContext(ctx context.Context) error {
// 	tx, ok := getTxFromContext(ctx)
// 	if !ok {
// 		return errors.New("アサーション失敗によりトランザクション取得不可")
// 	}
// 	return tx.Rollback().Error
// }

// Transaction ...
func (r TransactionManagementRepository) Transaction(ctx context.Context, fc func(context.Context) error) (err error) {
	tx := r.dbConn.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()
	ctx = context.WithValue(ctx, usecase.TransactionKey, tx)
	if err = fc(ctx); err != nil {
		return err
	}
	if err = tx.Commit().Error; err != nil {
		return err
	}
	return nil
}

func getTxFromContext(ctx context.Context) (*gorm.DB, bool) {
	tx, ok := ctx.Value(usecase.TransactionKey).(*gorm.DB)
	return tx, ok
}

// // TransactionFinisher defer内で呼び出されるトランザクション終了処理メソッド
// // 注）使用する際は必ず以下について確認すること
// // ・呼び出し側の関数スコープ内でグローバルに使用できるerr（error型の変数）が存在すること
// // →（例）Named return valueとして設定されたerr、対象関数スコープ内全体で使用できるように定義されたerr など
// func (r TransactionManagementRepository) TransactionFinisher(tx *gorm.DB, err error) {
// 	if err != nil {
// 		r.rollback(tx)
// 	} else {
// 		r.commit(tx)
// 	}
// }
