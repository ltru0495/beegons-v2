package api

import (
	"log"
	"net/http"

	"github.com/beegons/models"
	"github.com/gorilla/mux"
)

func DataObservedGet(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]
	data, err := models.GetDataObserved(id)
	// log.Println(data)
	if err != nil {
		log.Println(err)
		models.SendNotFound(w)
		return
	}
	models.SendData(w, data)
	return
}

func LastDataObserved(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]
	data, err := models.GetLastDataObserved(id)
	// log.Println(data)
	if err != nil {
		log.Println(err)
		models.SendNotFound(w)
		return
	}
	models.SendData(w, data)
	return
}
