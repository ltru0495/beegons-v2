package models

import (
	// "encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/schema"

	"github.com/beegons/utils"
	geojson "github.com/paulmach/go.geojson"
)

type Module struct {
	Id                string            `json:"id" bson:"id"`
	DataType          string            `json:"dataType"`
	Type              string            `json:"type" bson:"type"`
	Name              string            `json:"name" bson:"name"`
	Mac               string            `json:"mac" bson:"mac"`
	SupportedProtocol []string          `json:"supportedProtocol" bson:"protocol"`
	Coordinates       []float64         `json:"coordinates,omitempty"`
	Location          *geojson.Geometry `json:"location"`
}

type Location struct {
	Type        string    `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}

var prefix string = "urn:ngsi-ld:"

func (m *Module) DecodeModuleForm(r *http.Request) error {
	err := r.ParseForm()
	if err != nil {
		return err
	}
	decoder := schema.NewDecoder()
	_ = decoder.Decode(m, r.PostForm)
	g := geojson.NewPointGeometry(m.Coordinates)

	m.Location = g

	m.Coordinates = nil
	m.Type = "Module"
	return nil
}

func (m *Module) CreateModule() (err error) {
	m.Id = prefix + "Module:" + m.Name
	m.Type = "Module"
	err = utils.PostEntity(m)
	return
}

func (m *Module) CreateDataObserved() (err error) {
	d := make(map[string]interface{})
	d["id"] = prefix + "DataObserved:" + m.Name
	d["type"] = m.DataType + "Observed"
	d["dataType"] = m.DataType + "Observed"
	d["dateObserved"] = time.Now().Format("2006-01-02T15:04:05Z")
	d["refModule"] = m.Id
	err = utils.PostEntity(d)
	return
}

func (m *Module) CreateAlert() (err error) {
	d := make(map[string]interface{})
	d["id"] = prefix + "Alert:" + m.Name
	d["type"] = "Alert"
	d["dateObserved"] = time.Now().Format("2006-01-02T15:04:05Z")
	d["refModule"] = m.Id
	d["condition"] = ""
	err = utils.PostEntity(d)
	return
}

func (m *Module) CreateAlertSubscription() (err error) {
	id := prefix + "Alert:" + m.Name
	entities := []Entity{{Id: id}}
	subject := Subject{entities}

	url := utils.GetAlertURL() + "/alerts/notify"

	protocol := HTTP{URL: url}
	notification := Notification{HTTP: protocol, AttrsFormat: "keyValues"}

	data := Payload{
		Description:  "Notify Server of alerts generated",
		Subject:      subject,
		Notification: notification,
	}

	err = utils.PostSubscription(data)
	return err
}

func (m *Module) CreateCygnusSubscription() (err error) {
	id := prefix + "DataObserved:" + m.Name
	entities := []Entity{{Id: id}}
	subject := Subject{entities}

	url := utils.GetCygnusURL() + "/notify"

	protocol := HTTP{URL: url}
	notification := Notification{HTTP: protocol, AttrsFormat: "legacy"}

	data := Payload{
		Description:  "Notify Cygnus of all sensor changes",
		Subject:      subject,
		Notification: notification,
		//		Throttling:   5,
	}
	err = utils.PostSubscription(data)
	return err
}

func (m *Module) CreateDataSubscription() (err error) {
	id := prefix + "DataObserved:" + m.Name
	entities := []Entity{{Id: id}}
	subject := Subject{entities}

	url := utils.GetAlertURL() + "/data/notify"

	protocol := HTTP{URL: url}
	notification := Notification{HTTP: protocol, AttrsFormat: "keyValues"}

	data := Payload{
		Description:  "Notify Beegons of all sensor changes",
		Subject:      subject,
		Notification: notification,
	}
	err = utils.PostSubscription(data)
	return err
}

func (m *Module) CreateFlinkSubscription() (err error) {
	id := prefix + "DataObserved:" + m.Name
	entities := []Entity{{Id: id}}
	subject := Subject{entities}

	url := utils.GetFlinkURL()

	protocol := HTTP{URL: url}
	notification := Notification{HTTP: protocol, AttrsFormat: "normalized"}

	data := Payload{
		Description:  "Notify Flink of all sensor changes",
		Subject:      subject,
		Notification: notification,
		//		Throttling:   5,
	}

	err = utils.PostSubscription(data)
	return err
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
