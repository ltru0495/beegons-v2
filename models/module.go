package models

import (
	// "context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/schema"
	"strings"
	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	"errors"
	// "log"
	"net/http"
	"time"

	"github.com/beegons/utils"
)

type Module struct {
	// _Id                  primitive.ObjectID `json:"_id" bson:"_id"`
	Id                   string   `json:"id" bson:"id"`
	Type                 string   `json:"type" bson:"type"`
	DataType             string   `json:"dataType"`
	Mac                  string   `json:"mac" bson:"mac"`
	Name                 string   `json:"name" bson:"name"`
	State                string   `json:"state" bson:"state"`
	Protocol             string   `json:"protocol" bson:"protocol"`
	ControlledProperties []string `json:"controlledProperties" bson:"controlledProperties"`
	Latitude             float64  `json:"latitude"`
	Longitude            float64  `json:"longitude"`
}

type Sensor struct {
	Id        string `json:"id"`
	Name      string `json:"name" bson:"name"`
	RefModule string `json:"refModule" bson:"refModule"`
	Type      string `json:"type" bson:"type"`
	DataType  string `json:"dataType"`
	Parameter string `json:"parameter" bson:"parameter"`
	Model     string `json:"model" bson:"model"`
	Unit      string `json:"unit" bson:"unit"`
}

type ModuleDataSensors struct {
	Module  Module
	Data    DataObserved
	Sensors []Sensor
}

func (m *Module) DecodeModuleForm(r *http.Request) ([]Sensor, error) {
	var sensors []Sensor

	err := r.ParseForm()
	if err != nil {
		return sensors, err
	}
	decoder := schema.NewDecoder()
	_ = decoder.Decode(m, r.PostForm)

	sens := r.PostFormValue("sensors")
	err = json.Unmarshal([]byte(sens), &sensors)
	if err != nil {
		return sensors, err
	}

	var props []string
	for _, sensor := range sensors {
		props = append(props, sensor.Parameter)
	}
	m.ControlledProperties = props
	m.Type = "Module"
	return sensors, nil
}

func (m *Module) CreateModule() (err error) {
	m.Id = "urn:ngsi-ld:Module:" + m.Name
	m.Type = "Module"
	err = utils.PostEntity(m)
	return
}

func (m *Module) CreateSensors(sensors []Sensor) (err error) {
	for k, sensor := range sensors {
		sensor.Type = "Sensor"
		sensor.Id = strings.Replace(
			fmt.Sprintf("%s:S%03d", m.Id, k), "Module", "Sensor", -1)
		sensor.RefModule = m.Id

		err = utils.PostEntity(sensor)
		if err != nil {
			err = errors.New("Error while creating sensor")
			return
		}
	}
	return nil
}

func (m *Module) CreateDataObserved() (err error) {
	d := make(map[string]interface{})
	d["id"] = "urn:ngsi-ld:DataObserved:" + m.Name
	d["type"] = "DataObserved"
	d["dateObserved"] = time.Now().Format("2006-01-02T15:04:05Z")
	d["refModule"] = m.Id
	d["dataType"] = m.DataType
	for _, v := range m.ControlledProperties {
		d[v] = 0.0
	}
	err = utils.PostEntity(d)
	return
}

func GetAllModules() (modules []Module, err error) {
	err = utils.GetEntities("Module", &modules)
	return
}

func GetModule(id string) (module Module, err error) {
	err = utils.GetEntity(id, &module)
	if err != nil {
		return
	}
	return
}

func GetSensors(moduleid string) (sensors []Sensor, err error) {
	var s []Sensor
	err = utils.GetEntities("Sensor", &s)
	for _, v := range s {
		if v.RefModule == moduleid {
			sensors = append(sensors, v)
		}
	}
	if err != nil {
		return
	}
	return
}

/*==============================================================================================*/
/* 									DATABASE OPERATIONS 										*/
// const MODULE_COL = "modules"

// func (m *Module) Insert() error {
// 	m._Id = primitive.NewObjectID()
// 	res, err := getDatabase().Collection(MODULE_COL).InsertOne(context.Background(), m)
// 	m._Id = res.InsertedID.(primitive.ObjectID)
// 	return err
// }

// func (m *Module) Delete() error {
// 	_, err := getDatabase().Collection(MODULE_COL).DeleteOne(context.Background(), bson.D{{"_id", m.Id}})
// 	return err
// }

// func (m *Module) Update() error {
// 	update := bson.M{"$set": bson.M{
// 		"_id":      m._Id,
// 		"type":     m.Type,
// 		"mac":      m.Mac,
// 		"name":     m.Name,
// 		"state":    m.State,
// 		"protocol": m.Protocol,
// 	}}
// 	_, err := getDatabase().Collection(MODULE_COL).UpdateOne(context.Background(), bson.D{{"_id", m.Id}}, update)
// 	return err
// }

// func FindModule(id primitive.ObjectID) (m Module, err error) {
// 	filter := bson.D{{"_id", id}}
// 	err = getDatabase().Collection(MODULE_COL).FindOne(context.Background(), filter).Decode(&m)
// 	return
// }

// func AllModules() (modules []Module, err error) {
// 	cursor, err := getDatabase().Collection(MODULE_COL).Find(context.Background(), bson.D{{}})
// 	if err != nil {
// 		return
// 	}
// 	var m Module
// 	for cursor.Next(context.TODO()) {
// 		err = cursor.Decode(&m)
// 		if err != nil {
// 			return
// 		}
// 		modules = append(modules, m)
// 	}
// 	if err = cursor.Err(); err != nil {
// 		return
// 	}
// 	cursor.Close(context.TODO())
// 	return
// }
