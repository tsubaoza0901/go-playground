package presenters

import (
	"go-playground/m/v1/presenters/response"
	"go-playground/m/v1/usecases/data/output"
	"net/http"
)

// User ...
type User struct {
	*response.AppResponse
}

// NewUser ...
func NewUser() *User {
	return &User{}
}

// UserWithItem ...
func (p *User) UserWithItem(user *output.UserWithItem) {
	body := response.User{
		Name: user.Name,
		Age:  user.Age,
	}
	if user.Items != nil {
		items := make([]*response.Item, len(user.Items))
		for i, v := range user.Items {
			items[i] = &response.Item{
				Name: v.Name,
			}
		}
		body.Item = items
	}

	p.AppResponse = &response.AppResponse{
		Status: http.StatusOK,
		Body:   body,
	}
}

// UserList ...
func (p *User) UserList(users []*output.User) {
	body := make([]*response.User, len(users))
	for i, v := range users {
		body[i] = &response.User{
			Name: v.Name,
			Age:  v.Age,
		}
	}

	p.AppResponse = &response.AppResponse{
		Status: http.StatusOK,
		Body:   body,
	}
}

// User ...
func (p *User) User(user *output.User) {
	body := response.User{
		Name: user.Name,
		Age:  user.Age,
	}
	p.AppResponse = &response.AppResponse{
		Status: http.StatusOK,
		Body:   body,
	}
}

// Error ...
func (p *User) Error(err error) {
	p.AppResponse = &response.AppResponse{
		Status: http.StatusOK,
		Body:   err.Error(),
	}
}
