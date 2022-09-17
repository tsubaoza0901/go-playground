package dto

import "go-playground/m/v1/src/domain/model/grade"

// FetchGradeResult ...
type FetchGradeResult struct {
	ID   uint
	Name string
}

// NewFetchGradeResult ...
func NewFetchGradeResult(id uint, name string) *FetchGradeResult {
	return &FetchGradeResult{
		ID:   id,
		Name: name,
	}
}

// ToGradeModel ...
func (g FetchGradeResult) ToGradeModel() grade.Entity {
	gradeEntity := grade.MakeEntity(grade.ID(g.ID), grade.Name(g.Name))
	return *gradeEntity
}

// FetchGradeListResult ...
type FetchGradeListResult []*FetchGradeResult

// ToGradesModel ...
func (gs FetchGradeListResult) ToGradesModel() grade.Entities {
	grades := make(grade.Entities, len(gs))
	for i, g := range gs {
		grades[i] = g.ToGradeModel()
	}
	return grades
}
