package input

// UserCreate ...
type UserCreate struct {
	FirstName    string
	LastName     string
	Age          uint
	EmailAddress string
}

// NewUserCreate ...
func NewUserCreate() UserCreate {
	return UserCreate{}
}
