package response

// User ...
type User struct {
	Name     string  `json:"name"`
	Age      uint    `json:"age"`
	ItemList []*Item `json:"item_list,omitempty"`
}
