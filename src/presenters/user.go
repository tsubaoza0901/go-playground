package presenters

import (
	"go-playground/m/v1/presenters/response"
	"go-playground/m/v1/usecase/data/output"
)

// User ...
type User struct {
	response.AppResponse
}

// NewUser ...
func NewUser() *User {
	return &User{}
}

// User ...
func (p *User) User(user *output.User) {
	data := &response.User{
		ID:           user.ID,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		Age:          user.Age,
		EmailAddress: user.EmailAddress,
		GradeName:    user.GradeName,
	}
	p.AppResponse = response.AppResponse{
		Data: data,
	}
}

// UserList ...
func (p *User) UserList(userList []*output.User) {
	data := make([]*response.User, len(userList))
	for i, v := range userList {
		data[i] = &response.User{
			ID:           v.ID,
			FirstName:    v.FirstName,
			LastName:     v.LastName,
			Age:          v.Age,
			EmailAddress: v.EmailAddress,
			GradeName:    v.GradeName,
		}
	}
	p.AppResponse = response.AppResponse{
		Data: data,
	}
}

// AppError ...
func (p *User) AppError(errorCode int) {
	p.AppResponse = response.AppResponse{
		Error: (*response.ErrorCode)(&errorCode),
	}
}
