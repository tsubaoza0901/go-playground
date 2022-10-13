package request

// User ...
type User struct {
	Name string `json:"name"`
	Age  uint   `json:"age"`
}

// GetUserByID ...
type GetUserByID struct {
	ID uint `param:"id"`
}
