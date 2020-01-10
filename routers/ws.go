package routers

import (
	"github.com/beegons/controllers"
	"github.com/gorilla/mux"
)

func SetWSRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/ws/alert", controllers.WSAlert).Methods("GET")
	router.HandleFunc("/ws/data", controllers.WSData).Methods("GET")
	router.HandleFunc("/ws/ps", controllers.WSPS).Methods("GET")
	return router

}
