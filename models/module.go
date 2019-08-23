package models

import (
	// "encoding/json"
	"github.com/gorilla/schema"
	"net/http"
	"time"

	"github.com/beegons/utils"
)

type Module struct {
	Id                string    `json:"id" bson:"id"`
	DataType          string    `json:"dataType"`
	Type              string    `json:"type" bson:"type"`
	Name              string    `json:"name" bson:"name"`
	Mac               string    `json:"mac" bson:"mac"`
	SupportedProtocol []string  `json:"supportedProtocol" bson:"protocol"`
	Coordinates       []float64 `json:"coordinates"`

	// ControlledProperties []string `json:"controlledProperties" bson:"controlledProperties"`
}

func (m *Module) DecodeModuleForm(r *http.Request) error {
	err := r.ParseForm()
	if err != nil {
		return err
	}
	decoder := schema.NewDecoder()
	_ = decoder.Decode(m, r.PostForm)
	m.Type = "Module"
	return nil
}

func (m *Module) CreateModule() (err error) {
	m.Id = "urn:ngsi-ld:Module:" + m.Name
	m.Type = "Module"
	err = utils.PostEntity(m)
	return
}

func (m *Module) CreateDataObserved() (err error) {
	d := make(map[string]interface{})
	d["id"] = "urn:ngsi-ld:DataObserved:" + m.Name
	d["type"] = m.DataType + "Observed"
	d["dateObserved"] = time.Now().Format("2006-01-02T15:04:05Z")
	d["refModule"] = m.Id
	err = utils.PostEntity(d)

	err = utils.PostSubscription(d["id"].(string), d["type"].(string))
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
