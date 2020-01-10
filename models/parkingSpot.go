package models

import (
	// "encoding/json"
	"net/http"

	"github.com/gorilla/schema"

	"github.com/beegons/utils"
	geojson "github.com/paulmach/go.geojson"
)

type ParkingSpot struct {
	Id                string            `json:"id"`
	Name              string            `json:"name"`
	Type              string            `json:"type"`
	Status            bool              `json:"status"`
	EnergyStatus      string            	`json:"energyStatus"`
	Mac               string            `json:"mac" bson:"mac"`
	SupportedProtocol []string          `json:"supportedProtocol" bson:"protocol"`
	Coordinates       []float64         `json:"coordinates,omitempty"`
	Location          *geojson.Geometry `json:"location"`
}



func (p *ParkingSpot) DecodeParkingSpotForm(r *http.Request) error {
	err := r.ParseForm()
	if err != nil {
		return err
	}
	decoder := schema.NewDecoder()
	_ = decoder.Decode(p, r.PostForm)
	g := geojson.NewPointGeometry(p.Coordinates)

	p.Location = g

	p.Coordinates = nil
	p.Type = "ParkingSpot"

	p.EnergyStatus = "100"
	p.Status = true
	return nil
}


func (p *ParkingSpot) CreatePSSubscription() (err error) {
	id := prefix + "ParkingSpot:" + p.Name
	entities := []Entity{{Id: id}}
	subject := Subject{entities}

	url := utils.GetAlertURL() + "/ps/notify"

	protocol := HTTP{URL: url}
	notification := Notification{HTTP: protocol, AttrsFormat: "keyValues"}

	data := Payload{
		Description:  "Notify Beegons of all parkiing spot changes",
		Subject:      subject,
		Notification: notification,
	}
	err = utils.PostSubscription(data)
	return err
}


func (p *ParkingSpot) CreateParkingSpot() (err error) {
	p.Id = prefix + "ParkingSpot:" + p.Name
	p.Type = "ParkingSpot"
	err = utils.PostEntity(p)
	return
}

func (p *ParkingSpot) CreateDataSubscription() (err error) {
	id := prefix + "ParkingSpot:" + p.Name
	entities := []Entity{{Id: id}}
	subject := Subject{entities}

	url := utils.GetAlertURL() + "/parking/notify"

	protocol := HTTP{URL: url}
	notification := Notification{HTTP: protocol, AttrsFormat: "keyValues"}

	data := Payload{
		Description:  "Notify Beegons of all parking spot changes",
		Subject:      subject,
		Notification: notification,
	}
	err = utils.PostSubscription(data)
	return err
}

func GetAllParkingSpots() (parkingSpots []ParkingSpot, err error) {
	err = utils.GetEntities("ParkingSpot", &parkingSpots)
	return
}

func GetParkingSpot(id string) (parkingSpot ParkingSpot, err error) {
	err = utils.GetEntity(id, &parkingSpot)
	if err != nil {
		return
	}
	return
}
