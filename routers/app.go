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
	appRouter.HandleFunc("/module/modify", controllers.ModuleModify).Methods("GET", "POST")
	appRouter.HandleFunc("/modules", controllers.ModuleTable).Methods("GET")

	appRouter.HandleFunc("/alerts", controllers.Alerts).Methods("GET")

	appRouter.HandleFunc("/charts/realtime", controllers.ChartsRealTime).Methods("GET")
	appRouter.HandleFunc("/charts/historical", controllers.ChartsHistorical).Methods("GET")

	appRouter.HandleFunc("/alerts/notify", controllers.AlertsNotify).Methods("GET", "POST")
	appRouter.HandleFunc("/data/notify", controllers.DataNotify).Methods("GET", "POST")
	appRouter.HandleFunc("/ps/notify", controllers.ParkingSpotNotify).Methods("GET", "POST")

	appRouter.HandleFunc("/error", controllers.Error).Methods("GET")
	appRouter.HandleFunc("/maps/sensors", controllers.MapSensors).Methods("GET")
	appRouter.HandleFunc("/maps/parking", controllers.MapParking).Methods("GET")

	appRouter.HandleFunc("/parkingspot/create", controllers.ParkingSpotCreate).Methods("GET", "POST")
	appRouter.HandleFunc("/parkingspots", controllers.ParkingSpotTable).Methods("GET")

	// appRouter.HandleFunc("/login", controllers.Login).Methods("GET", "POST")
	// appRouter.HandleFunc("/logout", controllers.Logout).Methods("GET")

	// appRouter.HandleFunc("/data", controllers.OrionSubscription)
	router.PathPrefix("/").Handler(appRouter)
	return router

}
