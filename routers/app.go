package routers

import (
	"github.com/beegons/controllers"
	"github.com/gorilla/mux"
)

func SetAppRoutes(router *mux.Router) *mux.Router {
	appRouter := mux.NewRouter()
	appRouter.HandleFunc("/", controllers.Index)

	appRouter.HandleFunc("/user/create", controllers.UserCreate).Methods("GET", "POST")
	appRouter.HandleFunc("/users", controllers.UserTable).Methods("GET")

	appRouter.HandleFunc("/module/create", controllers.ModuleCreate).Methods("GET", "POST")
	appRouter.HandleFunc("/modules", controllers.ModuleTable).Methods("GET")

	appRouter.HandleFunc("/charts/realtime", controllers.ChartsRealTime).Methods("GET")
	appRouter.HandleFunc("/charts/historical", controllers.ChartsHistorical).Methods("GET")

	appRouter.HandleFunc("/alerts/notify", controllers.AlertsNotify).Methods("GET", "POST")

	// appRouter.HandleFunc("/data", controllers.OrionSubscription)
	router.PathPrefix("/").Handler(appRouter)
	return router

}
