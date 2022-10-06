package output

// User ...
type User struct {
	Name string
	Age  uint
}

// UserWithItem ...
type UserWithItem struct {
	Name  string
	Age   uint
	Items []*Item
}
