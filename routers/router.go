package routers

import (
	"github.com/gorilla/mux"
	"net/http"
)

func InitRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("assets/"))))
	router = SetApiRoutes(router)

	router = SetWSRoutes(router)

	router = SetAppRoutes(router)

	return router
}
