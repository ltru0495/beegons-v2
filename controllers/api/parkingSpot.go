package api

import (
	"log"
	"net/http"

	"github.com/beegons/models"
	"github.com/gorilla/mux"
)

func GetParkingSpots(w http.ResponseWriter, r *http.Request) {
	parkingSpots, err := models.GetAllParkingSpots()
	if err != nil {
		log.Println(err)
		models.SendNotFound(w)
		return
	}
	models.SendData(w, parkingSpots)
	return
}

func ParkingSpot(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	parkingSpot, err := models.GetParkingSpot(id)
	if err != nil {
		log.Println(err)
		models.SendNotFound(w)
		return
	}
	models.SendData(w, parkingSpot)
	return
}
