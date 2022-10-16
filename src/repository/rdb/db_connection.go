//go:generate mockgen -source=$GOFILE -package=mock -destination=$GOPATH/src/mock/$GOFILE

package rdb

import (
	"context"

	"gorm.io/gorm"
)

// IManageDBConn ...
type IManageDBConn interface {
	StartTransaction(ctx context.Context, fc func(context.Context) error) (err error)
	GetConnection(ctx context.Context) *gorm.DB
}
