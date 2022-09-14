package output

import (
	"fmt"
	"go-playground/m/v1/src/domain/model/user"
)

// User ...
type User struct {
	ID        uint
	FirstName string
	LastName  string
	Age       uint
	GradeName string
}

// MakeUser ...
func MakeUser(u user.General) User {
	return User{
		ID:        uint(u.ID()),
		FirstName: string(u.FirstName()),
		LastName:  string(u.LastName()),
		Age:       uint(u.Age()),
		GradeName: string(u.GradeName()),
	}
}

// MakeJPNFullName ...
func (u User) MakeJPNFullName() string {
	return fmt.Sprintf("%s %s", u.LastName, u.FirstName)
}

// Users ...
type Users []User

// MakeUsers ...
func MakeUsers(us user.Generals) Users {
	users := make(Users, len(us))
	for i, u := range us {
		users[i] = MakeUser(u)
	}
	return users
}
