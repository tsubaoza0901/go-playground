package response

// DealHistory ...
type DealHistory struct {
	Date     string `json:"date"`
	ItemName string `json:"item_name"`
	Amount   uint   `json:"amount"`
}
