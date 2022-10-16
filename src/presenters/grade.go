package presenters

import (
	"go-playground/m/v1/presenters/response"
	"go-playground/m/v1/usecase/data/output"
)

// Grade ...
type Grade struct {
	response.AppResponse
}

// NewGrade ...
func NewGrade() *Grade {
	return &Grade{}
}

// GradeList ...
func (p *Grade) GradeList(gradeList []*output.Grade) {
	data := make([]*response.Grade, len(gradeList))
	for i, v := range gradeList {
		data[i] = &response.Grade{
			ID:   v.ID,
			Name: v.Name,
		}
	}
	p.AppResponse = response.AppResponse{
		Data: data,
	}
}

// AppError ...
func (p *Grade) AppError(errorCode int) {
	p.AppResponse = response.AppResponse{
		Error: (*response.ErrorCode)(&errorCode),
	}
}
