package routers

import (
	"github.com/beegons/controllers/api"
	"github.com/gorilla/mux"
)

func SetApiRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/api/module/{moduleid}", api.Module).Methods("GET")
	router.HandleFunc("/api/module/{moduleid}/parameters", api.ModuleParameters).Methods("GET")
	router.HandleFunc("/api/modules", api.GetModules).Methods("GET")
	router.HandleFunc("/api/module/{moduleid}/realtime", api.ModuleRealTime).Methods("GET")
	router.HandleFunc("/api/data/{moduleid}/{parameter}/{start}/{end}", api.HistoricalData).Methods("GET")
	router.HandleFunc("/api/file/{moduleid}/{parameter}/{start}/{end}/{format}", api.HistoricalData).Methods("GET")


	router.HandleFunc("/api/parkingspots", api.GetParkingSpots).Methods("GET")
	router.HandleFunc("/api/parkingspot/{id}", api.ParkingSpot).Methods("GET")
	// router.HandleFunc("/api/modulewdata/{id}", api.RealtimeInfoGet).Methods("GET")
	// router.HandleFunc("/api/lastdata/{id}/{parameter}", api.LastDataObserved).Methods("GET")

	return router

}
