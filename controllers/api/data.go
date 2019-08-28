package api

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"

	"github.com/beegons/models"
)

func HistoricalData(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["moduleid"]
	parameter := vars["parameter"]

	m, err := models.GetModule(id)
	if err != nil {
		log.Println(err)
		models.SendNotFound(w)
		return
	}
	dataId := "urn:ngsi-ld:DataObserved:" + m.Name
	dataType := m.DataType + "Observed"

	if err != nil {
		log.Println(err)
		models.SendNotFound(w)
		return
	}

	start, err := time.Parse("2006-01-02T15:04Z", vars["start"])
	if err != nil {
		log.Println(err)
	}
	end, err := time.Parse("2006-01-02T15:04Z", vars["end"])

	res, err := models.GetHistoricalData(dataId, dataType, parameter, start, end)
	if err != nil {
		log.Println(err)
		models.SendNotFound(w)
		return
	}
	models.SendData(w, res)
}
