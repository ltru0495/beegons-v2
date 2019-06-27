package routers

import (
	"github.com/beegons/controllers"
	"github.com/gorilla/mux"
)

func SetAppRoutes(router *mux.Router) *mux.Router {
	appRouter := mux.NewRouter()
	appRouter.HandleFunc("/", controllers.Index)

	appRouter.HandleFunc("/user/create", controllers.UserCreate).Methods("GET")
	router.PathPrefix("/").Handler(appRouter)
	return router

}
