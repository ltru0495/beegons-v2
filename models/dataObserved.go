package models

import (
	"context"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/beegons/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DataObserved struct {
	Id           string             `json:"id" bson:"id"`
	RefModule    string             `json:"refModule"`
	Type         string             `json:"type" bson:"type"`
	DataType     string             `json:"dataType"`
	DateObserved string             `json:"dateObserved" bson:"dateObserved"`
	Parameters   map[string]float64 ``
}

type ModuleAndData struct {
	Module Module       `json:"module"`
	Data   DataObserved `json:"data"`
}

func mapToDataObserved(m map[string]interface{}) (d DataObserved) {
	params := make(map[string]float64)
	for k, v := range m {
		switch k {
		case "id":
			d.Id = v.(string)
		case "type":
			d.Type = v.(string)
		case "refModule":
			d.RefModule = v.(string)
		case "dataType":
			d.DataType = v.(string)
		case "dateObserved":
			d.DateObserved = v.(string)
		default:

			switch value := v.(type) {
			case float64:
				params[k] = value
			case string:
				floatValue, err := strconv.ParseFloat(value, 64)
				if err == nil {
					params[k] = floatValue
				}
				// log.Println(k, floatValue)
			default:
				// log.Println(value)

			}
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
func FilterDataByDate(id, dataType, parameter string, start, end time.Time) (d []CygnusDocument, err error) {
	collection := "sth_/_" + id + "_" + dataType
	start = FixDate(start)
	end = FixDate(end)
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

	if err = cursor.Err(); err != nil {
		return
	}
	cursor.Close(context.TODO())

	d = FixTimeZone(d)
	return
}

func GetLastData(id, dataType, parameter string, n int64) (d []CygnusDocument, err error) {
	collection := "sth_/_" + id + "_" + dataType
	options := options.Find()
	options.SetSort(bson.D{{"recvTime", -1}})
	options.SetLimit(n)

	// log.Println(id)
	// log.Println(dataType)
	// log.Println(parameter)
	// log.Println(n)
	// log.Println(collection)

	filter := bson.M{
		"attrName": parameter,
	}
	cursor, err := GetCygnusDatabase().Collection(collection).Find(context.Background(), filter, options)
	if err != nil {
		log.Println(err)
		return
	}

	err = cursor.All(context.Background(), &d)
	if err != nil {
		log.Println(err)
		return
	}

	if err = cursor.Err(); err != nil {
		return
	}
	cursor.Close(context.TODO())

	d = FixTimeZone(d)
	// log.Println(d)
	return

}

func FixDate(date time.Time) time.Time {
	return date.Add(time.Hour * 5)
}
func FixTimeZone(d []CygnusDocument) []CygnusDocument {
	for k, v := range d {
		date := v.RecvTime

		date = date.Add(-time.Hour * 5)
		d[k].RecvTime = date
	}
	return d
}
