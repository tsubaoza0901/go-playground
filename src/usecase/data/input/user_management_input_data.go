package input

// UserCreate ...
type UserCreate struct {
	FirstName string
	LastName  string
	Age       uint
}

// NewUserCreate ...
func NewUserCreate() UserCreate {
	return UserCreate{}
}
