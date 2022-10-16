//go:generate mockgen -source=$GOFILE -package=mock -destination=$GOPATH/src/mock/$GOFILE

package port

import "go-playground/m/v1/usecase/data/output"

// GradeOutput ...
type GradeOutput interface {
	GradeList([]*output.Grade)
	AppError(errorCode int)
}

// DealHistoryOutput ...
type DealHistoryOutput interface {
	DealHistoryList([]*output.DealHistory)
	AppError(errorCode int)
}

// UserOutput ...
type UserOutput interface {
	// User(*output.User)
	User(*output.User)
	UserList([]*output.User)
	AppError(errorCode int)
}

// BalanceOutput ...
type BalanceOutput interface {
	Balance(*output.Balance)
	AppError(errorCode int)
}
