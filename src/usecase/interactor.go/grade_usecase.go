package interactor

import (
	"context"
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
	grades, err := u.FetchAllGrades(ctx)
	if err != nil {
		return nil, err
	}
	return output.MakeGrades(grades.Entities), nil
}
