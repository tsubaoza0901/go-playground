package dependency

import "go-playground/m/v1/src/adapters/controllers"

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
