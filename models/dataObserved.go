package models

import (
	"context"
	"github.com/beegons/utils"
	"go.mongodb.org/mongo-driver/bson"
	"strings"
)

type DataObserved struct {
	Id           string             `json:"id" bson:"id"`
	RefModule    string             `json:"refModule"`
	Type         string             `json:"type" bson:"type`
	DataType     string             `json:"dataType"`
	DateObserved string             `json:"dateObserved" bson:"dateObserved"`
	Parameters   map[string]float64 `json:"parameters"`
}

type DataObservedOrion map[string]interface{}

func mapToData(m map[string]interface{}, properties []string) (d DataObserved) {
	d.Id = m["id"].(string)
	d.Type = m["type"].(string)
	d.DateObserved = m["dateObserved"].(string)

	params := make(map[string]float64)
	for _, prop := range properties {
		val, ok := m[prop]
		if ok {
			params[prop] = val.(float64)
		}
	}
	d.Parameters = params

	return
}

func GetDataObserved(id string) (d DataObserved, err error) {
	module, err := GetModule(id)
	if err != nil {
		return
	}
	id = strings.Replace(module.Id, "Module", "DataObserved", -1)
	var aux map[string]interface{}
	err = utils.GetEntity(id, &aux)

	d = mapToData(aux, module.ControlledProperties)

	return
}

func GetLastDataObserved(id string) (d []map[string]interface{}, err error) {
	cursor, err := getDatabase().Collection("data").Find(context.Background(), bson.D{{}})
	if err != nil {
		return
	}
	var a map[string]interface{}
	for cursor.Next(context.TODO()) {
		err = cursor.Decode(&a)
		if err != nil {
			return
		}
		d = append(d, a)
	}
	if err = cursor.Err(); err != nil {
		return
	}
	cursor.Close(context.TODO())
	return

}
