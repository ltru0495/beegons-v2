package routers

import (
	"github.com/beegons/controllers"
	"github.com/gorilla/mux"
)

func SetWSRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/ws/alert", controllers.WSAlert).Methods("GET")

	// router.HandleFunc("/api/modulewdata/{id}", api.RealtimeInfoGet).Methods("GET")
	// router.HandleFunc("/api/lastdata/{id}/{parameter}", api.LastDataObserved).Methods("GET")

	return router

}
