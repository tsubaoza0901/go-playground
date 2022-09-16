package model

import (
	"go-playground/m/v1/src/domain/model/grade"
	"go-playground/m/v1/src/usecase/repository/dto"
)

// Grade ...
type Grade struct {
	ID   uint
	Name string
}

// TableName ...
func (Grade) TableName() string {
	return "grades"
}

// MakeFetchGradeResultDTO ...
func MakeFetchGradeResultDTO(g Grade) *dto.FetchGradeResult {
	fetchGradeResult := dto.NewFetchGradeResult(
		grade.ID(g.ID),
		grade.Name(g.Name),
	)
	return fetchGradeResult
}

// Grades ...
type Grades []Grade

// MakeFetchGradeListResultDTO ...
func MakeFetchGradeListResultDTO(gs Grades) *dto.FetchGradeListResult {
	fetchGradeListResult := make(dto.FetchGradeListResult, len(gs))
	for i, g := range gs {
		fetchGradeListResult[i] = MakeFetchGradeResultDTO(g)
	}
	return &fetchGradeListResult
}
