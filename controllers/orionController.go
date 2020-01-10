package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/beegons/models"
	"github.com/beegons/utils"
)

func AlertsNotify(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		log.Println(err)
		return
	}

	var subNotification models.Subscription

	err = json.Unmarshal(b, &subNotification)
	if err != nil {
		log.Println(err)
		return
	}

	data := subNotification.Data[0]

	alertBytes, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
		return
	}

	hub := utils.GetWSAlertHub()
	hub.Broadcast <- alertBytes

	/*
		var subNotification models.Subscription
		var alert models.Alert

		err = json.Unmarshal(b, &subNotification)
		if err != nil {
			log.Println(err)
			return
		}
		// log.Println(subNotification)

		data := subNotification.Data[0]

		params := make(map[string]interface{})
		for k, v := range data {
			switch k {
			case "id":
				alert.Id = v.(string)
			case "type":
			case "condition":
				alert.Condition = v.(string)
			case "refModule":
				alert.RefModule = v.(string)
			case "dateObserved":
				alert.DateObserved = v.(string)
			default:
				params[k] = v
			}
		}
		alert.Parameters = params

		// log.Println(alert)

		alertBytes, err := json.Marshal(alert)
		if err != nil {
			log.Println(err)
			return
		}

		hub := utils.GetWSHub()
		hub.Broadcast <- alertBytes
	*/
}

func DataNotify(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		log.Println(err)
		return
	}

	var subNotification models.Subscription

	err = json.Unmarshal(b, &subNotification)
	if err != nil {
		log.Println(err)
		return
	}

	data := subNotification.Data[0]

	dataBytes, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
		return
	}

	hub := utils.GetWSDataHub()
	hub.Broadcast <- dataBytes

}


func ParkingSpotNotify(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		log.Println(err)
		return
	}

	var subNotification models.Subscription

	err = json.Unmarshal(b, &subNotification)
	if err != nil {
		log.Println(err)
		return
	}

	data := subNotification.Data[0]

	dataBytes, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
		return
	}

	hub := utils.GetWSPSHub()
	hub.Broadcast <- dataBytes

}
