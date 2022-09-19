package controllers

import (
	"go-playground/m/v1/adapters/controllers/rest/handler"

	gqlHandler "github.com/99designs/gqlgen/graphql/handler"
)

// AppController ...
type AppController struct {
	handler.UserHandler
	handler.GradeHandler
	handler.DealHistoryHandler
	handler.BalanceControlHandler
	*gqlHandler.Server
}

// NewAppController ...
func NewAppController(uh handler.UserHandler, gh handler.GradeHandler, dhh handler.DealHistoryHandler, bch handler.BalanceControlHandler, s *gqlHandler.Server) AppController {
	return AppController{uh, gh, dhh, bch, s}
}
