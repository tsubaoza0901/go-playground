package interactor

import (
	"context"
	"go-playground/m/v1/domain/model/grade"
	"go-playground/m/v1/usecase/data/output"
	"go-playground/m/v1/usecase/interactor/port"
	"go-playground/m/v1/usecase/rule"
)

// GradeUsecase ...
type GradeUsecase struct {
	port.IGradeRepository
	gradeOutputPort port.GradeOutput
}

// NewGradeUsecase ...
func NewGradeUsecase(gr port.IGradeRepository, gop port.GradeOutput) *GradeUsecase {
	return &GradeUsecase{gr, gop}
}

// RetrieveGrades ...
func (u *GradeUsecase) RetrieveGrades(ctx context.Context) {
	targetGradeList, err := u.fetchGradeList(ctx)
	if err != nil {
		u.gradeOutputPort.AppError(rule.InternalServerError)
		return
	}
	outputList := make([]*output.Grade, len(targetGradeList))
	for i, v := range targetGradeList {
		outputList[i] = &output.Grade{
			ID:   uint(v.ID()),
			Name: v.Name(),
		}
	}
	u.gradeOutputPort.GradeList(outputList)
}

func (u *GradeUsecase) fetchGradeList(ctx context.Context) (grade.Entities, error) {
	fetchResult, err := u.FetchGradeList(ctx)
	if err != nil {
		return nil, err
	}
	return fetchResult.ToGradesModel(), nil
}
