package response

import (
	"go-playground/m/v1/src/usecase/data/output"
)

// User ...
type User struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Age       uint   `json:"age"`
	GradeName string `json:"gradeName"`
}

// Users ...
type Users []User

// NewUser ...
func NewUser(u output.User) User {
	user := User{
		ID:        u.ID,
		Name:      u.MakeJPNFullName(),
		Age:       u.Age,
		GradeName: u.GradeName,
	}
	return user
}

// NewUsers ...
func NewUsers(us output.Users) Users {
	users := make([]User, len(us))
	for i, u := range us {
		users[i] = NewUser(u)
	}
	return users
}
