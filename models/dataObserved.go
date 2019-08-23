package models

import (
	"context"
	"github.com/beegons/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strings"
)

type DataObserved struct {
	Id           string             `json:"id" bson:"id"`
	RefModule    string             `json:"refModule"`
	Type         string             `json:"type" bson:"type`
	DataType     string             `json:"dataType"`
	DateObserved string             `json:"dateObserved" bson:"dateObserved"`
	Parameters   map[string]float64 ``
}

func mapToDataObserved(m map[string]interface{}) (d DataObserved) {
	params := make(map[string]float64)
	for k, v := range m {
		switch k {
		case "id", "type", "refModule", "dateObserved", "dataType":
			d.Id = v.(string)
		default:
			params[k] = v.(float64)
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

	d = mapToDataObserved(aux)

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

func GetLastDataObservedByParameter(id, parameter string) (d []map[string]interface{}, err error) {

	options := options.Find()

	// Sort by `_id` field descending
	options.SetSort(bson.D{{"dateObserved", 1}})

	cursor, err := getDatabase().Collection("data").Find(context.Background(), bson.D{{}}, options)
	if err != nil {
		return
	}
	var a map[string]interface{}

	for cursor.Next(context.TODO()) {
		aux := make(map[string]interface{})
		err = cursor.Decode(&a)
		if err != nil {
			return
		}
		aux["dateObserved"] = a["dateObserved"]
		aux[parameter] = a[parameter]
		d = append(d, aux)
	}
	if err = cursor.Err(); err != nil {
		return
	}
	cursor.Close(context.TODO())
	return

}
