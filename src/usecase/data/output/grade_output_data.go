package output

import "go-playground/m/v1/domain/model/grade"

// Grade ...
type Grade struct {
	ID   uint
	Name string
}

// MakeGrade ...
func MakeGrade(u grade.Entity) Grade {
	return Grade{
		ID:   uint(u.ID()),
		Name: u.Name(),
	}
}

// Grades ...
type Grades []Grade

// MakeGrades ...
func MakeGrades(gs grade.Entities) *Grades {
	grades := make(Grades, len(gs))
	for i, g := range gs {
		grades[i] = MakeGrade(g)
	}
	return &grades
}
