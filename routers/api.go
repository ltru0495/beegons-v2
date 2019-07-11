package routers

import (
	"github.com/beegons/controllers/api"
	"github.com/gorilla/mux"
)

func SetApiRoutes(router *mux.Router) *mux.Router {
	apiRouter := mux.NewRouter()

	apiRouter.HandleFunc("/api/module/{id}", api.ModuleGet).Methods("GET")
	router.PathPrefix("/").Handler(apiRouter)
	return router

}
