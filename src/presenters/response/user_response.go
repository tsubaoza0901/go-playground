package response

// User ...
type User struct {
	ID           uint   `json:"id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Age          uint   `json:"age"`
	EmailAddress string `json:"email"`
	GradeName    string `json:"grade_name"`
}
