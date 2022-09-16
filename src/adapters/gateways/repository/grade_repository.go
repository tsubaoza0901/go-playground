package repository

import (
	"context"
	dbModel "go-playground/m/v1/src/adapters/gateways/persistance/rdb/model"
	"go-playground/m/v1/src/usecase/repository/dto"

	"gorm.io/gorm"
)

// GradeRepository ...
type GradeRepository struct {
	dbConn *gorm.DB
}

// NewGradeRepository ...
func NewGradeRepository(conn *gorm.DB) GradeRepository {
	return GradeRepository{
		dbConn: conn,
	}
}

// FetchGradeList ...
func (r GradeRepository) FetchGradeList(ctx context.Context) (*dto.FetchGradeListResult, error) {
	grades := new(dbModel.Grades)
	if err := r.dbConn.Find(&grades).Error; err != nil {
		return nil, err
	}
	return dbModel.MakeFetchGradeListResultDTO(*grades), nil
}
