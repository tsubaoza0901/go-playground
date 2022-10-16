package repository

import (
	"context"
	dbModel "go-playground/m/v1/infrastructure/rdb/model"
	"go-playground/m/v1/repository/rdb"
	"go-playground/m/v1/usecase/dto"
)

// GradeRepository ...
type GradeRepository struct {
	rdb.IManageDBConn
}

// NewGradeRepository ...
func NewGradeRepository(mdc rdb.IManageDBConn) GradeRepository {
	return GradeRepository{mdc}
}

// FetchGradeList ...
func (r GradeRepository) FetchGradeList(ctx context.Context) (*dto.FetchGradeListResult, error) {
	grades := new(dbModel.Grades)
	if err := r.GetConnection(ctx).Find(&grades).Error; err != nil {
		return nil, err
	}
	return dbModel.MakeFetchGradeListResultDTO(*grades), nil
}
