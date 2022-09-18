package controllers

import (
	"go-playground/m/v1/src/adapters/controllers/http"

	gqlHandler "github.com/99designs/gqlgen/graphql/handler"
)

// AppController ...
type AppController struct {
	http.IUserHandler
	http.IGradeHandler
	http.IDealHistoryHandler
	http.IBalanceControlHandler
	*gqlHandler.Server
}

// NewAppController ...
func NewAppController(uh http.IUserHandler, gh http.IGradeHandler, dhh http.IDealHistoryHandler, bch http.IBalanceControlHandler, s *gqlHandler.Server) AppController {
	return AppController{uh, gh, dhh, bch, s}
}
