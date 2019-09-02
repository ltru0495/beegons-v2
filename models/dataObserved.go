package models

import (
	"context"
	"github.com/beegons/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strings"

	"log"
	"time"
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

func GetDataObserved(moduleid string) (d DataObserved, err error) {
	module, err := GetModule(moduleid)
	if err != nil {
		return
	}
	id := strings.Replace(module.Id, "Module", "DataObserved", -1)
	var aux map[string]interface{}
	err = utils.GetEntity(id, &aux)

	d = mapToDataObserved(aux)

	return
}

// sth_/_urn:ngsi-ld:DataObserved:MOD1_AirQualityObserved

/*
var data []models.Data
	iter := C.Find(bson.M{"date": bson.M{
		"$gt": fromDate,
		"$lt": toDate,
	}, "type": dataType, "id_moduleiot": moduleid, "id_sensor": sensorid}).Sort("date").Sort("-$natural").
			Limit(100).Sort("date").Sort("$natural").Iter()
*/
func GetHistoricalData(id, dataType, parameter string, start, end time.Time) (d []CygnusDocument, err error) {
	collection := "sth_/_" + id + "_" + dataType

	log.Println(start)
	log.Println(end)
	filter := bson.M{
		"attrName": parameter,
		"recvTime": bson.M{
			"$gte": start,
			"$lte": end,
		},
	}
	cursor, err := GetCygnusDatabase().Collection(collection).Find(context.Background(), filter)
	if err != nil {
		log.Println(err)
		return
	}

	err = cursor.All(context.Background(), &d)
	if err != nil {
		log.Println(err)
		return
	}
	for cursor.Next(context.TODO()) {
		aux := make(map[string]interface{})
		err = cursor.Decode(&aux)
		log.Println(aux)
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

	cursor, err := GetCygnusDatabase().Collection("data").Find(context.Background(), bson.D{{}}, options)
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
