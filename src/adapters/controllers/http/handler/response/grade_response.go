package response

import (
	"go-playground/m/v1/src/usecase/data/output"
)

// Grade ...
type Grade struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

// Grades ...
type Grades []Grade

// NewGrade ...
func NewGrade(g output.Grade) Grade {
	return Grade{
		ID:   g.ID,
		Name: g.Name,
	}
}

// NewGrades ...
func NewGrades(gs output.Grades) Grades {
	grades := make([]Grade, len(gs))
	for i, g := range gs {
		grades[i] = NewGrade(g)
	}
	return grades
}
