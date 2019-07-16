package routers

import (
	"github.com/beegons/controllers/api"
	"github.com/gorilla/mux"
)

func SetApiRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/api/module/{id}", api.ModuleGet).Methods("GET")
	router.HandleFunc("/api/aqo/{id}", api.AirQualityObservedGet).Methods("GET")
	return router

}
