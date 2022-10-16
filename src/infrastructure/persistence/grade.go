package persistence

import (
	"context"
	"go-playground/m/v1/infrastructure/rdb"
	dbModel "go-playground/m/v1/infrastructure/rdb/model"
	"go-playground/m/v1/usecase/dto"
)

// Grade ...
type Grade struct {
	rdb.ManageDBConn
}

// NewGrade ...
func NewGrade(mdc rdb.ManageDBConn) Grade {
	return Grade{mdc}
}

// FetchGradeList ...
func (r Grade) FetchGradeList(ctx context.Context) (*dto.FetchGradeListResult, error) {
	grades := new(dbModel.Grades)
	if err := r.GetConnection(ctx).Find(&grades).Error; err != nil {
		return nil, err
	}
	return dbModel.MakeFetchGradeListResultDTO(*grades), nil
}
