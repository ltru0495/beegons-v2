package api

import (
	"log"
	"net/http"

	"github.com/beegons/models"
	"github.com/gorilla/mux"
)

func ModuleGet(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]
	module, err := models.GetModule(id)
	if err != nil {
		log.Println(err)
		models.SendNotFound(w)
		return
	}
	models.SendData(w, module)
	return
}

func RealtimeInfoGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	module, err := models.GetModule(id)
	if err != nil {
		log.Println(err)
		models.SendNotFound(w)
		return
	}

	// data, err := models.GetDataObserved(id)
	// if err != nil {
	// 	log.Println(err)
	// 	models.SendNotFound(w)
	// }

	models.SendData(w, module)
	return
}
