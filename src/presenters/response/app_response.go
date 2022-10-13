package response

// AppResponse ...
type AppResponse struct {
	Status int         `json:"status"`
	Body   interface{} `json:"body"`
}
