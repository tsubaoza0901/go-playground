// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Grade struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type NewUser struct {
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Age         int    `json:"age"`
	Email       string `json:"email"`
	TopUpAmount int    `json:"topUpAmount"`
}

type User struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Age          int    `json:"age"`
	EmailAddress string `json:"emailAddress"`
	GradeName    string `json:"gradeName"`
}
