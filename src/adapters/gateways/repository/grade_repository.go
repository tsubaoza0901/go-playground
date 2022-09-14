package repository

import (
	"context"
	dbModel "go-playground/m/v1/src/adapters/gateways/persistance/rdb/model"
	"go-playground/m/v1/src/domain/model/grade"

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

// FetchAllGrades ...
func (r GradeRepository) FetchAllGrades(ctx context.Context) (*grade.FetchAllDTO, error) {
	grades := new(dbModel.Grades)
	if err := r.dbConn.Find(&grades).Error; err != nil {
		return nil, err
	}
	return dbModel.MakeFetchAllDTO(*grades), nil
}
