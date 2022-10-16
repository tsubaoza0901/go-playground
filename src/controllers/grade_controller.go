package controllers

import (
	"context"
	"go-playground/m/v1/usecase"
)

// Grade ...
type Grade struct {
	gradeUsecase usecase.IGradeUsecase
}

// NewGrade ...
func NewGrade(gu usecase.IGradeUsecase) *Grade {
	return &Grade{gu}
}

// GetGradeList ...
func (c *Grade) GetGradeList(ctx context.Context) {
	c.gradeUsecase.RetrieveGrades(ctx)
}
