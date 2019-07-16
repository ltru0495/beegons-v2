package models

import (
	"github.com/beegons/utils"
	"time"

	"log"
)

type AirQualityObserved struct {
	Id           string             `json:"id" bson:"id"`
	Type         string             `json:"type" bson:"type`
	DateObserved string             `json:"dateObserved" bson:"dateObserved"`
	Parameters   map[string]float64 `json:"parameters"`
}

type AirQualityObservedOrion map[string]interface{}

func NewAirQualityObserved(module Module) (a AirQualityObservedOrion) {
	a = make(map[string]interface{})
	a["id"] = "urn:ngsi-ld:AirQualityObserved:" + module.Name

	a["type"] = "AirQualityObserved"
	a["dateObserved"] = time.Now().Format("2006-01-02T15:04:05Z")

	for _, v := range module.ControlledProperties {
		a[v] = 0.0
	}
	return
}

func (a *AirQualityObservedOrion) CreateAirQualityObserved() (err error) {

	err = utils.PostEntity(a)
	return
}

func mapToAQO(m map[string]interface{}, properties []string) (aqo AirQualityObserved) {
	aqo.Id = m["id"].(string)
	aqo.Type = m["type"].(string)
	aqo.DateObserved = m["dateObserved"].(string)

	params := make(map[string]float64)
	for _, prop := range properties {
		val, ok := m[prop]
		if ok {
			params[prop] = val.(float64)
		}
	}
	aqo.Parameters = params

	return
}

func GetAirQualityObserved(id string) (aqo AirQualityObserved, err error) {
	module, err := GetModule(id)
	if err != nil {
		return
	}
	id = "urn:ngsi-ld:AirQualityObserved:" + id
	var aux map[string]interface{}
	err = utils.GetEntity(id, &aux)
	log.Println(aux, "ASDFSADF")
	aqo = mapToAQO(aux, module.ControlledProperties)

	return
}
