package controllers

import (
	"github.com/beegons/models"
	"github.com/beegons/utils"

	"log"
	"net/http"
)

func ParkingSpotTable(w http.ResponseWriter, r *http.Request) {
	parkingSpots, err := models.GetAllParkingSpots()
	log.Println(parkingSpots)
	if err != nil {
		log.Println(err)
	}

	context := make(map[string]interface{})
	context["ParkingSpots"] = parkingSpots

	utils.RenderTemplate(w, "parking_spot_table", context)
}

func ParkingSpotCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		} else {
			parkingSpot := new(models.ParkingSpot)
			err := parkingSpot.DecodeParkingSpotForm(r)

			err = parkingSpot.CreateParkingSpot()
			if err != nil {
				log.Println(err)
				models.SendUnprocessableEntity(w)
				return
			}
			err = parkingSpot.CreatePSSubscription()
			if err != nil {
				log.Println(err)
				log.Println("Error While creating data subscription")
				models.SendUnprocessableEntity(w)
				return
			}


			/**************************************************/

			res := models.CreateDefaultResponse(w)
			res.Message = "ParkingSpot has been created"
			res.Content = parkingSpot
			res.Send()
			return
		}
	}
	utils.RenderTemplate(w, "parking_spot_create", nil)
}
