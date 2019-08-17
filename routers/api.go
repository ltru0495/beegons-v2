package routers

import (
	"github.com/beegons/controllers/api"
	"github.com/gorilla/mux"
)

func SetApiRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/api/module/{id}", api.ModuleGet).Methods("GET")
	router.HandleFunc("/api/modulewdata/{id}", api.RealtimeInfoGet).Methods("GET")
	router.HandleFunc("/api/lastdata/{id}", api.LastDataObserved).Methods("GET")
	router.HandleFunc("/api/data/{id}", api.DataObservedGet).Methods("GET")
	return router

}
