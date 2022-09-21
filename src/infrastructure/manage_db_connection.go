package infrastructure

// type contextKey string

// const transactionKey contextKey = "transaction"

// // ManageDBConn ...
// type ManageDBConn struct {
// 	db *gorm.DB
// }

// // NewManageDBConn ...
// func NewManageDBConn(conn *gorm.DB) ManageDBConn {
// 	return ManageDBConn{conn}
// }

// // StartTransaction ...
// func (r ManageDBConn) StartTransaction(ctx context.Context, fc func(context.Context) error) (err error) {
// 	tx := r.db.Begin()
// 	if tx.Error != nil {
// 		return tx.Error
// 	}

// 	defer func() {
// 		if err != nil {
// 			tx.Rollback()
// 		}
// 	}()
// 	ctx = context.WithValue(ctx, transactionKey, tx)
// 	if err = fc(ctx); err != nil {
// 		return err
// 	}
// 	if err = tx.Commit().Error; err != nil {
// 		return err
// 	}
// 	return nil
// }

// // GetConnection ...
// func (r ManageDBConn) GetConnection(ctx context.Context) *gorm.DB {
// 	tx, ok := r.getTxFromContext(ctx)
// 	if !ok {
// 		return r.db
// 	}
// 	return tx
// }

// func (r ManageDBConn) getTxFromContext(ctx context.Context) (*gorm.DB, bool) {
// 	tx, ok := ctx.Value(transactionKey).(*gorm.DB)
// 	return tx, ok
// }
