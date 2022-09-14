package model

import (
	"go-playground/m/v1/src/domain/model/grade"
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

func (g Grade) makeGradeEntity() grade.Entity {
	gradeEntity := grade.NewEntity(g.Name)
	gradeEntity.SetID(g.ID)
	return *gradeEntity
}

// Grades ...
type Grades []Grade

// MakeFetchAllDTO ...
func MakeFetchAllDTO(gs Grades) *grade.FetchAllDTO {
	gradeEntities := make(grade.Entities, len(gs))
	for i, g := range gs {
		gradeEntities[i] = g.makeGradeEntity()
	}
	gradeFetchAllDTO := grade.NewFetchAllDTO(gradeEntities)

	return &gradeFetchAllDTO
}
