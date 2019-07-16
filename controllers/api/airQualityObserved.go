package api

import (
	"log"
	"net/http"

	"github.com/beegons/models"
	"github.com/gorilla/mux"
)

func AirQualityObservedGet(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]
	aqo, err := models.GetAirQualityObserved(id)
	// log.Println(aqo)
	if err != nil {
		log.Println(err)
		models.SendNotFound(w)
		return
	}
	models.SendData(w, aqo)
	return
}
