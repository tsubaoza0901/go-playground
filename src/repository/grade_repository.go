package repository

import (
	"context"
	"go-playground/m/v1/repository/persistence"
	"go-playground/m/v1/usecase/dto"
)

// Grade ...
type Grade struct {
	gradeDataAccess persistence.GradeDataAccess
}

// NewGrade ...
func NewGrade(gda persistence.GradeDataAccess) Grade {
	return Grade{gda}
}

// FetchGradeList ...
func (r Grade) FetchGradeList(ctx context.Context) (*dto.FetchGradeListResult, error) {
	return r.gradeDataAccess.FetchGradeList(ctx)
}
