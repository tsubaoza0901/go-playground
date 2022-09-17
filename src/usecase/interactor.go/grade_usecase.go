package interactor

import (
	"context"
	"go-playground/m/v1/src/domain/model/grade"
	"go-playground/m/v1/src/usecase/data/output"
	"go-playground/m/v1/src/usecase/repository"
)

// GradeUsecase ...
type GradeUsecase struct {
	repository.IGradeRepository
}

// NewGradeUsecase ...
func NewGradeUsecase(r repository.IGradeRepository) GradeUsecase {
	return GradeUsecase{r}
}

// RetrieveGrades ...
func (u GradeUsecase) RetrieveGrades(ctx context.Context) (*output.Grades, error) {
	targetGradeList, err := u.fetchGradeList(ctx)
	if err != nil {
		return nil, err
	}
	return output.MakeGrades(targetGradeList), nil
}

func (u GradeUsecase) fetchGradeList(ctx context.Context) (grade.Entities, error) {
	fetchResult, err := u.FetchGradeList(ctx)
	if err != nil {
		return nil, err
	}
	return fetchResult.ToGradesModel(), nil
}
