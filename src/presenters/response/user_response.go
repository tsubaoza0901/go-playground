package response

// User ...
type User struct {
	Name string  `json:"name"`
	Age  uint    `json:"age"`
	Item []*Item `json:"item,omitempty"`
}
