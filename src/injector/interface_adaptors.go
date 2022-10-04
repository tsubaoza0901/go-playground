package injector

// // WebFW ...
// type WebFW struct {
// 	c echo.Context
// 	// e *echo.Echo
// }

// // NewWebFW ...
// func NewWebFW(c echo.Context) *WebFW {
// 	return &WebFW{c}
// }

// // SetWebFW ...
// func (f *WebFW) SetWebFW(c echo.Context) {
// 	f.InitAppHandlers()
// 	NewAppPresenters(c)
// }

// // AppControllers ...
// type AppControllers struct {
// 	AppHandlers
// }

// // InitAppControllers ...
// func (f *WebFW) InitAppControllers() AppControllers {
// 	return AppControllers{
// 		AppHandlers: f.InitAppHandlers(),
// 	}
// }

// // AppHandlers ...
// type AppHandlers struct {
// 	handler.User
// }

// // InitAppHandlers ...
// func (f *WebFW) InitAppHandlers() AppHandlers {
// 	return AppHandlers{
// 		User: f.InitUserHandler(),
// 	}
// }

// // InitUserHandler ...
// func (f *WebFW) InitUserHandler() *handler.UserHandler {
// 	return handler.NewUserHandler(
// 		f.InitUserInteractor(),
// 	)
// }

// // ORM ...
// type ORM struct {
// 	dbConn string
// }

// // NewORM ...
// func NewORM(dbConn string) *ORM {
// 	return &ORM{dbConn}
// }

// // AppPresenters ...
// type AppPresenters struct {
// 	*presenters.UserPresenter
// }

// // NewAppPresenters ...
// func NewAppPresenters(c echo.Context) AppPresenters {
// 	return AppPresenters{
// 		UserPresenter: presenters.NewUserPresenter(c),
// 	}
// }

// // InitUserPresenter ...
// func (d *Dependency) InitUserPresenter() *presenters.UserPresenter {
// 	return presenters.NewUserPresenter(d.c)
// }

// // InitUserGateway ...
// func (d *ORM) InitUserGateway() *gateways.UserGateway {
// 	return gateways.NewUserGateway(d.dbConn)
// }
