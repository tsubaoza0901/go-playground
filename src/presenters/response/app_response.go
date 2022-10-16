package response

// AppResponse ...
type AppResponse struct {
	Data  interface{}
	Error *ErrorCode
}

// // AppError ...
// func (a *AppResponse) AppError(errorCode int) {
// 	a = &AppResponse{
// 		Error: (*ErrorCode)(&errorCode),
// 	}
// }
