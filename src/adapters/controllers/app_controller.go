package controllers

import "go-playground/m/v1/src/adapters/controllers/http"

// AppController ...
type AppController struct {
	http.IUserHandler
	http.IGradeHandler
	http.IDealHistoryHandler
	http.IBalanceControlHandler
}

// NewAppController ...
func NewAppController(uh http.IUserHandler, gh http.IGradeHandler, dhh http.IDealHistoryHandler, bch http.IBalanceControlHandler) AppController {
	return AppController{uh, gh, dhh, bch}
}
