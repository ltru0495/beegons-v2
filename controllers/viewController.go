package controllers

import (
	// "fmt"
	// "log"
	"net/http"
	// "strings"
	// "github.com/gorilla/mux"
	// "github.com/gorilla/schema"
	// "encoding/json"
	"fmt"
	"github.com/beegons/utils"
)

func Index(w http.ResponseWriter, r *http.Request) {

	context := make(map[string]interface{})
	// url := "http://localhost:1026/v2/entities?options=keyValues"
	// req, err := http.NewRequest("GET", url, nil)
	// CheckErr(err)

	// res, err := http.DefaultClient.Do(req)
	// CheckErr(err)
	// defer res.Body.Close()
	// var body []Module
	// err = json.NewDecoder(res.Body).Decode(&body)
	// CheckErr(err)
	// fmt.Println(body)

	modules, err := utils.GetEntities("Module")
	if err != nil {
		fmt.Println(err)
	}

	context["Modules"] = modules
	utils.RenderTemplate(w, "index", context)

}

type Module struct {
	Id                 string   `json:"id"`
	Type               string   `json:"type"`
	SerialNumber       string   `json:"serialNumber"`
	DeviceState        string   `json:"deviceState"`
	ControlledProperty []string `json:"controlledProperty"`
}

func CheckErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

// [
//     {
//         "category": [
//             "sensor"
//         ],
//         "id": "urn:ngsi-ld:Device:device-9845A",
//         "type": "Device"
//     },
//     {
//         "batteryLevel": 0.75,
//         "category": [
//             "sensor"
//         ],
//         "controlledProperty": [
//             "humidity",
//             "temperature",
//             "o3"
//         ],
//         "dateFirstUsed": {
//             "@type": "DateTime",
//             "@value": "2014-09-11T11:00:00Z"
//         },
//         "deviceState": "ok",
//         "id": "urn:ngsi-ld:Device:MOD1",
//         "serialNumber": "9845A",
//         "type": "Device"
//     }
// ]
