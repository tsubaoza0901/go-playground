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

// UserUpdate ...
type UserUpdate struct {
	ID           uint
	LastName     string
	EmailAddress string
	GradeID      uint
}

// NewUserUpdate ...
func NewUserUpdate() UserUpdate {
	return UserUpdate{}
}
