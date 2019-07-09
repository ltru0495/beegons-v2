package controllers

import (
	"github.com/beegons/models"
	"github.com/beegons/utils"

	"log"
	"net/http"
)

func ModuleCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		} else {
			module := new(models.Module)
			sensors, err := module.DecodeModuleForm(r)

			if err != nil {
				log.Println(err)
			}

			// log.Println(module)

			err = utils.PostEntity(module)
			if err != nil {
				log.Println(err)
			}
			res := models.CreateDefaultResponse(w)
			res.Message = "Module has been created"
			res.Data = sensors
			res.Send()
			return
		}
	}
	utils.RenderTemplate(w, "module_create", nil)
}

func ModuleTable(w http.ResponseWriter, r *http.Request) {
	modules, err := models.AllModules()
	if err != nil {
		log.Println(err)
	}

	context := make(map[string]interface{})
	context["Modules"] = modules

	utils.RenderTemplate(w, "module_table", context)
}

/*

curl -iX POST \
  'http://localhost:1026/v2/entities' \
  -H 'Content-Type: application/json' \
  -d '
{
    "id": "urn:ngsi-ld:Store:001",
    "type": "Store",
    "address": {
        "type": "PostalAddress",
        "value": {
            "streetAddress": "Bornholmer Straße 65",
            "addressRegion": "Berlin",
            "addressLocality": "Prenzlauer Berg",
            "postalCode": "10439"
        }
    },
    "location": {
        "type": "geo:json",
        "value": {
             "type": "Point",
             "coordinates": [13.3986, 52.5547]
        }
    },
    "name": {
        "type": "Text",
        "value": "Bösebrücke Einkauf"
    }
}'
*/
