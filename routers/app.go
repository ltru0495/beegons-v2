package routers

import (
	"github.com/gorilla/mux"
)

func SetAppRoutes(router *mux.Router) *mux.Router {
	appRouter := mux.NewRouter()
	appRouter.HandleFunc("/", nil)

	router.PathPrefix("/").Handler(appRouter)
	return router

}
