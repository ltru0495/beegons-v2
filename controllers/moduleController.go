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
			module.CreateModule()

			aqo := models.NewAirQualityObserved(*module)
			aqo.CreateAirQualityObserved()

			if err != nil {
				log.Println(err)
				models.SendUnprocessableEntity(w)
				return
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
	modules, err := models.GetAllModules()
	if err != nil {
		log.Println(err)
	}

	context := make(map[string]interface{})
	context["Modules"] = modules

	utils.RenderTemplate(w, "module_table", context)
}
