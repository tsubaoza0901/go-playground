package dependency

import (
	"go-playground/m/v1/controllers"
	graphqlController "go-playground/m/v1/framework/web/graphql/graph/controllers"
	"go-playground/m/v1/presenters"

	"go-playground/m/v1/framework/web"
	"go-playground/m/v1/framework/web/graphql/graph"
	"go-playground/m/v1/framework/web/graphql/graph/generated"
	"go-playground/m/v1/framework/web/rest"
	"go-playground/m/v1/usecase/interactor/port"

	gqlHandler "github.com/99designs/gqlgen/graphql/handler"
)

// InitWebAPI ...
func (i Injection) InitWebAPI() *web.API {
	return web.NewAPI(
		i.InitAppHandler(),
	)
}

// InitAppHandler ...
func (i Injection) InitAppHandler() *web.AppHandler {
	return &web.AppHandler{
		BalanceControlHandler: i.InitBalanceControlHandler(),
		DealHistoryHandler:    i.InitDealHistoryHandler(),
		GradeHandler:          i.InitGradeHandler(),
		UserHandler:           i.InitUserHandler(),
		// i.InitGraphQLHandlerServer(),
	}
}

// InitUserHandler ...
func (i Injection) InitUserHandler() *rest.UserHandler {
	userPresenter := presenters.NewUser()
	return rest.NewUserHandler(
		i.InitUserController(userPresenter),
		userPresenter,
	)
}

// InitUserController ...
func (i Injection) InitUserController(userOutputPort port.UserOutput) *controllers.User {
	return controllers.NewUser(
		i.InitUserManagementUsecase(userOutputPort),
	)
}

// InitGradeHandler ...
func (i Injection) InitGradeHandler() *rest.GradeHandler {
	gradePresenter := presenters.NewGrade()
	return rest.NewGradeHandler(
		i.InitGradeController(gradePresenter),
		gradePresenter,
	)
}

// InitGradeController ...
func (i Injection) InitGradeController(gradeOutputPort port.GradeOutput) *controllers.Grade {
	return controllers.NewGrade(
		i.InitGradeUsecase(gradeOutputPort),
	)
}

// InitDealHistoryHandler ...
func (i Injection) InitDealHistoryHandler() *rest.DealHistoryHandler {
	dealHistoryPresenter := presenters.NewDealHistory()
	return rest.NewDealHistoryHandler(
		i.InitDealHistoryController(dealHistoryPresenter),
		dealHistoryPresenter,
	)
}

// InitDealHistoryController ...
func (i Injection) InitDealHistoryController(dealHistoryOutputPort port.DealHistoryOutput) *controllers.DealHistory {
	return controllers.NewDealHistory(
		i.InitDealHistoryUsecase(dealHistoryOutputPort),
	)
}

// InitBalanceControlHandler ...
func (i Injection) InitBalanceControlHandler() *rest.BalanceControlHandler {
	balancePresenter := presenters.NewBalance()
	return rest.NewBalanceControlHandler(
		i.InitBalanceController(balancePresenter),
		balancePresenter,
	)
}

// InitBalanceController ...
func (i Injection) InitBalanceController(balanceOutputPort port.BalanceOutput) *controllers.BalanceControl {
	return controllers.NewBalanceControl(
		i.InitBalanceControlUsecase(balanceOutputPort),
	)
}

// InitGraphQLHandlerServer ...
func (i Injection) InitGraphQLHandlerServer() *gqlHandler.Server {
	return gqlHandler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{
				Resolvers: i.InitResolver(),
			},
		),
	)
}

// InitResolver ...
func (i Injection) InitResolver() *graph.Resolver {
	userPresenter := presenters.NewUser()
	return graph.NewResolver(
		i.InitGraphQLUserController(userPresenter),
		userPresenter,
	)
}

// InitGraphQLUserController ...
func (i Injection) InitGraphQLUserController(userOutputPort port.UserOutput) *graphqlController.User {
	return graphqlController.NewUser(
		i.InitUserManagementUsecase(userOutputPort),
	)
}
