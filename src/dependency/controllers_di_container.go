package dependency

import (
	"go-playground/m/v1/adapters/controllers"
	"go-playground/m/v1/adapters/controllers/graphql/graph"
	"go-playground/m/v1/adapters/controllers/graphql/graph/generated"
	"go-playground/m/v1/adapters/controllers/http/handler"

	gqlHandler "github.com/99designs/gqlgen/graphql/handler"
)

// InitAppController ...
func (i Injection) InitAppController() controllers.AppController {
	return controllers.NewAppController(
		i.InitUserHandler(),
		i.InitGradeHandler(),
		i.InitDealHistoryHandler(),
		i.InitBalanceControlHandler(),
		i.InitGraphQLHandlerServer(),
	)
}

// InitUserHandler ...
func (i Injection) InitUserHandler() handler.UserHandler {
	return handler.NewUserHandler(
		i.InitUserManagementUsecase(),
	)
}

// InitGradeHandler ...
func (i Injection) InitGradeHandler() handler.GradeHandler {
	return handler.NewGradeHandler(
		i.InitGradeUsecase(),
	)
}

// InitDealHistoryHandler ...
func (i Injection) InitDealHistoryHandler() handler.DealHistoryHandler {
	return handler.NewDealHistoryHandler(
		i.InitDealUsecase(),
	)
}

// InitBalanceControlHandler ...
func (i Injection) InitBalanceControlHandler() handler.BalanceControlHandler {
	return handler.NewBalanceControlHandler(
		i.InitBalanceControlUsecase(),
	)
}

// InitGraphQLHandlerServer ...
func (i Injection) InitGraphQLHandlerServer() *gqlHandler.Server {
	return gqlHandler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
		Resolvers: i.InitResolver(),
	}))
}

// InitResolver ...
func (i Injection) InitResolver() *graph.Resolver {
	return graph.NewResolver(
		i.InitUserManagementUsecase(),
	)
}
