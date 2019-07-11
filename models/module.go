package models

import (
	// "context"
	"encoding/json"
	"github.com/gorilla/schema"
	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

type Module struct {
	// _Id                  primitive.ObjectID `json:"_id" bson:"_id"`
	Id                   string   `json:"id" bson:"id"`
	Type                 string   `json:"type" bson:"type"`
	Mac                  string   `json:"mac" bson:"mac"`
	Name                 string   `json:"name" bson:"name"`
	State                string   `json:"state" bson:"state"`
	Protocol             string   `json:"protocol" bson:"protocol"`
	ControlledProperties []string `json:"controlledProperties" bson:"controlledProperties"`
}

type Sensor struct {
	Name  string `json:"name" bson:"name"`
	Type  string `json:"type" bson:"type"`
	Model string `json:"model" bson:"model"`
	Unit  string `json:"Unit" bson:"Unit"`
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
		props = append(props, sensor.Type)
	}
	m.ControlledProperties = props
	return sensors, nil
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
